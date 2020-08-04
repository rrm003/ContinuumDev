package test

import (
	"testing"
)

type datadivide struct {
	x      int
	y      int
	output float64
	err    string
}

func TestDivide(t *testing.T) {
	input := []datadivide{{1, 2, float64(1 / 2), ""},
		{0, 0, float64(0), "Divide by zero"},
		{0, 10, 0, ""},
		{1, -1, -1, ""},
		{21, -1, -21, ""},
		{-20, -90, float64(20 / 90), ""}}

	for _, in := range input {
		if out, err := Divide(in.x, in.y); in.err != err && in.output != out {
			t.Fail()
		}
	}
}
