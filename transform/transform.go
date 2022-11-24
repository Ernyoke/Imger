package transform

import (
	"errors"
	"github.com/ernyoke/imger/utils"
	"image"
	"math"
)

// RotateGray rotates a grayscale image counterclockwise with a given angle. The point which will represent the center
// ot the rotation is specified by the anchor argument. The result image can have its original size or it can be
// resized to fit in the area of the image.
// Example of usage:
//
// 		res, err := transform.RotateGray(img, 90.0, {512, 512}, true)
//
func RotateGray(img *image.Gray, angle float64, anchor image.Point, resizeToFit bool) (*image.Gray, error) {
	size := img.Bounds().Size()
	if anchor.X < 0 || anchor.Y < 0 || anchor.X > size.X || anchor.Y > size.Y {
		return nil, errors.New("invalid anchor position")
	}
	radians := angleToRadians(angle)
	newSize := size
	if resizeToFit {
		newSize = computeFitSize(size, radians)
	}
	result := image.NewGray(image.Rect(0, 0, newSize.X, newSize.Y))
	utils.ParallelForEachPixel(newSize, func(x, y int) {
		result.SetGray(x, y, img.GrayAt(getOriginalPixelPosition(x, y, radians, anchor, computeOffset(size, newSize))))
	})
	return result, nil
}

// RotateRGBA rotates an RGBA image counterclockwise with a given angle. The point which will represent the center
// ot the rotation is specified by the anchor argument. The result image can have its original size or it can be
// resized to fit in the area of the image.
// Example of usage:
//
// 		res, err := transform.RotateGray(img, 90.0, {512, 512}, true)
//
func RotateRGBA(img *image.RGBA, angle float64, anchor image.Point, resizeToFit bool) (*image.RGBA, error) {
	size := img.Bounds().Size()
	if anchor.X < 0 || anchor.Y < 0 || anchor.X > size.X || anchor.Y > size.Y {
		return nil, errors.New("invalid anchor position")
	}
	radians := angleToRadians(angle)
	newSize := size
	if resizeToFit {
		newSize = computeFitSize(size, radians)
	}
	result := image.NewRGBA(image.Rect(0, 0, newSize.X, newSize.Y))
	utils.ParallelForEachPixel(newSize, func(x, y int) {
		result.SetRGBA(x, y, img.RGBAAt(getOriginalPixelPosition(x, y, radians, anchor, computeOffset(size, newSize))))
	})
	return result, nil
}

func angleToRadians(angle float64) float64 {
	return angle * (math.Pi / 180)
}

func computeFitSize(size image.Point, radians float64) image.Point {
	a := math.Abs(float64(size.X) * math.Sin(radians))
	b := math.Abs(float64(size.X) * math.Cos(radians))
	c := math.Abs(float64(size.Y) * math.Sin(radians))
	d := math.Abs(float64(size.Y) * math.Cos(radians))
	return image.Point{X: int(c + b), Y: int(a + d)}
}

func computeOffset(size image.Point, fittingSize image.Point) image.Point {
	return image.Point{X: (fittingSize.X - size.X) / 2, Y: (fittingSize.Y - size.Y) / 2}
}

func getOriginalPixelPosition(x int, y int, radians float64, anchor image.Point, offset image.Point) (int, int) {
	dx := x - anchor.X - offset.X
	dy := y - anchor.Y - offset.Y
	originalX := int(math.Floor(math.Cos(radians)*float64(dx) - math.Sin(radians)*float64(dy) + float64(anchor.X)))
	originalY := int(math.Floor(math.Sin(radians)*float64(dx) + math.Cos(radians)*float64(dy) + float64(anchor.Y)))
	return originalX, originalY
}
