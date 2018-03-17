package resize

import "math"

// Filter - Interface for resampling filters
type Filter interface {
	Interpolate(float64) float64
	GetS() float64
}

// Linear - Struct for Linear filter
type Linear struct{}

// NewLinear creates a new Linear filter
func NewLinear() *Linear {
	return &Linear{}
}

// Interpolate returns the coefficient for x value using Linear interpolation
func (r *Linear) Interpolate(x float64) float64 {
	x = math.Abs(x)
	if x < 1.0 {
		return 1.0 - x
	}
	return 0
}

// GetS returns the support value for Linear filter
func (r *Linear) GetS() float64 {
	return 1.0
}

// CatmullRom - Struct for Catmull-Rom filter
type CatmullRom struct{}

// NewCatmullRom creates a new Catmull-Rom filter
func NewCatmullRom() *CatmullRom {
	return &CatmullRom{}
}

// Interpolate returns the coefficient for x value using Catmull-Rom interpolation
func (r *CatmullRom) Interpolate(x float64) float64 {
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

// GetS returns the support value for Catmull-Rom filter
func (r *CatmullRom) GetS() float64 {
	return 2.0
}

// Lanczos - struct for Lanczos filter
type Lanczos struct{}

// NewLanczos creates a new Lanczos filter
func NewLanczos() *Lanczos {
	return &Lanczos{}
}

// Interpolate returns the coefficient for x value using Lanczos interpolation
func (r *Lanczos) Interpolate(x float64) float64 {
	x = math.Abs(x)
	if x > 0.0 && x < 3.0 {
		return (3.0 * math.Sin(math.Pi*x) * math.Sin(math.Pi*(x/3.0))) / (math.Pi * math.Pi * x * x)
	}
	return 0.0
}

// GetS returns the support value for Lanczos filter
func (r *Lanczos) GetS() float64 {
	return 3.0
}
