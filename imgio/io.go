package imgio

import (
	"errors"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// Reads and decodes image from a given path. Supported extensions are: jpg, jpeg, png
func decode(path string) (image.Image, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	extension := strings.ToLower(filepath.Ext(path))
	switch extension {
	case ".jpg":
		fallthrough
	case ".jpeg":
		return jpeg.Decode(file)
	case ".png":
		return png.Decode(file)
	}
	return nil, errors.New("unsupported extension")
}

// Encodes and writes image to the given path
func encode(img image.Image, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	extension := strings.ToLower(filepath.Ext(path))
	switch extension {
	case ".jpg":
		fallthrough
	case ".jpeg":
		return jpeg.Encode(file, img, nil)
	case ".png":
		return png.Encode(file, img)
	}
	return errors.New("unsupported extension")
}

// ImreadGray reads the image from the given path and return a grayscale image. Returns an error if the path is not
// readable or the specified resource does not exist.
func ImreadGray(path string) (*image.Gray, error) {
	img, err := decode(path)
	if err != nil {
		return nil, err
	}
	bounds := img.Bounds()
	gray := image.NewGray(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(gray, bounds, img, bounds.Min, draw.Src)
	return gray, nil
}

// ImreadGray16 reads the image from the given path and return a grayscale16 image. Returns an error if the path is not
// readable or the specified resource does not exist.
func ImreadGray16(path string) (*image.Gray16, error) {
	img, err := decode(path)
	if err != nil {
		return nil, err
	}
	bounds := img.Bounds()
	gray16 := image.NewGray16(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(gray16, bounds, img, bounds.Min, draw.Src)
	return gray16, nil
}

// ImreadRGBA reads the image from the given path and return a RGBA image. Returns an error if the path is not readable
// or the specified resource does not exist.
func ImreadRGBA(path string) (*image.RGBA, error) {
	img, err := decode(path)
	if err != nil {
		return nil, err
	}
	bounds := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)
	return rgba, nil
}

// ImreadRGBA64 reads the image from the given path and return a RGBA64 image.
// Returns an error if the path is not readable or the specified resource does not exist.
func ImreadRGBA64(path string) (*image.RGBA64, error) {
	img, err := decode(path)
	if err != nil {
		return nil, err
	}
	bounds := img.Bounds()
	rgba64 := image.NewRGBA64(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba64, bounds, img, bounds.Min, draw.Src)
	return rgba64, nil
}

// Imwrite saves the image under the location specified by the "path" string. Returns an error if the location is
// not writable.
func Imwrite(img image.Image, path string) error {
	return encode(img, path)
}
