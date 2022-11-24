# Imger
[![MIT License](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://github.com/anthonynsimon/bild/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/Ernyoke/Imger)](https://goreportcard.com/report/github.com/Ernyoke/Imger)

This repository contains a collection of image processing algorithms written in pure Go.

## Currently supported
* IO (ImreadGray, ImreadGray16, ImreadRGBA, ImreadRGBA64, Imwrite). Supported extensions: jpg, jpeg, png
* Grayscale
* Blend (AddScalarToGray, AddGray, AddGrayWeighted)
* Threshold (Binary, BinaryInv, Trunc, ToZero, ToZeroInv, Otsu)
* Image padding (BorderConstant, BorderReplicate, BorderReflect)
* Convolution
* Blur (Average - Box, Gaussian)
* Edge detection (Sobel, Laplacian, Canny)
* Resize (Nearest Neighbour, Linear, Catmull-Rom, Lanczos)
* Effects (Pixelate, Sepia, Emboss, Sharpen, Invert)
* Transform (Rotate)

## Install
```bash
go get -u github.com/ernyoke/imger/...
```

## Running the Tests

```bash
go test ./...
```

## License
This project is under the MIT License. See the LICENSE file for the full license text.