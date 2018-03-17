package utils

import (
	"image"
	"math"
	"runtime"
	"sync"
)

// ParallelForEachPixel loops through the image and calls f functions for each [x, y] position.
// The image is divided into N * N blocks, where N is the number of available processor threads. For each block a
// parallel Goroutine is started.
func ParallelForEachPixel(size image.Point, f func(x int, y int)) {
	procs := runtime.GOMAXPROCS(0)
	var waitGroup sync.WaitGroup
	for i := 0; i < procs; i++ {
		startX := i * int(math.Floor(float64(size.X)/float64(procs)))
		var endX int
		if i < procs-1 {
			endX = (i + 1) * int(math.Floor(float64(size.X)/float64(procs)))
		} else {
			endX = size.X
		}
		for j := 0; j < procs; j++ {
			startY := j * int(math.Floor(float64(size.Y)/float64(procs)))
			var endY int
			if j < procs-1 {
				endY = (j + 1) * int(math.Floor(float64(size.Y)/float64(procs)))
			} else {
				endY = size.Y
			}
			waitGroup.Add(1)
			go func(sX int, eX int, sY int, eY int) {
				defer waitGroup.Done()
				for x := sX; x < eX; x++ {
					for y := sY; y < eY; y++ {
						f(x, y)
					}
				}
			}(startX, endX, startY, endY)
			waitGroup.Wait()
		}
	}
}
