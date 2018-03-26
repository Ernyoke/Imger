package transform

import (
	"errors"
	"github.com/Ernyoke/Imger/utils"
	"image"
	"math"
)

func RotateGray(img *image.Gray, angle int, anchor image.Point, resizeToFit bool) (*image.Gray, error) {
	size := img.Bounds().Size()
	if anchor.X < 0 || anchor.Y < 0 || anchor.X > size.X || anchor.Y > size.Y {
		return nil, errors.New("invalid anchor position")
	}
	radians := float64(angle) * (math.Pi / 180)
	newSize := size
	if resizeToFit {
		a := math.Abs(float64(size.X) * math.Sin(radians))
		b := math.Abs(float64(size.X) * math.Cos(radians))
		c := math.Abs(float64(size.Y) * math.Sin(radians))
		d := math.Abs(float64(size.Y) * math.Cos(radians))
		newSize.X = int(c + b)
		newSize.Y = int(a + d)
	}
	offsetX := (newSize.X - size.X) / 2
	offsetY := (newSize.Y - size.Y) / 2
	result := image.NewGray(image.Rect(0, 0, newSize.X, newSize.Y))
	utils.ParallelForEachPixel(newSize, func(x, y int) {
		dx := x - anchor.X - offsetX
		dy := y - anchor.Y - offsetY

		newX := int(math.Floor(math.Cos(radians)*float64(dx) - math.Sin(radians)*float64(dy) + float64(anchor.X)))
		newY := int(math.Floor(math.Sin(radians)*float64(dx) + math.Cos(radians)*float64(dy) + float64(anchor.Y)))

		result.SetGray(x, y, img.GrayAt(newX, newY))
	})
	return result, nil
}

func RotateRGBA(img *image.RGBA, angle int, anchor image.Point, resizeToFit bool) (*image.RGBA, error) {
	size := img.Bounds().Size()
	if anchor.X < 0 || anchor.Y < 0 || anchor.X > size.X || anchor.Y > size.Y {
		return nil, errors.New("invalid anchor position")
	}
	radians := float64(angle) * (math.Pi / 180)
	newSize := size
	if resizeToFit {
		a := math.Abs(float64(size.X) * math.Sin(radians))
		b := math.Abs(float64(size.X) * math.Cos(radians))
		c := math.Abs(float64(size.Y) * math.Sin(radians))
		d := math.Abs(float64(size.Y) * math.Cos(radians))
		newSize.X = int(c + b)
		newSize.Y = int(a + d)
	}
	offsetX := (newSize.X - size.X) / 2
	offsetY := (newSize.Y - size.Y) / 2
	result := image.NewRGBA(image.Rect(0, 0, newSize.X, newSize.Y))
	utils.ParallelForEachPixel(newSize, func(x, y int) {
		dx := x - anchor.X - offsetX
		dy := y - anchor.Y - offsetY

		newX := int(math.Floor(math.Cos(radians)*float64(dx) - math.Sin(radians)*float64(dy) + float64(anchor.X)))
		newY := int(math.Floor(math.Sin(radians)*float64(dx) + math.Cos(radians)*float64(dy) + float64(anchor.Y)))

		result.SetRGBA(x, y, img.RGBAAt(newX, newY))
	})
	return result, nil
}
