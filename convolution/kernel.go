package convolution

import (
	"errors"
	"image"
)

type Matrix interface {
	At(x, y int) float64
}

type Kernel struct {
	content [][]float64
	Width int
	Height int
}

func NewKernel(width int, height int) (*Kernel, error){
	if width < 0 || height < 0 {
		return nil, errors.New("negative kernel size")
	}
	m := make([][]float64, height)
	for i := range m {
		m[i] = make([]float64, width)
	}
	return &Kernel{content: m, Width: width, Height: height}, nil
}

func (k *Kernel) At(x, y int) float64 {
	return k.content[x][y]
}

func (k *Kernel) Set(x int, y int, value float64) {
	k.content[x][y] = value
}

func (k *Kernel) Size() image.Point {
	return image.Point{X: k.Width, Y: k.Height}
}
