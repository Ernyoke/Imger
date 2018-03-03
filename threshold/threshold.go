package threshold

import (
	"errors"
	"github.com/Ernyoke/Imger/utils"
	"image"
	"image/color"
)

type Method int

const (
	ThreshBinary Method = iota
	ThreshBinaryInv
	ThreshTrunc
	ThreshToZero
	ThreshToZeroInv
)

func Threshold(img *image.Gray, t uint8, method Method) (*image.Gray, error) {
	var setPixel func(*image.Gray, int, int)
	switch method {
	case ThreshBinary:
		setPixel = func(gray *image.Gray, x int, y int) {
			pixel := img.GrayAt(x, y).Y
			if pixel < t {
				gray.SetGray(x, y, color.Gray{Y: utils.MinUint8})
			} else {
				gray.SetGray(x, y, color.Gray{Y: utils.MaxUint8})
			}
		}
	case ThreshBinaryInv:
		setPixel = func(gray *image.Gray, x int, y int) {
			pixel := img.GrayAt(x, y).Y
			if pixel < t {
				gray.SetGray(x, y, color.Gray{Y: utils.MaxUint8})
			} else {
				gray.SetGray(x, y, color.Gray{Y: utils.MinUint8})
			}
		}
	case ThreshTrunc:
		{
			setPixel = func(gray *image.Gray, x int, y int) {
				pixel := img.GrayAt(x, y).Y
				if pixel < t {
					gray.SetGray(x, y, color.Gray{Y: pixel})
				} else {
					gray.SetGray(x, y, color.Gray{Y: t})
				}
			}
		}
	case ThreshToZero:
		setPixel = func(gray *image.Gray, x int, y int) {
			pixel := img.GrayAt(x, y).Y
			if pixel < t {
				gray.SetGray(x, y, color.Gray{Y: utils.MinUint8})
			} else {
				gray.SetGray(x, y, color.Gray{Y: pixel})
			}
		}
	case ThreshToZeroInv:
		setPixel = func(gray *image.Gray, x int, y int) {
			pixel := img.GrayAt(x, y).Y
			if pixel < t {
				gray.SetGray(x, y, color.Gray{Y: pixel})
			} else {
				gray.SetGray(x, y, color.Gray{Y: utils.MinUint8})
			}
		}
	default:
		return nil, errors.New("invalid threshold method")
	}
	return threshold(img, setPixel), nil
}

func Threshold16(img *image.Gray16, t uint16, method Method) (*image.Gray16, error) {
	var setPixel func(*image.Gray16, int, int)
	switch method {
	case ThreshBinary:
		setPixel = func(gray *image.Gray16, x int, y int) {
			pixel := img.Gray16At(x, y).Y
			if pixel < t {
				gray.SetGray16(x, y, color.Gray16{Y: utils.MinUint16})
			} else {
				gray.SetGray16(x, y, color.Gray16{Y: utils.MaxUint16})
			}
		}
	case ThreshBinaryInv:
		setPixel = func(gray *image.Gray16, x int, y int) {
			pixel := img.Gray16At(x, y).Y
			if pixel < t {
				gray.SetGray16(x, y, color.Gray16{Y: utils.MaxUint16})
			} else {
				gray.SetGray16(x, y, color.Gray16{Y: utils.MinUint16})
			}
		}
	case ThreshTrunc:
		{
			setPixel = func(gray *image.Gray16, x int, y int) {
				pixel := img.Gray16At(x, y).Y
				if pixel < t {
					gray.SetGray16(x, y, color.Gray16{Y: pixel})
				} else {
					gray.SetGray16(x, y, color.Gray16{Y: t})
				}
			}
		}
	case ThreshToZero:
		setPixel = func(gray *image.Gray16, x int, y int) {
			pixel := img.Gray16At(x, y).Y
			if pixel < t {
				gray.SetGray16(x, y, color.Gray16{Y: utils.MinUint16})
			} else {
				gray.SetGray16(x, y, color.Gray16{Y: pixel})
			}
		}
	case ThreshToZeroInv:
		setPixel = func(gray *image.Gray16, x int, y int) {
			pixel := img.Gray16At(x, y).Y
			if pixel < t {
				gray.SetGray16(x, y, color.Gray16{pixel})
			} else {
				gray.SetGray16(x, y, color.Gray16{Y: utils.MinUint16})
			}
		}
	default:
		return nil, errors.New("invalid threshold method")
	}
	return threshold16(img, setPixel), nil
}

// -------------------------------------------------------------------------------------------------------
func threshold(img *image.Gray, setPixel func(*image.Gray, int, int)) *image.Gray {
	size := img.Bounds().Size()
	gray := image.NewGray(img.Bounds())
	utils.ForEachPixel(size, func(x, y int) {
		setPixel(gray, x, y)
	})
	return gray
}

func threshold16(img *image.Gray16, setPixel16 func(*image.Gray16, int, int)) *image.Gray16 {
	size := img.Bounds().Size()
	gray := image.NewGray16(img.Bounds())
	utils.ForEachPixel(size, func(x, y int) {
		setPixel16(gray, x, y)
	})
	return gray
}
