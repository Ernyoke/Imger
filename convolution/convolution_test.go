package convolution

import (
	"github.com/ernyoke/imger/padding"
	"github.com/ernyoke/imger/utils"
	"image"
	"testing"
)

// ---------------------------------Unit tests------------------------------------
func Test_ConvolveGray_0Kernel(t *testing.T) {
	gray := image.Gray{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3,
		Pix: []uint8{
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
		},
	}
	expected := image.Gray{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3,
		Pix: []uint8{
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		},
	}
	kernel := Kernel{[][]float64{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}, 3, 3}
	conv, _ := ConvolveGray(&gray, &kernel, image.Point{X: 1, Y: 1}, padding.BorderConstant)
	size := conv.Bounds().Size()
	utils.ForEachPixel(size, func(x, y int) {
		pExp := expected.GrayAt(x, y).Y
		pRes := conv.GrayAt(x, y).Y
		if pExp != pRes {
			t.Errorf("Expected: %d - Actual: %d at %d, %d", pExp, pRes, x, y)
		}
	})
}

func Test_ConvolveGray_1Kernel(t *testing.T) {
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
	conv, _ := ConvolveGray(&gray, &kernel, image.Point{X: 1, Y: 1}, padding.BorderConstant)
	size := conv.Bounds().Size()
	utils.ForEachPixel(size, func(x, y int) {
		pExp := expected.GrayAt(x, y).Y
		pRes := conv.GrayAt(x, y).Y
		if pExp != pRes {
			t.Errorf("Expected: %d - Actual: %d at %d, %d", pExp, pRes, x, y)
		}
	})
}

func Test_ConvoleGray_11Kernel(t *testing.T) {
	gray := image.Gray{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3,
		Pix: []uint8{
			0x80, 0xFF, 0xFF,
			0xFF, 0x40, 0x40,
			0xFF, 0xFF, 0x40,
		},
	}
	expected := image.Gray{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3,
		Pix: []uint8{
			0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0x80,
			0xFF, 0xFF, 0x40,
		},
	}
	var kernel Kernel
	kernel = Kernel{[][]float64{
		{0, 0, 0},
		{0, 1, 1},
		{0, 0, 0},
	}, 3, 3}
	conv, _ := ConvolveGray(&gray, &kernel, image.Point{X: 1, Y: 1}, padding.BorderConstant)
	size := conv.Bounds().Size()
	utils.ForEachPixel(size, func(x, y int) {
		pExp := expected.GrayAt(x, y)
		pRes := conv.GrayAt(x, y)
		if pExp != pRes {
			t.Errorf("Expected: %d - Actual: %d at %d, %d", pExp, pRes, x, y)
		}
	})
}

func Test_ConvoleRGBA_1Kernel(t *testing.T) {
	rgba := image.RGBA{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3 * 4,
		Pix: []uint8{
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0x01, 0x02, 0x03, 0x04, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
		},
	}
	expected := image.RGBA{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3 * 4,
		Pix: []uint8{
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0x01, 0x02, 0x03, 0x04, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
		},
	}
	var kernel Kernel
	kernel = Kernel{[][]float64{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}, 3, 3}
	conv, _ := ConvolveRGBA(&rgba, &kernel, image.Point{X: 1, Y: 1}, padding.BorderConstant)
	size := conv.Bounds().Size()
	utils.ForEachPixel(size, func(x, y int) {
		pExp := expected.RGBAAt(x, y)
		pRes := conv.RGBAAt(x, y)
		if pExp != pRes {
			t.Errorf("Expected: %d - Actual: %d at %d, %d", pExp, pRes, x, y)
		}
	})
}

func Test_ConvoleRGBA_11Kernel(t *testing.T) {
	rgba := image.RGBA{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3 * 4,
		Pix: []uint8{
			0x80, 0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x40, 0x40, 0x40, 0x40,
		},
	}
	expected := image.RGBA{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3 * 4,
		Pix: []uint8{
			0xFF, 0xFF, 0xFF, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x40, 0x80, 0x80, 0x80, 0x40,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x40, 0x40, 0x40, 0x40,
		},
	}
	var kernel Kernel
	kernel = Kernel{[][]float64{
		{0, 0, 0},
		{0, 1, 1},
		{0, 0, 0},
	}, 3, 3}
	conv, _ := ConvolveRGBA(&rgba, &kernel, image.Point{X: 1, Y: 1}, padding.BorderConstant)
	size := conv.Bounds().Size()
	utils.ForEachPixel(size, func(x, y int) {
		pExp := expected.RGBAAt(x, y)
		pRes := conv.RGBAAt(x, y)
		if pExp != pRes {
			t.Errorf("Expected: %d - Actual: %d at %d, %d", pExp, pRes, x, y)
		}
	})
}

// -------------------------------------------------------------------------------
