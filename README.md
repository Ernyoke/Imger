# Imger

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
go get -u https://github.com/Ernyoke/Imger/...
```

## Test and examples
You can find tests and examples for every package on the `dev` branch.

## License
This project is under the MIT License. See the LICENSE file for the full license text.
