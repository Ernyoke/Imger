package utils

import (
	"image"
	"testing"
)

func Test_ParallelForEachPixel(t *testing.T) {
	const (
		N = 777
		M = 999
	)
	expected := [N][M]int{}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			expected[i][j] = i + j
		}
	}
	actual := [N][M]int{}
	ParallelForEachPixel(image.Point{X: N, Y: M}, func(x int, y int) {
		actual[x][y] = x + y
	})
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if expected[i][j] != actual[i][j] {
				t.Errorf("Expected gray: %d - actual gray: %d at: %d %d", expected[i][j], actual[i][j], i, j)
			}
		}
	}
}
