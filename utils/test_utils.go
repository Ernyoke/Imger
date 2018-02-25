package utils

/*
	Utility functions used for testing, do not use this in the production code.
*/
import (
	"fmt"
	"image"
	"testing"
)

// Compares to Gray images and prints out if there is a difference between the pixels
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

// Compares to RGBA images and prints out if there is a difference between the pixels
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
				t.Errorf("Expected red: %d - actual green: %d at: %d %d", c1.G, c2.G, y, x)
			}
			if c1.B != c2.B {
				t.Errorf("Expected red: %d - actual blue: %d at: %d %d", c1.B, c2.G, y, x)
			}
			if c1.A != c2.A {
				t.Errorf("Expected red: %d - actual alpha: %d at: %d %d", c1.A, c2.G, y, x)
			}
		}
	}
}

// Print out gray image pixels to console
func PrintGray(t *testing.T, gray *image.Gray) {
	size := gray.Bounds().Size()
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			fmt.Printf("0x%x ", gray.GrayAt(x, y).Y)
		}
		fmt.Printf("\n")
	}
}
