# Imger
[![MIT License](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://github.com/anthonynsimon/bild/blob/master/LICENSE) 
[![Build Status](https://travis-ci.org/Ernyoke/Imger.svg?branch=dev)](https://travis-ci.org/Ernyoke/Imger)
[![Go Report Card](https://goreportcard.com/badge/github.com/Ernyoke/Imger)](https://goreportcard.com/report/github.com/Ernyoke/Imger)

This repository contains a collection of image processing algorithms written in pure Go. The packages are under development, their API may change over time.

## Currently supported
* IO (ImreadGray, ImreadGray16, ImreadRGBA, ImreadRGBA64, Imwrite). Supported extensions: jpg, jpeg, png
* Grayscale
* Threshold (Binary, BinaryInv, Trunc, ToZero, ToZeroInv)
* Image padding (BorderConstant, BorderReplicate, BorderReflect)
* Convolution
* Blur (Average - Box, Gaussian)
* Edge detection (Sobel, Laplacian)
* Resize (Nearest Neighbour, Linear, Catmull-Rom, Lanczos)

## Install
```
go get -u github.com/Ernyoke/Imger/...
```

## Merge into ```master```
```
git checkout master
git merge --no-commit --no-ff dev
```

## License
This project is under the MIT License. See the LICENSE file for the full license text.