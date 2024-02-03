package threshold

import (
	"fmt"
	"image"
	"testing"

	"github.com/ernyoke/imger/imgio"
)

// -----------------------------Acceptance tests------------------------------------
func setupTestCaseGray(t *testing.T) *image.Gray {
	path := "../res/girl.jpg"
	img, err := imgio.ImreadGray(path)
	if err != nil {
		t.Errorf("Could not read image from path: %s", path)
	}
	return img
}

func setupTestCaseOtsu(t *testing.T) *image.Gray {
	path := "../res/building.jpg"
	img, err := imgio.ImreadGray(path)
	if err != nil {
		t.Errorf("Could not read image from path: %s", path)
	}
	return img
}

func setupTestCaseGray16(t *testing.T) *image.Gray16 {
	path := "../res/girl.jpg"
	img, err := imgio.ImreadGray16(path)
	if err != nil {
		t.Errorf("Could not read image from path: %s", path)
	}
	return img
}

func tearDownTestCase(t *testing.T, img image.Image, path string) {
	err := imgio.Imwrite(img, path)
	if err != nil {
		t.Errorf("Could not write image to path: %s", path)
	}
}

func Test_Acceptance_ThresholdBinray(t *testing.T) {
	gray := setupTestCaseGray(t)
	thresh, _ := Threshold(gray, 100, ThreshBinary)
	tearDownTestCase(t, thresh, "../res/threshold/threshBin.jpg")
}

func Test_Acceptance_ThresholdBinrayInv(t *testing.T) {
	gray := setupTestCaseGray(t)
	thresh, _ := Threshold(gray, 100, ThreshBinaryInv)
	tearDownTestCase(t, thresh, "../res/threshold/threshBinInv.jpg")
}

func Test_Acceptance_ThresholdTrunc(t *testing.T) {
	gray := setupTestCaseGray(t)
	thresh, _ := Threshold(gray, 100, ThreshTrunc)
	tearDownTestCase(t, thresh, "../res/threshold/threshTrunc.jpg")
}

func Test_Acceptance_ThresholdToZero(t *testing.T) {
	gray := setupTestCaseGray(t)
	thresh, _ := Threshold(gray, 100, ThreshToZero)
	tearDownTestCase(t, thresh, "../res/threshold/threshToZero.jpg")
}

func Test_Acceptance_ThresholdToZeroInv(t *testing.T) {
	gray := setupTestCaseGray(t)
	thresh, _ := Threshold(gray, 100, ThreshToZeroInv)
	tearDownTestCase(t, thresh, "../res/threshold/threshBin.jpg")
}

func Test_Acceptance_Threshold16Bin(t *testing.T) {
	gray := setupTestCaseGray16(t)
	thresh, _ := Threshold16(gray, 32000, ThreshBinary)
	tearDownTestCase(t, thresh, "../res/threshold/thresh16Bin.jpg")
}

func Test_Acceptance_Threshold16BinInv(t *testing.T) {
	gray := setupTestCaseGray16(t)
	thresh, _ := Threshold16(gray, 32000, ThreshBinaryInv)
	tearDownTestCase(t, thresh, "../res/threshold/thresh16BinInv.jpg")
}

func Test_Acceptance_Threshold16Trunc(t *testing.T) {
	gray := setupTestCaseGray16(t)
	thresh, _ := Threshold16(gray, 32000, ThreshTrunc)
	tearDownTestCase(t, thresh, "../res/threshold/thresh16Trunc.jpg")
}

func Test_Acceptance_Threshold16ToZero(t *testing.T) {
	gray := setupTestCaseGray16(t)
	thresh, _ := Threshold16(gray, 32000, ThreshToZero)
	tearDownTestCase(t, thresh, "../res/threshold/thresh16ToZero.jpg")
}

func Test_Acceptance_Threshold16ToZeroInv(t *testing.T) {
	gray := setupTestCaseGray16(t)
	thresh, _ := Threshold16(gray, 32000, ThreshToZeroInv)
	tearDownTestCase(t, thresh, "../res/threshold/thresh16ToZeroInv.jpg")
}

func Test_Acceptance_OtsuThreshold(t *testing.T) {
	gray := setupTestCaseOtsu(t)
	thresh, _ := OtsuThreshold(gray, ThreshBinary)
	tearDownTestCase(t, thresh, "../res/threshold/otsuThreshBin.jpg")
}

func Test_Acceptance_OtsuThreshold_Cropped(t *testing.T) {
	thresholdMethods := map[string]Method{
		"Bin":       ThreshBinary,
		"BinInv":    ThreshBinaryInv,
		"Trunc":     ThreshTrunc,
		"ToZero":    ThreshToZero,
		"ToZeroInv": ThreshToZeroInv,
	}

	for name, meth := range thresholdMethods {
		t.Run(name, func(t *testing.T) {
			gray := setupTestCaseOtsu(t)
			cropped := gray.SubImage(image.Rect(100, 100, gray.Bounds().Max.X-100, gray.Bounds().Max.Y-100))
			thresh, err := OtsuThreshold(cropped.(*image.Gray), meth)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			tearDownTestCase(t, thresh, fmt.Sprintf("../res/threshold/otsuThresh%sCropped.jpg", name))
		})
	}
}

//---------------------------------------------------------------------------------
