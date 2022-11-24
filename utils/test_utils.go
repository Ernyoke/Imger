package utils

/*
	Utility functions used for testing, do not use this in the production code.
*/
import (
	"fmt"
	"image"
	"math"
	"testing"
)

// CompareGrayImages Compares two Gray images and prints out if there is a difference between the pixels
func CompareGrayImages(t *testing.T, expected *image.Gray, actual *image.Gray) {
	expectedSize := expected.Bounds().Size()
	actualSize := actual.Bounds().Size()
	if !expectedSize.Eq(actualSize) {
		t.Fatalf("expected (size: %d %d) and actual (size: %d %d) have different sizes:", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	for x := 0; x < expected.Bounds().Size().X; x++ {
		for y := 0; y < expected.Bounds().Size().Y; y++ {
			c1 := expected.GrayAt(x, y)
			c2 := actual.GrayAt(x, y)
			if c1.Y != c2.Y {
				t.Errorf("Expected gray: %d - actual gray: %d at: %d %d", c1.Y, c2.Y, y, x)
			}
		}
	}
}

// CompareGrayImagesWithOffset Compares two Gray images within a given interval (pixel +/- offset) and prints out if there is a difference between the pixels
func CompareGrayImagesWithOffset(t *testing.T, expected *image.Gray, actual *image.Gray, offset uint16) {
	expectedSize := expected.Bounds().Size()
	actualSize := actual.Bounds().Size()
	if !expectedSize.Eq(actualSize) {
		t.Fatalf("expected (size: %d %d) and actual (size: %d %d) have different sizes:", expectedSize.X, expectedSize.Y, actualSize.X, actualSize.Y)
	}
	for x := 0; x < expected.Bounds().Size().X; x++ {
		for y := 0; y < expected.Bounds().Size().Y; y++ {
			c1 := expected.GrayAt(x, y)
			c2 := actual.GrayAt(x, y)
			if uint16(c1.Y) >= uint16(c2.Y)-offset && uint16(c1.Y) <= uint16(c2.Y)+offset {
				continue
			} else {
				t.Errorf("Expected gray: %d - actual gray: %d at: %d %d", c1.Y, c2.Y, y, x)
			}
		}
	}
}

// CompareRGBAImages Compares two RGBA images and prints out if there is a difference between the pixels
func CompareRGBAImages(t *testing.T, expected *image.RGBA, actual *image.RGBA) {
	if !expected.Bounds().Size().Eq(actual.Bounds().Size()) {
		t.Fatal("img1 and img2 have different sizes!")
	}
	for x := 0; x < expected.Bounds().Size().X; x++ {
		for y := 0; y < expected.Bounds().Size().Y; y++ {
			c1 := expected.RGBAAt(x, y)
			c2 := actual.RGBAAt(x, y)
			if c1.R != c2.R {
				t.Errorf("Expected red: %d - actual red: %d at: %d %d", c1.R, c2.R, y, x)
			}
			if c1.G != c2.G {
				t.Errorf("Expected green: %d - actual green: %d at: %d %d", c1.G, c2.G, y, x)
			}
			if c1.B != c2.B {
				t.Errorf("Expected blue: %d - actual blue: %d at: %d %d", c1.B, c2.B, y, x)
			}
			if c1.A != c2.A {
				t.Errorf("Expected alpha: %d - actual alpha: %d at: %d %d", c1.A, c2.A, y, x)
			}
		}
	}
}

// CompareRGBAImagesWithOffset Compares two RGBA images within a given interval (pixel +/- offset) and prints out if there is a difference between the pixels
func CompareRGBAImagesWithOffset(t *testing.T, expected *image.RGBA, actual *image.RGBA, offset uint16) {
	if !expected.Bounds().Size().Eq(actual.Bounds().Size()) {
		t.Fatal("img1 and img2 have different sizes!")
	}
	for x := 0; x < expected.Bounds().Size().X; x++ {
		for y := 0; y < expected.Bounds().Size().Y; y++ {
			c1 := expected.RGBAAt(x, y)
			c2 := actual.RGBAAt(x, y)
			if uint16(c1.R) < uint16(c2.R)-offset || uint16(c1.R) > uint16(c2.R)+offset {
				t.Errorf("Expected red: %d - actual red: %d at: %d %d", c1.R, c2.R, y, x)
			}
			if uint16(c1.G) < uint16(c2.G)-offset || uint16(c1.G) > uint16(c2.G)+offset {
				t.Errorf("Expected green: %d - actual green: %d at: %d %d", c1.G, c2.G, y, x)
			}
			if uint16(c1.B) < uint16(c2.B)-offset || uint16(c1.B) > uint16(c2.B)+offset {
				t.Errorf("Expected blue: %d - actual blue: %d at: %d %d", c1.B, c2.B, y, x)
			}
			if uint16(c1.A) < uint16(c2.A)-offset || uint16(c1.A) > uint16(c2.A)+offset {
				t.Errorf("Expected alpha: %d - actual alpha: %d at: %d %d", c1.A, c2.A, y, x)
			}
		}
	}
}

// PrintGray Print out gray image pixels to console
func PrintGray(t *testing.T, gray *image.Gray) {
	size := gray.Bounds().Size()
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			fmt.Printf("0x%x ", gray.GrayAt(x, y).Y)
		}
		fmt.Printf("\n")
	}
}

// PrintRGBA Print out gray image pixels to console
func PrintRGBA(t *testing.T, rgba *image.RGBA) {
	size := rgba.Bounds().Size()
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			fmt.Printf("0x%x ", rgba.RGBAAt(x, y))
		}
		fmt.Printf("\n")
	}
}

// IsEqualFloat64 Compares 2 float values and returns true if they are inside of the interval of [-eps, +eps]
func IsEqualFloat64(x float64, y float64) bool {
	eps := 0.0000001
	return math.Abs(x-y) <= eps
}
