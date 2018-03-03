package resize

import "math"

type Filter interface {
	interpolate(float64) float64
	getS() float64
}

type Linear struct {
	S float64
}

func NewLinear() *Linear {
	return &Linear{ S: 1.0 }
}

func (r *Linear) interpolate(x float64) float64 {
	x = math.Abs(x)
	if x < 1.0 {
		return 1.0 - x
	}
	return 0
}

func (r* Linear) getS() float64 {
	return r.S
}

type CatmullRom struct {
	S float64
}

func NewCatmullRom() *CatmullRom {
	return &CatmullRom{ S: 2.0 }
}

func (r *CatmullRom) interpolate(x float64) float64 {
	b := 0.0
	c := 0.5
	var w [4]float64
	x = math.Abs(x)

	if x < 1.0 {
		w[0] = 0
		w[1] = 6 - 2*b
		w[2] = (-18 + 12*b + 6*c) * x * x
		w[3] = (12 - 9*b - 6*c) * x * x * x
	} else if x <= 2.0 {
		w[0] = 8*b + 24*c
		w[1] = (-12*b - 48*c) * x
		w[2] = (6*b + 30*c) * x * x
		w[3] = (-b - 6*c) * x * x * x
	} else {
		return 0
	}
	return (w[0] + w[1] + w[2] + w[3]) / 6
}

func (r* CatmullRom) getS() float64 {
	return r.S
}

type Lanczos struct {
	S float64
}

func NewLanczos() *Lanczos {
	return &Lanczos{ S: 3.0 }
}

func (r *Lanczos) interpolate(x float64) float64 {
	x = math.Abs(x)
	if x == 0 {
		return 1.0
	} else if x < 3.0 {
		return (3.0 * math.Sin(math.Pi*x) * math.Sin(math.Pi*(x/3.0))) / (math.Pi * math.Pi * x * x)
	}
	return 0.0
}

func (r* Lanczos) getS() float64 {
	return r.S
}

