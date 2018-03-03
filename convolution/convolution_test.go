package convolution

import (
	"github.com/Ernyoke/Imger/padding"
	"github.com/Ernyoke/Imger/utils"
	"image"
	"testing"
)

// ---------------------------------Unit tests------------------------------------
func TestGrayScale(t *testing.T) {
	var gray image.Gray
	gray = image.Gray{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3,
		Pix: []uint8{
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
		},
	}
	var expected image.Gray
	expected = image.Gray{
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
	utils.ForEchPixel(size, func(x, y int) {
		pExp := expected.GrayAt(x, y).Y
		pRes := conv.GrayAt(x, y).Y
		if pExp != pRes {
			t.Errorf("Expected: %d - Actual: %d at %d, %d", pExp, pRes, x, y)
		}
	})
}

// -------------------------------------------------------------------------------
