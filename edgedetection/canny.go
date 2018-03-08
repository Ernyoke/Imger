package edgedetection

import (
	"image"
	"github.com/Ernyoke/Imger/blur"
	"github.com/Ernyoke/Imger/padding"
	"math"
	"fmt"
)

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

//func nonMaxSuppression(size image.Point, g [][]float64, theta [][]float64) *image.Gray {
//	thinEdges := image.NewGray(image.Rect(0, 0, size.X, size.Y))
//}

func CannyGray(img *image.Gray, lower uint8, upper uint8, kernelSize uint) (*image.Gray, error) {
	blured, error := blur.GaussianBlurGray(img, float64(kernelSize), 6, padding.BorderReflect)
	if error != nil {
		return nil, error
	}
	gx, error := VerticalSobelGray(blured, padding.BorderReflect)
	if error != nil {
		return nil, error
	}
	gy, error := HorizontalSobelGray(blured, padding.BorderReflect)
	if error != nil {
		return nil, error
	}
	size := img.Bounds().Size()
	theta := make([][]float64, size.X)
	g := make([][]float64, size.X)
	for x := 0; x < size.X; x++ {
		theta[x] = make([]float64, size.Y)
		g[x] = make([]float64, size.Y)
		for y := 0; y < size.Y; y++ {
			px := float64(gx.GrayAt(x, y).Y)
			py := float64(gy.GrayAt(x, y).Y)
			g[x][y] = math.Sqrt(px * px + py * py)
			theta[x][y] = orientation(math.Atan2(float64(gx.GrayAt(x, y).Y), float64(gy.GrayAt(x, y).Y)))
		}
	}

	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			fmt.Printf("%f ", theta[x][y])
		}
		fmt.Printf("\n")
	}

	return gx, nil
}