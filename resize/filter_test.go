package resize

import (
	"testing"
	"github.com/ernyoke/imger/utils"
)

// --------------------------------Unit tests----------------------------------------
func Test_Linear_Positive_Valid(t *testing.T) {
	linear := NewLinear()
	expected := 0.5
	actual := linear.interpolate(0.5)
	if !utils.IsEqualFloat64(expected, actual) {
		t.Errorf("Expected %f is not equal to actual: %f\n", expected, actual)
	}
}

func Test_Linear_Negative_Valid(t *testing.T) {
	linear := NewLinear()
	expected := 0.5
	actual := linear.interpolate(-0.5)
	if !utils.IsEqualFloat64(expected, actual) {
		t.Errorf("Expected %f is not equal to actual: %f\n", expected, actual)
	}
}

func Test_Linear_Positive_Invalid(t *testing.T) {
	linear := NewLinear()
	expected := 0.0
	actual := linear.interpolate(1.5)
	if !utils.IsEqualFloat64(expected, actual) {
		t.Errorf("Expected %f is not equal to actual: %f\n", expected, actual)
	}
}

func Test_Linear_Negative_Invalid(t *testing.T) {
	linear := NewLinear()
	expected := 0.0
	actual := linear.interpolate(-1.5)
	if !utils.IsEqualFloat64(expected, actual) {
		t.Errorf("Expected %f is not equal to actual: %f\n", expected, actual)
	}
}

func Test_CatmullRom_Positive_Betwen0and1(t *testing.T) {
	catmullRom := NewCatmullRom()
	expected := 0.5625
	actual := catmullRom.interpolate(0.5)
	if !utils.IsEqualFloat64(expected, actual) {
		t.Errorf("Expected %f is not equal to actual: %f\n", expected, actual)
	}
}

func Test_CatmullRom_Negative_Betwen0and1(t *testing.T) {
	catmullRom := NewCatmullRom()
	expected := 0.5625
	actual := catmullRom.interpolate(-0.5)
	if !utils.IsEqualFloat64(expected, actual) {
		t.Errorf("Expected %f is not equal to actual: %f\n", expected, actual)
	}
}

func Test_CatmullRom_Positive_Betwen1and2(t *testing.T) {
	catmullRom := NewCatmullRom()
	expected := -0.0625
	actual := catmullRom.interpolate(1.5)
	if !utils.IsEqualFloat64(expected, actual) {
		t.Errorf("Expected %f is not equal to actual: %f\n", expected, actual)
	}
}

func Test_CatmullRom_Negative_Betwen1and2(t *testing.T) {
	catmullRom := NewCatmullRom()
	expected := -0.0625
	actual := catmullRom.interpolate(-1.5)
	if !utils.IsEqualFloat64(expected, actual) {
		t.Errorf("Expected %f is not equal to actual: %f\n", expected, actual)
	}
}

func Test_CatmullRom_Positive_Invalid(t *testing.T) {
	linear := NewLinear()
	expected := 0.0
	actual := linear.interpolate(2.6)
	if !utils.IsEqualFloat64(expected, actual) {
		t.Errorf("Expected %f is not equal to actual: %f\n", expected, actual)
	}
}

func Test_CatmullRom_Negative_Invalid(t *testing.T) {
	linear := NewLinear()
	expected := 0.0
	actual := linear.interpolate(-2.6)
	if !utils.IsEqualFloat64(expected, actual) {
		t.Errorf("Expected %f is not equal to actual: %f\n", expected, actual)
	}
}

func Test_Lanczos_Positive_Valid(t *testing.T) {
	lanczos := NewLanczos()
	expected := 0.60792710185
	actual := lanczos.interpolate(0.5)
	if !utils.IsEqualFloat64(expected, actual) {
		t.Errorf("Expected %f is not equal to actual: %f\n", expected, actual)
	}
}

func Test_Lanczos_Negative_Valid(t *testing.T) {
	lanczos := NewLanczos()
	expected := 0.60792710185
	actual := lanczos.interpolate(-0.5)
	if !utils.IsEqualFloat64(expected, actual) {
		t.Errorf("Expected %f is not equal to actual: %f\n", expected, actual)
	}
}

func Test_Lanczos_Positive_Invalid(t *testing.T) {
	lanczos := NewLanczos()
	expected := 0.0
	actual := lanczos.interpolate(3.5)
	if !utils.IsEqualFloat64(expected, actual) {
		t.Errorf("Expected %f is not equal to actual: %f\n", expected, actual)
	}
}

func Test_Lanczos_Negative_Invalid(t *testing.T) {
	lanczos := NewLanczos()
	expected := 0.0
	actual := lanczos.interpolate(-3.5)
	if !utils.IsEqualFloat64(expected, actual) {
		t.Errorf("Expected %f is not equal to actual: %f\n", expected, actual)
	}
}

func Test_Lanczos_0(t *testing.T) {
	lanczos := NewLanczos()
	expected := 0.0
	actual := lanczos.interpolate(0.0)
	if !utils.IsEqualFloat64(expected, actual) {
		t.Errorf("Expected %f is not equal to actual: %f\n", expected, actual)
	}
}
// ----------------------------------------------------------------------------------
