package test

import (
	"testing"
)

type dataadd struct {
	x      int
	y      int
	output int
}

func TestAdd(t *testing.T) {
	input := []dataadd{{1, 2, 3},
		{0, 0, 0},
		{0, 10, 10},
		{1, -1, 0},
		{21, -1, 20},
		{-20, -90, -110}}

	for _, in := range input {
		if out := Add(in.x, in.y); out != in.output {
			t.Fail()
		}
	}
}
