package padding

import (
	"image"
	"errors"
	"image/color"
)

type Border int

const (
	BorderConstant Border = iota
	BorderReplicate
	BorderReflect
)

type Paddings struct {
	PaddingLeft int
	PaddingRight int
	PaddingTop int
	PaddingBottom int
}

func topPaddingReplicate(img image.Image, p Paddings, setPixel func(int, int, color.Color)) {
	originalSize := img.Bounds().Size()
	for x := p.PaddingLeft; x < originalSize.X + p.PaddingLeft; x++ {
		firstPixel := img.At(x - p.PaddingLeft, p.PaddingTop)
		for y := 0; y < p.PaddingTop; y++ {
			setPixel(x, y, firstPixel)
		}
	}
}

func bottomPaddingReplicate(img image.Image, p Paddings, setPixel func(int, int, color.Color)) {
	originalSize := img.Bounds().Size()
	for x := p.PaddingLeft; x < originalSize.X + p.PaddingLeft; x++ {
		lastPixel := img.At(x - p.PaddingLeft, originalSize.Y - 1)
		for y := p.PaddingTop + originalSize.Y; y < originalSize.Y + p.PaddingTop + p.PaddingBottom; y++ {
			setPixel(x, y, lastPixel)
		}
	}
}

func leftPaddingReplicate(img image.Image, padded image.Image, p Paddings, setPixel func(int, int, color.Color)) {
	originalSize := img.Bounds().Size()
	for y := 0; y < originalSize.Y + p.PaddingBottom + p.PaddingTop; y++ {
		firstPixel := padded.At(p.PaddingLeft, y)
		for x := 0; x < p.PaddingLeft; x++ {
			setPixel(x, y, firstPixel)
		}
	}
}

func rightPaddingReplicate(img image.Image, padded image.Image, p Paddings, setPixel func(int, int, color.Color)) {
	originalSize := img.Bounds().Size()
	for y := 0; y < originalSize.Y + p.PaddingBottom + p.PaddingTop; y++ {
		lastPixel := padded.At(originalSize.X + p.PaddingLeft - 1, y)
		for x := originalSize.X + p.PaddingLeft; x < originalSize.X + p.PaddingLeft + p.PaddingRight; x++ {
			setPixel(x, y, lastPixel)
		}
	}
}

func topPaddingReflect(img image.Image, p Paddings, setPixel func(int, int, color.Color)) {
	originalSize := img.Bounds().Size()
	for x := p.PaddingLeft; x < originalSize.X + p.PaddingLeft; x++ {
		for y := 0; y < p.PaddingTop; y++ {
			pixel := img.At(x - p.PaddingLeft, p.PaddingTop - y)
			setPixel(x, y, pixel)
		}
	}
}

func bottomPaddingReflect(img image.Image, p Paddings, setPixel func(int, int, color.Color)) {
	originalSize := img.Bounds().Size()
	for x := p.PaddingLeft; x < originalSize.X + p.PaddingLeft; x++ {
		for y := p.PaddingTop + originalSize.Y; y < originalSize.Y + p.PaddingTop + p.PaddingBottom; y++ {
			pixel := img.At(x - p.PaddingLeft, originalSize.Y - (y - p.PaddingTop - originalSize.Y) - 1)
			setPixel(x, y, pixel)
		}
	}
}

func leftPaddingReflect(img image.Image, padded image.Image, p Paddings, setPixel func(int, int, color.Color)) {
	originalSize := img.Bounds().Size()
	for y := 0; y < originalSize.Y + p.PaddingBottom + p.PaddingTop; y++ {
		for x := 0; x < p.PaddingLeft; x++ {
			pixel := padded.At(2 * p.PaddingLeft - x, y)
			setPixel(x, y, pixel)
		}
	}
}

func rightPaddingReflect(img image.Image, padded image.Image, p Paddings, setPixel func(int, int, color.Color)) {
	originalSize := img.Bounds().Size()
	for y := 0; y < originalSize.Y + p.PaddingBottom + p.PaddingTop; y++ {
		for x := originalSize.X + p.PaddingLeft; x < originalSize.X + p.PaddingLeft + p.PaddingRight; x++ {
			pixel := padded.At(originalSize.X + p.PaddingRight - (x - originalSize.X - p.PaddingLeft), y)
			setPixel(x, y, pixel)
		}
	}
}


func PaddingGray(img *image.Gray, kernelSize image.Point, anchor image.Point, border Border) (*image.Gray, error) {
	originalSize := img.Bounds().Size()
	p, error := calculatePaddings(kernelSize, anchor)
	if error != nil {
		return nil, error
	}
	rect := getReactangleFromPaddings(p, originalSize)
	padded := image.NewGray(rect)

	for x := p.PaddingLeft; x < originalSize.X + p.PaddingLeft; x++ {
		for y := p.PaddingTop; y < originalSize.Y + p.PaddingTop; y++ {
			padded.Set(x, y, img.GrayAt(x - p.PaddingLeft, y - p.PaddingTop))
		}
	}

	switch border {
	case BorderConstant:
		// do nothing
	case BorderReplicate:
		topPaddingReplicate(img, p, func(x int, y int, pixel color.Color){
			padded.Set(x, y, pixel)
		})
		bottomPaddingReplicate(img, p, func(x int, y int, pixel color.Color){
			padded.Set(x, y, pixel)
		})
		leftPaddingReplicate(img, padded, p, func(x int, y int, pixel color.Color){
			padded.Set(x, y, pixel)
		})
		rightPaddingReplicate(img, padded, p, func(x int, y int, pixel color.Color){
			padded.Set(x, y, pixel)
		})
	case BorderReflect:
		topPaddingReflect(img, p, func(x int, y int, pixel color.Color){
			padded.Set(x, y, pixel)
		})
		bottomPaddingReflect(img, p, func(x int, y int, pixel color.Color){
			padded.Set(x, y, pixel)
		})
		leftPaddingReflect(img, padded, p, func(x int, y int, pixel color.Color){
			padded.Set(x, y, pixel)
		})
		rightPaddingReflect(img, padded, p, func(x int, y int, pixel color.Color){
			padded.Set(x, y, pixel)
		})
	default:
		return nil, errors.New("unknown border type")
	}
	return padded, nil
}

func PaddingRGBA(img *image.RGBA, kernelSize image.Point, anchor image.Point, border Border) (*image.RGBA, error) {
	originalSize := img.Bounds().Size()
	p, error := calculatePaddings(kernelSize, anchor)
	if error != nil {
		return nil, error
	}
	rect := getReactangleFromPaddings(p, originalSize)
	padded := image.NewRGBA(rect)

	for x := p.PaddingLeft; x < originalSize.X + p.PaddingLeft; x++ {
		for y := p.PaddingTop; y < originalSize.Y + p.PaddingTop; y++ {
			padded.Set(x, y, img.RGBAAt(x - p.PaddingLeft, y - p.PaddingTop))
		}
	}

	switch border {
	case BorderConstant:
		// do nothing
	case BorderReplicate:
		topPaddingReplicate(img, p, func(x int, y int, pixel color.Color){
			padded.Set(x, y, pixel)
		})
		bottomPaddingReplicate(img, p, func(x int, y int, pixel color.Color){
			padded.Set(x, y, pixel)
		})
		leftPaddingReplicate(img, padded, p, func(x int, y int, pixel color.Color){
			padded.Set(x, y, pixel)
		})
		rightPaddingReplicate(img, padded, p, func(x int, y int, pixel color.Color){
			padded.Set(x, y, pixel)
		})
	case BorderReflect:
		topPaddingReflect(img, p, func(x int, y int, pixel color.Color){
			padded.Set(x, y, pixel)
		})
		bottomPaddingReflect(img, p, func(x int, y int, pixel color.Color){
			padded.Set(x, y, pixel)
		})
		leftPaddingReflect(img, padded, p, func(x int, y int, pixel color.Color){
			padded.Set(x, y, pixel)
		})
		rightPaddingReflect(img, padded, p, func(x int, y int, pixel color.Color){
			padded.Set(x, y, pixel)
		})
	default:
		return nil, errors.New("unknown border type")
	}
	return padded, nil
}

// -------------------------------------------------------------------------------------------------------
func calculatePaddings(kernelSize image.Point, anchor image.Point) (Paddings, error) {
	var p Paddings
	if kernelSize.X < 0 || kernelSize.Y < 0 {
		return p, errors.New("negative size")
	}
	if anchor.X < 0 || anchor.Y < 0 {
		return p, errors.New("negative anchor value")
	}
	if anchor.X > kernelSize.X || anchor.Y > kernelSize.Y {
		return p, errors.New("anc" + "hor value outside of the kernel")
	}

	p = Paddings{PaddingLeft: anchor.X, PaddingRight: kernelSize.X - anchor.X, PaddingTop: anchor.Y, PaddingBottom: kernelSize.Y - anchor.Y}

	return p, nil
}

func getReactangleFromPaddings(p Paddings, imgSize image.Point,) image.Rectangle {
	x := p.PaddingLeft + p.PaddingRight + imgSize.X
	y := p.PaddingTop + p.PaddingBottom + imgSize.Y
	return image.Rect(0, 0, x, y)
}
