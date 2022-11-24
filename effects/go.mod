module github.com/ernyoke/imger/effects

go 1.18

replace github.com/ernyoke/imger/utils => ../utils

replace github.com/ernyoke/imger/padding => ../padding

replace github.com/ernyoke/imger/grayscale => ../grayscale

replace github.com/ernyoke/imger/convolution => ../convolution

replace github.com/ernyoke/imger/blend => ../blend

replace github.com/ernyoke/imger/resize => ../resize

replace github.com/ernyoke/imger/imgio => ../imgio

require (
	github.com/ernyoke/imger/blend v0.0.0-00010101000000-000000000000
	github.com/ernyoke/imger/convolution v0.0.0-00010101000000-000000000000
	github.com/ernyoke/imger/grayscale v0.0.0-00010101000000-000000000000
	github.com/ernyoke/imger/imgio v0.0.0-00010101000000-000000000000
	github.com/ernyoke/imger/padding v0.0.0-00010101000000-000000000000
	github.com/ernyoke/imger/resize v0.0.0-00010101000000-000000000000
	github.com/ernyoke/imger/utils v0.0.0-00010101000000-000000000000
)
