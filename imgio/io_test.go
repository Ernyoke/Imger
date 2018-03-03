package imgio

import (
	"image"
	"testing"
)

// -----------------------------Acceptance tests------------------------------------
func Test_ImreadGray(t *testing.T) {
	path := "../res/girl.jpg"
	img, err := ImreadGray(path)
	if err != nil {
		t.Fatal("Could not read file!")
	}
	expectedSize := image.Point{X: 403, Y: 403}
	actualSize := img.Bounds().Size()
	if !(expectedSize.X == actualSize.X && expectedSize.Y == actualSize.Y) {
		t.Errorf("Expected size [%d %d] does not equal actual size [%d %d]", expectedSize.X, expectedSize.Y,
			actualSize.X, actualSize.Y)
	}
}

func Test_ImreadGray16(t *testing.T) {
	path := "../res/girl.jpg"
	img, err := ImreadGray16(path)
	if err != nil {
		t.Fatal("Could not read file!")
	}
	expectedSize := image.Point{X: 403, Y: 403}
	actualSize := img.Bounds().Size()
	if !(expectedSize.X == actualSize.X && expectedSize.Y == actualSize.Y) {
		t.Errorf("Expected size [%d %d] does not equal actual size [%d %d]", expectedSize.X, expectedSize.Y,
			actualSize.X, actualSize.Y)
	}
}

func Test_ImreadRGBA(t *testing.T) {
	path := "../res/girl.jpg"
	img, err := ImreadRGBA(path)
	if err != nil {
		t.Fatal("Could not read file!")
	}
	expectedSize := image.Point{X: 403, Y: 403}
	actualSize := img.Bounds().Size()
	if !(expectedSize.X == actualSize.X && expectedSize.Y == actualSize.Y) {
		t.Errorf("Expected size [%d %d] does not equal actual size [%d %d]", expectedSize.X, expectedSize.Y,
			actualSize.X, actualSize.Y)
	}
}

func Test_ImreadRGBA64(t *testing.T) {
	path := "../res/girl.jpg"
	img, err := ImreadRGBA64(path)
	if err != nil {
		t.Fatal("Could not read file!")
	}
	expectedSize := image.Point{X: 403, Y: 403}
	actualSize := img.Bounds().Size()
	if !(expectedSize.X == actualSize.X && expectedSize.Y == actualSize.Y) {
		t.Errorf("Expected size [%d %d] does not equal actual size [%d %d]", expectedSize.X, expectedSize.Y,
			actualSize.X, actualSize.Y)
	}
}

func Test_Imread_InexistentFile(t *testing.T) {
	path := "../res/inexistent.jpg"
	_, err := ImreadRGBA64(path)
	if err != nil {
		// ok
		return
	}
	t.Fatal("Should not reach this point!")
}

func Test_ImwriteJPG(t *testing.T) {
	path := "../res/girl.jpg"
	img, err := ImreadRGBA(path)
	if err != nil {
		t.Fatal("Could not read file!")
	}
	outPath := "../res/io/outputJPG.jpg"
	errOut := Imwrite(img, outPath)
	if errOut != nil {
		t.Fatalf("Could not write to this location: %s! Error: %s", outPath, errOut)
	}
}

func Test_ImwritePNG(t *testing.T) {
	path := "../res/girl.jpg"
	img, err := ImreadRGBA(path)
	if err != nil {
		t.Fatal("Could not read file!")
	}
	outPath := "../res/io/outputPNG.png"
	errOut := Imwrite(img, outPath)
	if errOut != nil {
		t.Fatalf("Could not write to this location: %s! Error: %s", outPath, errOut)
	}
}

func Test_Imwrite_InvalidExtension(t *testing.T) {
	path := "../res/girl.jpg"
	img, err := ImreadRGBA(path)
	if err != nil {
		t.Fatal("Could not read file!")
	}
	outPath := "../res/io/invalid.xxx"
	errOut := Imwrite(img, outPath)
	if errOut != nil {
		// ok
		return
	}
	t.Fatal("Should not reach this point!")
}
