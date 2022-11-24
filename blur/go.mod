module github.com/ernyoke/imger/blur

go 1.18

replace github.com/ernyoke/imger/convolution => ../convolution

replace github.com/ernyoke/imger/padding => ../padding

replace github.com/ernyoke/imger/utils => ../utils

replace github.com/ernyoke/imger/imgio => ../imgio

require (
	github.com/ernyoke/imger/convolution v0.1.0
	github.com/ernyoke/imger/imgio v0.1.0
	github.com/ernyoke/imger/padding v0.1.0
	github.com/ernyoke/imger/utils v0.1.0
)
