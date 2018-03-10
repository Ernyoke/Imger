package edgedetection

import (
	"github.com/Ernyoke/Imger/blur"
	"github.com/Ernyoke/Imger/padding"
	"github.com/Ernyoke/Imger/utils"
	"image"
	"image/color"
	"math"
	"github.com/Ernyoke/Imger/grayscale"
)

// Computes the edges of a given image using the Canny edge detection algorithm.
func CannyGray(img *image.Gray, lower float64, upper float64, kernelSize uint) (*image.Gray, error) {

	// blur the image using Gaussian filter
	blurred, error := blur.GaussianBlurGray(img, float64(kernelSize), 3, padding.BorderReflect)
	if error != nil {
		return nil, error
	}

	// get vertical and horizontal edges using Sobel filter
	vertical, error := VerticalSobelGray(blurred, padding.BorderReflect)
	if error != nil {
		return nil, error
	}
	horizontal, error := HorizontalSobelGray(blurred, padding.BorderReflect)
	if error != nil {
		return nil, error
	}

	// calculate the gradient values and orientation angles for each pixel
	g, theta := gradientAndOrientation(vertical, horizontal)

	// "thin" the edges using non-max suppression procedure
	thinEdges := nonMaxSuppression(blurred, g, theta)

	// hysteresis
	hist := threshold(thinEdges, g, lower, upper)

	return hist, nil
}

func CannyRGBA(img *image.RGBA, lower float64, upper float64, kernelSize uint) (*image.Gray, error) {
	return CannyGray(grayscale.Grayscale(img), lower, upper, kernelSize)
}

func gradientAndOrientation(vertical *image.Gray, horizontal *image.Gray) ([][]float64, [][]float64) {
	size := vertical.Bounds().Size()
	theta := make([][]float64, size.X)
	g := make([][]float64, size.X)
	for x := 0; x < size.X; x++ {
		theta[x] = make([]float64, size.Y)
		g[x] = make([]float64, size.Y)
		for y := 0; y < size.Y; y++ {
			px := float64(vertical.GrayAt(x, y).Y)
			py := float64(horizontal.GrayAt(x, y).Y)
			g[x][y] = math.Hypot(px, py)
			theta[x][y] = orientation(math.Atan2(float64(vertical.GrayAt(x, y).Y), float64(horizontal.GrayAt(x, y).Y)))
		}
	}
	return g, theta
}

func isBetween(val float64, lowerBound float64, upperBound float64) bool {
	return val >= lowerBound && val < upperBound
}

func orientation(x float64) float64 {
	angle := 180 * x / math.Pi
	if isBetween(angle, -22.5, 22.5) || isBetween(angle, -180, -157.5) {
		return 0
	}
	if isBetween(angle, 22.5, 67.5) || isBetween(angle, -157.5, -112.5) {
		return 0
	}
	if isBetween(angle, 157.5, 180) || isBetween(angle, -22.5, 0) {
		return 45
	}
	if isBetween(angle, 67.5, 112.5) || isBetween(angle, -112.5, -67.5) {
		return 90
	}
	if isBetween(angle, 112.5, 157.5) || isBetween(angle, -67.5, -22.5) {
		return 135
	}
	return -1
}

func isBiggerThenNeighbours(val float64, neighbour1 float64, neighbour2 float64) bool {
	return val > neighbour1 && val > neighbour2
}

func nonMaxSuppression(img *image.Gray, g [][]float64, theta [][]float64) *image.Gray {
	size := img.Bounds().Size()
	thinEdges := image.NewGray(image.Rect(0, 0, size.X, size.Y))
	utils.ForEachPixel(size, func(x, y int) {
		isLocalMax := false
		if x > 0 && x < size.X-1 && y > 0 && y < size.Y-1 {
			switch theta[x][y] {
			case 45:
				if isBiggerThenNeighbours(g[x][y], g[x-1][y-1], g[x+1][y+1]) {
					isLocalMax = true
				}
			case 90:
				if isBiggerThenNeighbours(g[x][y], g[x][y+1], g[x][y-1]) {
					isLocalMax = true
				}
			case 135:
				if isBiggerThenNeighbours(g[x][y], g[x+1][y-1], g[x-1][y+1]) {
					isLocalMax = true
				}
			case 0:
				if isBiggerThenNeighbours(g[x][y], g[x+1][y], g[x-1][y]) {
					isLocalMax = true
				}
			}
		} else {
			isLocalMax = true
		}
		if isLocalMax {
			thinEdges.SetGray(x, y, color.Gray{Y: utils.MaxUint8})
		}
	})
	return thinEdges
}

func threshold(img *image.Gray, g [][]float64, lowerBound float64, upperBound float64) *image.Gray {
	size := img.Bounds().Size()
	res := image.NewGray(image.Rect(0, 0, size.X, size.Y))
	utils.ForEachPixel(size, func(x int, y int) {
		p := img.GrayAt(x, y)
		if p.Y == utils.MaxUint8 {
			if g[x][y] < lowerBound {
				res.SetGray(x, y, color.Gray{Y: utils.MinUint8})
			}
			if g[x][y] > upperBound {
				res.SetGray(x, y, color.Gray{Y: utils.MaxUint8})
			}
		}
	})
	utils.ForEachPixel(size, func(x int, y int) {
		p := img.GrayAt(x, y)
		if p.Y == utils.MaxUint8 {
			if g[x][y] >= lowerBound && g[x][y] <= upperBound {
				if res.GrayAt(x-1, y-1).Y == utils.MaxUint8 || res.GrayAt(x-1, y).Y == utils.MaxUint8 ||
					res.GrayAt(x-1, y+1).Y == utils.MaxUint8 || res.GrayAt(x, y-1).Y == utils.MaxUint8 ||
					res.GrayAt(x, y+1).Y == utils.MaxUint8 || res.GrayAt(x+1, y-1).Y == utils.MaxUint8 ||
					res.GrayAt(x+1, y).Y == utils.MaxUint8 || res.GrayAt(x+1, y+1).Y == utils.MaxUint8 {
					res.SetGray(x, y, color.Gray{Y: utils.MinUint8})
				}
			}
		}
	})
	return res
}
