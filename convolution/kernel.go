package convolution

import (
	"errors"
	"image"
	"math"
)

type Matrix interface {
	At(x, y int) float64
}

type Kernel struct {
	Content [][]float64
	Width   int
	Height  int
}

func NewKernel(width int, height int) (*Kernel, error) {
	if width < 0 || height < 0 {
		return nil, errors.New("negative kernel size")
	}
	m := make([][]float64, height)
	for i := range m {
		m[i] = make([]float64, width)
	}
	return &Kernel{Content: m, Width: width, Height: height}, nil
}

func (k *Kernel) At(x, y int) float64 {
	return k.Content[x][y]
}

func (k *Kernel) Set(x int, y int, value float64) {
	k.Content[x][y] = value
}

func (k *Kernel) Size() image.Point {
	return image.Point{X: k.Width, Y: k.Height}
}

func (k *Kernel) AbSum() float64 {
	var sum float64
	for x := 0; x < k.Height; x++ {
		for y := 0; y < k.Width; y++ {
			sum += math.Abs(k.At(x, y))
		}
	}
	return sum
}

func (k *Kernel) Normalize() *Kernel {
	normalized, _ := NewKernel(k.Width, k.Height)
	sum := k.AbSum()
	if sum == 0 {
		sum = 1

	}
	for x := 0; x < k.Height; x++ {
		for y := 0; y < k.Width; y++ {
			normalized.Set(x, y, k.At(x, y)/sum)
		}
	}
	return normalized
}
