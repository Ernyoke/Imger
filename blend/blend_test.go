package blend

import (
	"github.com/ernyoke/imger/utils"
	"image"
	"testing"
)

func Test_AddScalarToGray(t *testing.T) {
	input := image.Gray{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3,
		Pix: []uint8{
			0xFF, 0x80, 0x56,
			0x56, 0x80, 0xFD,
			0x00, 0x1A, 0xBB,
		},
	}
	expected := &image.Gray{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3,
		Pix: []uint8{
			0xFF, 0x8A, 0x60,
			0x60, 0x8A, 0xFF,
			0x0A, 0x24, 0xC5,
		},
	}
	result := AddScalarToGray(&input, 10)
	utils.CompareGrayImages(t, expected, result)
}

func Test_AddGray(t *testing.T) {
	input1 := image.Gray{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3,
		Pix: []uint8{
			0xFF, 0x80, 0x56,
			0xFD, 0xFD, 0xFD,
			0x00, 0x00, 0xBB,
		},
	}
	input2 := image.Gray{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3,
		Pix: []uint8{
			0xFF, 0x0A, 0x56,
			0x01, 0x02, 0x03,
			0x00, 0xFF, 0xBB,
		},
	}
	expected := &image.Gray{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3,
		Pix: []uint8{
			0xFF, 0x8A, 0xAC,
			0xFE, 0xFF, 0xFF,
			0x00, 0xFF, 0xFF,
		},
	}
	result, err := AddGray(&input1, &input2)
	if err != nil {
		t.Fatalf("Error should not be returned. Error value: %s", err)
	}
	utils.CompareGrayImages(t, expected, result)
}

func Test_AddGrayWeighted(t *testing.T) {
	input1 := image.Gray{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3,
		Pix: []uint8{
			0xFF, 0x80, 0x56,
			0xFD, 0xFD, 0xFD,
			0x00, 0x00, 0xBB,
		},
	}
	input2 := image.Gray{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3,
		Pix: []uint8{
			0xFF, 0x0A, 0x56,
			0x01, 0x02, 0x03,
			0x00, 0xFF, 0xBB,
		},
	}
	expected := &image.Gray{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3,
		Pix: []uint8{
			0xFF, 0x45, 0x56,
			0x7F, 0x7F, 0x80,
			0x00, 0x7F, 0xBB,
		},
	}
	result, err := AddGrayWeighted(&input1, 0.5, &input2, 0.5)
	if err != nil {
		t.Fatalf("Error should not be returned. Error value: %s", err)
	}
	utils.CompareGrayImages(t, expected, result)
}

func Test_AddGrayWeightedError(t *testing.T) {
	input1 := image.Gray{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3,
		Pix: []uint8{
			0xFF, 0x80, 0x56,
			0xFD, 0xFD, 0xFD,
			0x00, 0x00, 0xBB,
		},
	}
	input2 := image.Gray{
		Rect:   image.Rect(0, 0, 3, 2),
		Stride: 3,
		Pix: []uint8{
			0xFF, 0x0A, 0x56,
			0x01, 0x02, 0x03,
		},
	}
	_, err := AddGrayWeighted(&input1, 0.5, &input2, 0.5)
	if err != nil {
		if err.Error() != "the size of the two image does not match" {
			t.Fatalf("Invalid error message!")
		}
	} else {
		t.Fatalf("Should not reach this point")
	}
}
