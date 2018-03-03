package resize

import "math"

type Filter interface {
	interpolate(float64) float64
	getS() float64
}

type Linear struct {}

func NewLinear() *Linear {
	return &Linear{}
}

func (r *Linear) interpolate(x float64) float64 {
	x = math.Abs(x)
	if x < 1.0 {
		return 1.0 - x
	}
	return 0
}

func (r *Linear) getS() float64 {
	return 1.0
}

type CatmullRom struct {}

func NewCatmullRom() *CatmullRom {
	return &CatmullRom{}
}

func (r *CatmullRom) interpolate(x float64) float64 {
	b := 0.0
	c := 0.5
	x = math.Abs(x)

	if x < 1.0 {
		return (6 - 2*b + (-18+12*b+6*c)*math.Pow(x, 2) + (12-9*b-6*c)*math.Pow(x, 3)) / 6
	} else if x <= 2.0 {
		return (8*b + 24*c + (-12*b-48*c)*x + (6*b+30*c)*math.Pow(x, 2) + (-b-6*c)*math.Pow(x, 3)) / 6
	}
	return 0
}

func (r *CatmullRom) getS() float64 {
	return 2.0
}

type Lanczos struct {}

func NewLanczos() *Lanczos {
	return &Lanczos{}
}

func (r *Lanczos) interpolate(x float64) float64 {
	x = math.Abs(x)
	if x > 0.0 && x < 3.0 {
		return (3.0 * math.Sin(math.Pi*x) * math.Sin(math.Pi*(x/3.0))) / (math.Pi * math.Pi * x * x)
	}
	return 0.0
}

func (r *Lanczos) getS() float64 {
	return 3.0
}
