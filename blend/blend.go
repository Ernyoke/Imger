package blend

import (
	"errors"
	"github.com/Ernyoke/Imger/utils"
	"image"
	"image/color"
)

func AddGray(img1 *image.Gray, img2 *image.Gray) (*image.Gray, error) {
	size1 := img1.Bounds().Size()
	size2 := img2.Bounds().Size()
	if size1.X != size2.X || size1.Y != size2.Y {
		return nil, errors.New("the size of the two image does not match")
	}
	res := image.NewGray(img1.Bounds())
	utils.ForEachPixel(size1, func(x int, y int) {
		p1 := img1.GrayAt(x, y)
		p2 := img2.GrayAt(x, y)
		sum := uint16(p1.Y) + uint16(p2.Y)
		if sum > 255 {
			sum = 255
		}
		res.SetGray(x, y, color.Gray{uint8(sum)})
	})
	return res, nil
}

func AddGrayWeighted(img1 *image.Gray, w1 float64, img2 *image.Gray, w2 float64) (*image.Gray, error) {
	size1 := img1.Bounds().Size()
	size2 := img2.Bounds().Size()
	if size1.X != size2.X || size1.Y != size2.Y {
		return nil, errors.New("the size of the two image does not match")
	}
	res := image.NewGray(img1.Bounds())
	utils.ForEachPixel(size1, func(x int, y int) {
		p1 := img1.GrayAt(x, y)
		p2 := img2.GrayAt(x, y)
		sum := float64(p1.Y)*w1 + float64(p2.Y)*w2
		if sum > 255 {
			sum = 255
		}
		res.SetGray(x, y, color.Gray{uint8(sum)})
	})
	return res, nil
}
