package convolution

import (
	"testing"
	"image"
	"github.com/ernyoke/imgur/padding"
)

func TestGrayScale(t *testing.T) {
	var gray image.Gray
	gray = image.Gray {
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3,
		Pix: []uint8{
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
		},
	}
	var expected image.Gray
	expected = image.Gray {
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3,
		Pix: []uint8{
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
		},
	}
	var kernel Kernel
	kernel = Kernel{[][]float64{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}, 3, 3}
	conv, _ := ConvolveGray(&gray, &kernel, image.Point{1, 1}, padding.BorderConstant)
	size := conv.Bounds().Size()
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			rRes, gRes, bRes, aRes := conv.At(x, y).RGBA()
			rExp, gExp, bExp, aExp := expected.At(x, y).RGBA()
			if rRes == rExp && gRes == gExp && bRes == bExp && aRes == aExp {
				continue
			} else {
				t.Errorf("Expected: %d %d %d %d - Actual: %d %d %d %d at %d, %d", rExp, gExp, bExp, aExp, rRes, gRes, bRes, aRes, x, y)
			}
		}
	}
}
