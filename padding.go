package imger

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

type paddings struct {
	paddingLeft int
	paddingRight int
	paddingTop int
	paddingBottom int
}

func topPaddingReplicate(img image.Image, p paddings, setPixel func(int, int, color.Color)) {
	originalSize := img.Bounds().Size()
	for x := p.paddingLeft; x < originalSize.X + p.paddingLeft; x++ {
		firstPixel := img.At(x - p.paddingLeft, p.paddingTop)
		for y := 0; y < p.paddingTop; y++ {
			setPixel(x, y, firstPixel)
		}
	}
}

func bottomPaddingReplicate(img image.Image, p paddings, setPixel func(int, int, color.Color)) {
	originalSize := img.Bounds().Size()
	for x := p.paddingLeft; x < originalSize.X + p.paddingLeft; x++ {
		lastPixel := img.At(x - p.paddingLeft, originalSize.Y - 1)
		for y := p.paddingTop + originalSize.Y; y < originalSize.Y + p.paddingTop + p.paddingBottom; y++ {
			setPixel(x, y, lastPixel)
		}
	}
}

func leftPaddingReplicate(img image.Image, padded image.Image, p paddings, setPixel func(int, int, color.Color)) {
	originalSize := img.Bounds().Size()
	for y := 0; y < originalSize.Y + p.paddingBottom + p.paddingTop; y++ {
		firstPixel := padded.At(p.paddingLeft, y)
		for x := 0; x < p.paddingLeft; x++ {
			setPixel(x, y, firstPixel)
		}
	}
}

func rightPaddingReplicate(img image.Image, padded image.Image, p paddings, setPixel func(int, int, color.Color)) {
	originalSize := img.Bounds().Size()
	for y := 0; y < originalSize.Y + p.paddingBottom + p.paddingTop; y++ {
		lastPixel := padded.At(originalSize.X + p.paddingLeft - 1, y)
		for x := originalSize.X + p.paddingLeft; x < originalSize.X + p.paddingLeft + p.paddingRight; x++ {
			setPixel(x, y, lastPixel)
		}
	}
}

func topPaddingReflect(img image.Image, p paddings, setPixel func(int, int, color.Color)) {
	originalSize := img.Bounds().Size()
	for x := p.paddingLeft; x < originalSize.X + p.paddingLeft; x++ {
		for y := 0; y < p.paddingTop; y++ {
			pixel := img.At(x - p.paddingLeft, p.paddingTop - y)
			setPixel(x, y, pixel)
		}
	}
}

func bottomPaddingReflect(img image.Image, p paddings, setPixel func(int, int, color.Color)) {
	originalSize := img.Bounds().Size()
	for x := p.paddingLeft; x < originalSize.X + p.paddingLeft; x++ {
		for y := p.paddingTop + originalSize.Y; y < originalSize.Y + p.paddingTop + p.paddingBottom; y++ {
			pixel := img.At(x - p.paddingLeft, originalSize.Y - (y - p.paddingTop - originalSize.Y) - 1)
			setPixel(x, y, pixel)
		}
	}
}

func leftPaddingReflect(img image.Image, padded image.Image, p paddings, setPixel func(int, int, color.Color)) {
	originalSize := img.Bounds().Size()
	for y := 0; y < originalSize.Y + p.paddingBottom + p.paddingTop; y++ {
		for x := 0; x < p.paddingLeft; x++ {
			pixel := padded.At(2 * p.paddingLeft - x, y)
			setPixel(x, y, pixel)
		}
	}
}

func rightPaddingReflect(img image.Image, padded image.Image, p paddings, setPixel func(int, int, color.Color)) {
	originalSize := img.Bounds().Size()
	for y := 0; y < originalSize.Y + p.paddingBottom + p.paddingTop; y++ {
		for x := originalSize.X + p.paddingLeft; x < originalSize.X + p.paddingLeft + p.paddingRight; x++ {
			pixel := padded.At(originalSize.X + p.paddingRight - (x - originalSize.X - p.paddingLeft), y)
			setPixel(x, y, pixel)
		}
	}
}


func PaddingGray(img *image.Gray, kernelSize image.Point, anchor image.Point, border Border) (image.Image, error) {
	originalSize := img.Bounds().Size()
	rect, p, error := getPaddings(originalSize, kernelSize, anchor)
	if error != nil {
		return nil, error
	}
	padded := image.NewGray(rect)

	for x := p.paddingLeft; x < originalSize.X + p.paddingLeft; x++ {
		for y := p.paddingTop; y < originalSize.Y + p.paddingTop; y++ {
			padded.Set(x, y, img.At(x - p.paddingLeft, y - p.paddingTop))
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

func getPaddings(imgSize image.Point, kernelSize image.Point, anchor image.Point) (image.Rectangle, paddings, error) {
	var p paddings
	var rect image.Rectangle
	if kernelSize.X < 0 || kernelSize.Y < 0 {
		return rect, p, errors.New("negative size")
	}
	if anchor.X < 0 || anchor.Y < 0 {
		return rect, p, errors.New("negative anchor value")
	}
	if anchor.X > kernelSize.X || anchor.Y > kernelSize.Y {
		return rect, p, errors.New("anchor value outside of the kernel")
	}

	p = paddings{paddingLeft: anchor.X, paddingRight: kernelSize.X - anchor.X, paddingTop: anchor.Y, paddingBottom: kernelSize.Y - anchor.Y}

	imgSize.X += p.paddingLeft + p.paddingRight
	imgSize.Y += p.paddingTop + p.paddingBottom

	rect = image.Rect(0, 0, imgSize.X, imgSize.Y)

	return rect, p, nil
}

func PaddingRGBA(img *image.RGBA, kernelSize image.Point, anchor image.Point, border Border) (image.Image, error) {
	originalSize := img.Bounds().Size()
	rect, p, error := getPaddings(originalSize, kernelSize, anchor)
	if error != nil {
		return nil, error
	}
	padded := image.NewRGBA(rect)

	for x := p.paddingLeft; x < originalSize.X + p.paddingLeft; x++ {
		for y := p.paddingTop; y < originalSize.Y + p.paddingTop; y++ {
			padded.Set(x, y, img.At(x - p.paddingLeft, y - p.paddingTop))
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
