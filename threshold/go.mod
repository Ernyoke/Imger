module github.com/ernyoke/imger/threshold

go 1.18

replace github.com/ernyoke/imger/utils => ../utils

replace github.com/ernyoke/imger/histogram => ../histogram

replace github.com/ernyoke/imger/imgio => ../imgio

require (
	github.com/ernyoke/imger/histogram v0.1.0
	github.com/ernyoke/imger/imgio v0.1.0
	github.com/ernyoke/imger/utils v0.1.0
)
