package histogram

import (
	"github.com/ernyoke/imger/imgio"
	"image"
	"testing"
)

// --------------------------------Unit tests---------------------------------------

func Test_Histogram_GrayScale(t *testing.T) {
	gray := image.Gray{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3,
		Pix: []uint8{
			0x01, 0x02, 0x80, 0x80, 0x02, 0xFF, 0xFF, 0xFF, 0xFF,
		},
	}
	hist := HistogramGray(&gray)
	for i, h := range hist {
		switch i {
		case 0x01:
			expected := uint64(1)
			if h != expected {
				t.Errorf("Histogram value for %d should be %d!", h, expected)
			}
		case 0x02:
			expected := uint64(2)
			if h != expected {
				t.Errorf("Histogram value for %d should be %d!", h, expected)
			}
		case 0x80:
			expected := uint64(2)
			if h != expected {
				t.Errorf("Histogram value for %d should be %d!", h, expected)
			}
		case 0xFF:
			expected := uint64(4)
			if h != expected {
				t.Errorf("Histogram value for %d should be %d!", h, expected)
			}
		default:
			expected := uint64(0)
			if h != expected {
				t.Errorf("Histogram value for %d should be %d!", h, expected)
			}
		}
	}
}

func Test_Histogram_RGBA(t *testing.T) {
	rgba := image.RGBA{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3 * 4,
		Pix: []uint8{
			0x01, 0x01, 0x01, 0xFF, 0x02, 0x02, 0x02, 0xFF, 0x80, 0x80, 0x80, 0xFF,
			0x80, 0x80, 0x80, 0xFF, 0x02, 0x02, 0x02, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
		},
	}
	hist := HistogramRGBA(&rgba)
	for _, hc := range hist {
		for i, h := range hc {
			switch i {
			case 0x01:
				expected := uint64(1)
				if h != expected {
					t.Errorf("Histogram value for %d should be %d!", h, expected)
				}
			case 0x02:
				expected := uint64(2)
				if h != expected {
					t.Errorf("Histogram value for %d should be %d!", h, expected)
				}
			case 0x80:
				expected := uint64(2)
				if h != expected {
					t.Errorf("Histogram value for %d should be %d!", h, expected)
				}
			case 0xFF:
				expected := uint64(4)
				if h != expected {
					t.Errorf("Histogram value for %d should be %d!", h, expected)
				}
			default:
				expected := uint64(0)
				if h != expected {
					t.Errorf("Histogram value for %d should be %d!", h, expected)
				}
			}
		}
	}
}

func Test_Histogram_RGBARed(t *testing.T) {
	rgba := image.RGBA{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3 * 4,
		Pix: []uint8{
			0x01, 0x01, 0x01, 0xFF, 0x02, 0x66, 0x06, 0xFF, 0x80, 0x80, 0x80, 0xFF,
			0x80, 0x80, 0x80, 0xFF, 0x02, 0x12, 0x02, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
		},
	}
	hist := HistogramRGBARed(&rgba)
	for i, h := range hist {
		switch i {
		case 0x01:
			expected := uint64(1)
			if h != expected {
				t.Errorf("Histogram value for %d should be %d!", h, expected)
			}
		case 0x02:
			expected := uint64(2)
			if h != expected {
				t.Errorf("Histogram value for %d should be %d!", h, expected)
			}
		case 0x80:
			expected := uint64(2)
			if h != expected {
				t.Errorf("Histogram value for %d should be %d!", h, expected)
			}
		case 0xFF:
			expected := uint64(4)
			if h != expected {
				t.Errorf("Histogram value for %d should be %d!", h, expected)
			}
		default:
			expected := uint64(0)
			if h != expected {
				t.Errorf("Histogram value for %d should be %d!", h, expected)
			}
		}
	}
}

func Test_Histogram_RGBAGreen(t *testing.T) {
	rgba := image.RGBA{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3 * 4,
		Pix: []uint8{
			0x45, 0x01, 0x08, 0xFF, 0x02, 0x02, 0x02, 0xFF, 0x80, 0x80, 0x80, 0xFF,
			0x56, 0x80, 0x80, 0xFF, 0x02, 0x02, 0x02, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
		},
	}
	hist := HistogramRGBAGreen(&rgba)
	for i, h := range hist {
		switch i {
		case 0x01:
			expected := uint64(1)
			if h != expected {
				t.Errorf("Histogram value for %d should be %d!", h, expected)
			}
		case 0x02:
			expected := uint64(2)
			if h != expected {
				t.Errorf("Histogram value for %d should be %d!", h, expected)
			}
		case 0x80:
			expected := uint64(2)
			if h != expected {
				t.Errorf("Histogram value for %d should be %d!", h, expected)
			}
		case 0xFF:
			expected := uint64(4)
			if h != expected {
				t.Errorf("Histogram value for %d should be %d!", h, expected)
			}
		default:
			expected := uint64(0)
			if h != expected {
				t.Errorf("Histogram value for %d should be %d!", h, expected)
			}
		}
	}
}

func Test_Histogram_RGBABlue(t *testing.T) {
	rgba := image.RGBA{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3 * 4,
		Pix: []uint8{
			0x01, 0x11, 0x01, 0xFF, 0x02, 0x02, 0x02, 0xFF, 0x80, 0x80, 0x80, 0xFF,
			0x80, 0x81, 0x80, 0xFF, 0x02, 0x02, 0x02, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFD, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
		},
	}
	hist := HistogramRGBABlue(&rgba)
	for i, h := range hist {
		switch i {
		case 0x01:
			expected := uint64(1)
			if h != expected {
				t.Errorf("Histogram value for %d should be %d!", h, expected)
			}
		case 0x02:
			expected := uint64(2)
			if h != expected {
				t.Errorf("Histogram value for %d should be %d!", h, expected)
			}
		case 0x80:
			expected := uint64(2)
			if h != expected {
				t.Errorf("Histogram value for %d should be %d!", h, expected)
			}
		case 0xFF:
			expected := uint64(4)
			if h != expected {
				t.Errorf("Histogram value for %d should be %d!", h, expected)
			}
		default:
			expected := uint64(0)
			if h != expected {
				t.Errorf("Histogram value for %d should be %d!", h, expected)
			}
		}
	}
}

// -----------------------------Acceptance tests------------------------------------
func setupTestCaseGray(t *testing.T) *image.Gray {
	path := "../res/girl.jpg"
	img, err := imgio.ImreadGray(path)
	if err != nil {
		t.Errorf("Could not read image from path: %s", path)
	}
	return img
}

func setupTestCaseRGBA(t *testing.T) *image.RGBA {
	path := "../res/girl.jpg"
	img, err := imgio.ImreadRGBA(path)
	if err != nil {
		t.Errorf("Could not read image from path: %s", path)
	}
	return img
}

func tearDownTestCase(t *testing.T, img image.Image, path string) {
	err := imgio.Imwrite(img, path)
	if err != nil {
		t.Errorf("Could not write image to path: %s", path)
	}
}

func Test_Acceptance_DrawHistogram_GrayScale(t *testing.T) {
	gray := setupTestCaseGray(t)
	expectedSize := image.Point{X: 512, Y: 600}
	hist := DrawHistogramGray(gray, expectedSize)
	actualSize := hist.Bounds().Size()
	if actualSize.X != expectedSize.X && actualSize.Y != expectedSize.Y {
		t.Fatalf("Size of expected [%d %d] does not match size of actual [%d %d]", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, hist, "../res/histogram/gray.jpg")
}

func Test_Acceptance_DrawHistogram_RGBA(t *testing.T) {
	rgba := setupTestCaseRGBA(t)
	expectedSize := image.Point{X: 512, Y: 600}
	hist := DrawHistogramRGBA(rgba, expectedSize)
	actualSize := hist.Bounds().Size()
	if actualSize.X != expectedSize.X && actualSize.Y != expectedSize.Y {
		t.Fatalf("Size of expected [%d %d] does not match size of actual [%d %d]", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	tearDownTestCase(t, hist, "../res/histogram/rgba.jpg")
}
