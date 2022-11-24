package resize

import (
	"github.com/ernyoke/imger/utils"
	"testing"
)

// --------------------------------Unit tests----------------------------------------
func Test_Linear_Positive_Valid(t *testing.T) {
	linear := NewLinear()
	expected := 0.5
	actual := linear.Interpolate(0.5)
	if !utils.IsEqualFloat64(expected, actual) {
		t.Errorf("Expected %f is not equal to actual: %f\n", expected, actual)
	}
}

func Test_Linear_Negative_Valid(t *testing.T) {
	linear := NewLinear()
	expected := 0.5
	actual := linear.Interpolate(-0.5)
	if !utils.IsEqualFloat64(expected, actual) {
		t.Errorf("Expected %f is not equal to actual: %f\n", expected, actual)
	}
}

func Test_Linear_Positive_Invalid(t *testing.T) {
	linear := NewLinear()
	expected := 0.0
	actual := linear.Interpolate(1.5)
	if !utils.IsEqualFloat64(expected, actual) {
		t.Errorf("Expected %f is not equal to actual: %f\n", expected, actual)
	}
}

func Test_Linear_Negative_Invalid(t *testing.T) {
	linear := NewLinear()
	expected := 0.0
	actual := linear.Interpolate(-1.5)
	if !utils.IsEqualFloat64(expected, actual) {
		t.Errorf("Expected %f is not equal to actual: %f\n", expected, actual)
	}
}

func Test_CatmullRom_Positive_Betwen0and1(t *testing.T) {
	catmullRom := NewCatmullRom()
	expected := 0.5625
	actual := catmullRom.Interpolate(0.5)
	if !utils.IsEqualFloat64(expected, actual) {
		t.Errorf("Expected %f is not equal to actual: %f\n", expected, actual)
	}
}

func Test_CatmullRom_Negative_Betwen0and1(t *testing.T) {
	catmullRom := NewCatmullRom()
	expected := 0.5625
	actual := catmullRom.Interpolate(-0.5)
	if !utils.IsEqualFloat64(expected, actual) {
		t.Errorf("Expected %f is not equal to actual: %f\n", expected, actual)
	}
}

func Test_CatmullRom_Positive_Betwen1and2(t *testing.T) {
	catmullRom := NewCatmullRom()
	expected := -0.0625
	actual := catmullRom.Interpolate(1.5)
	if !utils.IsEqualFloat64(expected, actual) {
		t.Errorf("Expected %f is not equal to actual: %f\n", expected, actual)
	}
}

func Test_CatmullRom_Negative_Betwen1and2(t *testing.T) {
	catmullRom := NewCatmullRom()
	expected := -0.0625
	actual := catmullRom.Interpolate(-1.5)
	if !utils.IsEqualFloat64(expected, actual) {
		t.Errorf("Expected %f is not equal to actual: %f\n", expected, actual)
	}
}

func Test_CatmullRom_Positive_Invalid(t *testing.T) {
	linear := NewLinear()
	expected := 0.0
	actual := linear.Interpolate(2.6)
	if !utils.IsEqualFloat64(expected, actual) {
		t.Errorf("Expected %f is not equal to actual: %f\n", expected, actual)
	}
}

func Test_CatmullRom_Negative_Invalid(t *testing.T) {
	linear := NewLinear()
	expected := 0.0
	actual := linear.Interpolate(-2.6)
	if !utils.IsEqualFloat64(expected, actual) {
		t.Errorf("Expected %f is not equal to actual: %f\n", expected, actual)
	}
}

func Test_Lanczos_Positive_Valid(t *testing.T) {
	lanczos := NewLanczos()
	expected := 0.60792710185
	actual := lanczos.Interpolate(0.5)
	if !utils.IsEqualFloat64(expected, actual) {
		t.Errorf("Expected %f is not equal to actual: %f\n", expected, actual)
	}
}

func Test_Lanczos_Negative_Valid(t *testing.T) {
	lanczos := NewLanczos()
	expected := 0.60792710185
	actual := lanczos.Interpolate(-0.5)
	if !utils.IsEqualFloat64(expected, actual) {
		t.Errorf("Expected %f is not equal to actual: %f\n", expected, actual)
	}
}

func Test_Lanczos_Positive_Invalid(t *testing.T) {
	lanczos := NewLanczos()
	expected := 0.0
	actual := lanczos.Interpolate(3.5)
	if !utils.IsEqualFloat64(expected, actual) {
		t.Errorf("Expected %f is not equal to actual: %f\n", expected, actual)
	}
}

func Test_Lanczos_Negative_Invalid(t *testing.T) {
	lanczos := NewLanczos()
	expected := 0.0
	actual := lanczos.Interpolate(-3.5)
	if !utils.IsEqualFloat64(expected, actual) {
		t.Errorf("Expected %f is not equal to actual: %f\n", expected, actual)
	}
}

func Test_Lanczos_0(t *testing.T) {
	lanczos := NewLanczos()
	expected := 0.0
	actual := lanczos.Interpolate(0.0)
	if !utils.IsEqualFloat64(expected, actual) {
		t.Errorf("Expected %f is not equal to actual: %f\n", expected, actual)
	}
}

// ----------------------------------------------------------------------------------
