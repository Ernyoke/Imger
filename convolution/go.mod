module github.com/ernyoke/imger/convolution

go 1.18

replace github.com/ernyoke/imger/utils => ../utils

replace github.com/ernyoke/imger/padding => ../padding

replace github.com/ernyoke/imger/imgio => ../imgio

require (
	github.com/ernyoke/imger/padding v0.0.0-20221124211744-c532613c62c5
	github.com/ernyoke/imger/utils v0.0.0-20221124211744-c532613c62c5
)
