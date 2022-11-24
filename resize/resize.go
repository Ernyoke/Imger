package resize

import (
	"errors"
	"github.com/ernyoke/imger/utils"
	"image"
	"image/color"
	"math"
)

// Interpolation method types
type Interpolation int

const (
	// InterNearest - takes the nearest pixel.
	InterNearest Interpolation = iota
	// InterLinear - Linear interpolation between two pixels. More info: https://en.wikipedia.org/wiki/Linear_interpolation
	InterLinear
	// InterCatmullRom - Catmull-Rom resampling. More info: https://en.wikipedia.org/wiki/Centripetal_Catmull%E2%80%93Rom_spline
	InterCatmullRom
	// InterLanczos - Lanczos resampling. More info: https://en.wikipedia.org/wiki/Lanczos_resampling
	InterLanczos
)

func resizeNearestGray(img *image.Gray, fx float64, fy float64) (*image.Gray, error) {
	oldSize := img.Bounds().Size()
	newSize := image.Point{X: int(float64(oldSize.X) * fx), Y: int(float64(oldSize.Y) * fy)}
	newImg := image.NewGray(image.Rect(0, 0, newSize.X, newSize.Y))
	utils.ParallelForEachPixel(newSize, func(x int, y int) {
		oldXTemp := float64(x) / fx
		var oldX int
		if fraction := oldXTemp - float64(int(oldXTemp)); fraction >= 0.5 {
			oldX = int(oldXTemp + 1)
		} else {
			oldX = int(oldXTemp)
		}
		oldYTemp := float64(y) / fy
		var oldY int
		if fraction := oldYTemp - float64(int(oldYTemp)); fraction >= 0.5 {
			oldY = int(oldYTemp + 1)
		} else {
			oldY = int(oldYTemp)
		}
		newImg.SetGray(x, y, img.GrayAt(oldX, oldY))
	})
	return newImg, nil
}

func resizeLinearGray(img *image.Gray, fx float64, fy float64) (*image.Gray, error) {
	res, err := resizeHorizontalGray(img, fx, NewLinear())
	if err != nil {
		return nil, err
	}
	res, err = resizeVerticalGray(res, fy, NewLinear())
	if err != nil {
		return nil, err
	}

	return res, nil
}

func resizeCatmullRomGray(img *image.Gray, fx float64, fy float64) (*image.Gray, error) {
	res, err := resizeHorizontalGray(img, fx, NewCatmullRom())
	if err != nil {
		return nil, err
	}
	res, err = resizeVerticalGray(res, fy, NewCatmullRom())
	if err != nil {
		return nil, err
	}

	return res, nil
}

func resizeLanczosGray(img *image.Gray, fx float64, fy float64) (*image.Gray, error) {
	res, err := resizeHorizontalGray(img, fx, NewLanczos())
	if err != nil {
		return nil, err
	}
	res, err = resizeVerticalGray(res, fy, NewLanczos())
	if err != nil {
		return nil, err
	}

	return res, nil
}

func resizeHorizontalGray(img *image.Gray, fx float64, filter Filter) (*image.Gray, error) {
	originalSize := img.Bounds().Size()
	newWidth := int(float64(originalSize.X) * fx)
	res := image.NewGray(image.Rect(0, 0, newWidth, originalSize.Y))
	dfx := 1 / fx

	radius := math.Ceil(fx * filter.GetS())
	for y := 0; y < originalSize.Y; y++ {
		for x := 0; x < newWidth; x++ {
			ix := (float64(x)+0.5)*dfx - 0.5
			start := utils.ClampInt(int(ix-radius+0.5), 0, originalSize.X)
			end := utils.ClampInt(int(ix+radius), 0, originalSize.X)
			var fPix float64
			var sum float64
			for i := start; i < end; i++ {
				filterValue := filter.Interpolate(float64(i)-ix) / fx
				pix := img.GrayAt(i, y)
				fPix += float64(pix.Y) * filterValue
				sum += filterValue
			}
			res.SetGray(x, y, color.Gray{uint8(utils.ClampF64(fPix/sum+0.5, 0, 255))})
		}
	}
	return res, nil
}

func resizeVerticalGray(img *image.Gray, fy float64, filter Filter) (*image.Gray, error) {
	originalSize := img.Bounds().Size()
	newHeight := int(float64(originalSize.Y) * fy)
	res := image.NewGray(image.Rect(0, 0, originalSize.X, newHeight))
	dfy := 1 / fy

	radius := math.Ceil(fy * filter.GetS())
	for y := 0; y < newHeight; y++ {
		iy := (float64(y)+0.5)*dfy - 0.5
		start := utils.ClampInt(int(iy-radius+0.5), 0, originalSize.Y)
		end := utils.ClampInt(int(iy+radius), 0, originalSize.Y)
		for x := 0; x < originalSize.X; x++ {
			var sum float64
			var fPix float64
			for i := start; i < end; i++ {
				filterValue := filter.Interpolate(float64(i)-iy) / fy
				pix := img.GrayAt(x, i)
				fPix += float64(pix.Y) * filterValue
				sum += filterValue
			}
			res.SetGray(x, y, color.Gray{uint8(utils.ClampF64(fPix/sum+0.5, 0, 255))})
		}
	}
	return res, nil
}

func resizeNearestRGBA(img *image.RGBA, fx float64, fy float64) (*image.RGBA, error) {
	oldSize := img.Bounds().Size()
	newSize := image.Point{X: int(float64(oldSize.X) * fx), Y: int(float64(oldSize.Y) * fy)}
	newImg := image.NewRGBA(image.Rect(0, 0, newSize.X, newSize.Y))
	utils.ParallelForEachPixel(newSize, func(x int, y int) {
		oldXTemp := float64(x) / fx
		var oldX int
		if fraction := oldXTemp - float64(int(oldXTemp)); fraction >= 0.5 {
			oldX = int(oldXTemp + 1)
		} else {
			oldX = int(oldXTemp)
		}
		oldYTemp := float64(y) / fy
		var oldY int
		if fraction := oldYTemp - float64(int(oldYTemp)); fraction >= 0.5 {
			oldY = int(oldYTemp + 1)
		} else {
			oldY = int(oldYTemp)
		}
		newImg.SetRGBA(x, y, img.RGBAAt(oldX, oldY))
	})
	return newImg, nil
}

func resizeLinearRGBA(img *image.RGBA, fx float64, fy float64) (*image.RGBA, error) {
	res, err := resizeHorizontalRGBA(img, fx, NewLinear())
	if err != nil {
		return nil, err
	}
	res, err = resizeVerticalRGBA(res, fy, NewLinear())
	if err != nil {
		return nil, err
	}

	return res, nil
}

func resizeCatmullRomRGBA(img *image.RGBA, fx float64, fy float64) (*image.RGBA, error) {
	res, err := resizeHorizontalRGBA(img, fx, NewCatmullRom())
	if err != nil {
		return nil, err
	}
	res, err = resizeVerticalRGBA(res, fy, NewCatmullRom())
	if err != nil {
		return nil, err
	}

	return res, nil
}

func resizeLanczosRGBA(img *image.RGBA, fx float64, fy float64) (*image.RGBA, error) {
	res, err := resizeHorizontalRGBA(img, fx, NewLanczos())
	if err != nil {
		return nil, err
	}
	res, err = resizeVerticalRGBA(res, fy, NewLanczos())
	if err != nil {
		return nil, err
	}

	return res, nil
}

func resizeHorizontalRGBA(img *image.RGBA, fx float64, filter Filter) (*image.RGBA, error) {
	originalSize := img.Bounds().Size()
	newWidth := int(float64(originalSize.X) * fx)
	res := image.NewRGBA(image.Rect(0, 0, newWidth, originalSize.Y))
	dfx := 1 / fx

	radius := math.Ceil(fx * filter.GetS())
	for y := 0; y < originalSize.Y; y++ {
		for x := 0; x < newWidth; x++ {
			ix := (float64(x)+0.5)*dfx - 0.5
			start := utils.ClampInt(int(ix-radius+0.5), 0, originalSize.X)
			end := utils.ClampInt(int(ix+radius), 0, originalSize.X)
			var fPixR float64
			var fPixG float64
			var fPixB float64
			var fPixA float64
			var sum float64
			for i := start; i < end; i++ {
				filterValue := filter.Interpolate(float64(i)-ix) / fx
				pix := img.RGBAAt(i, y)
				fPixR += float64(pix.R) * filterValue
				fPixG += float64(pix.G) * filterValue
				fPixB += float64(pix.B) * filterValue
				fPixA += float64(pix.A) * filterValue
				sum += filterValue
			}
			res.SetRGBA(x, y, color.RGBA{R: uint8(utils.ClampF64(fPixR/sum+0.5, 0, 255)),
				G: uint8(utils.ClampF64(fPixG/sum+0.5, 0, 255)),
				B: uint8(utils.ClampF64(fPixB/sum+0.5, 0, 255)),
				A: uint8(utils.ClampF64(fPixA/sum+0.5, 0, 255))})
		}
	}
	return res, nil
}

func resizeVerticalRGBA(img *image.RGBA, fy float64, filter Filter) (*image.RGBA, error) {
	originalSize := img.Bounds().Size()
	newHeight := int(float64(originalSize.Y) * fy)
	res := image.NewRGBA(image.Rect(0, 0, originalSize.X, newHeight))
	dfy := 1 / fy

	radius := math.Ceil(fy * filter.GetS())
	for y := 0; y < newHeight; y++ {
		iy := (float64(y)+0.5)*dfy - 0.5
		start := utils.ClampInt(int(iy-radius+0.5), 0, originalSize.Y)
		end := utils.ClampInt(int(iy+radius), 0, originalSize.Y)
		for x := 0; x < originalSize.X; x++ {
			var fPixR float64
			var fPixG float64
			var fPixB float64
			var fPixA float64
			var sum float64
			for i := start; i < end; i++ {
				filterValue := filter.Interpolate(float64(i)-iy) / fy
				pix := img.RGBAAt(x, i)
				fPixR += float64(pix.R) * filterValue
				fPixG += float64(pix.G) * filterValue
				fPixB += float64(pix.B) * filterValue
				fPixA += float64(pix.A) * filterValue
				sum += filterValue
			}
			res.SetRGBA(x, y, color.RGBA{R: uint8(utils.ClampF64(fPixR/sum+0.5, 0, 255)),
				G: uint8(utils.ClampF64(fPixG/sum+0.5, 0, 255)),
				B: uint8(utils.ClampF64(fPixB/sum+0.5, 0, 255)),
				A: uint8(utils.ClampF64(fPixA/sum+0.5, 0, 255))})
		}
	}
	return res, nil
}

// ResizeGray resizes an grayscale (Gray) image.
// Input parameters: rbga imaga which will be resized; fx, fy scaling factors, their value has to be a positive float,
// the new size of the image will be computed as originalWidth * fx and originalHeight * fy; interpolation method,
// currently the following methods are supported: InterNearest, InterLinear, InterCatmullRom, InterLanczos.
// Example of usage:
//
//		res, err := resize.ResizeGray(img, 2.5, 3.5, resize.InterLinear)
//
func ResizeGray(img *image.Gray, fx float64, fy float64, interpolation Interpolation) (*image.Gray, error) {
	if fx < 0 || fy < 0 {
		return nil, errors.New("scale value should be greater then 0")
	}
	switch interpolation {
	case InterNearest:
		return resizeNearestGray(img, fx, fy)
	case InterLinear:
		return resizeLinearGray(img, fx, fy)
	case InterCatmullRom:
		return resizeCatmullRomGray(img, fx, fy)
	case InterLanczos:
		return resizeLanczosGray(img, fx, fy)
	}
	return nil, errors.New("invalid interpolation method")
}

// ResizeRGBA resizes an RGBA image.
// Input parameters: rbga imaga which will be resized; fx, fy scaling factors, their value has to be a positive float,
// the new size of the image will be computed as originalWidth * fx and originalHeight * fy; interpolation method,
// currently the following methods are supported: InterNearest, InterLinear, InterCatmullRom, InterLanczos.
// Example of usage:
//
//		res, err := resize.ResizeRGBA(img, 2.5, 3.5, resize.InterLinear)
//
func ResizeRGBA(img *image.RGBA, fx float64, fy float64, interpolation Interpolation) (*image.RGBA, error) {
	if fx < 0 || fy < 0 {
		return nil, errors.New("scale value should be greater then 0")
	}
	switch interpolation {
	case InterNearest:
		return resizeNearestRGBA(img, fx, fy)
	case InterLinear:
		return resizeLinearRGBA(img, fx, fy)
	case InterCatmullRom:
		return resizeCatmullRomRGBA(img, fx, fy)
	case InterLanczos:
		return resizeLanczosRGBA(img, fx, fy)
	}
	return nil, errors.New("invalid interpolation method")
}
