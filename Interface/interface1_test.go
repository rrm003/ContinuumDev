package main

import (
	"testing"
)

func Testarea(t *testing.T) {
	
}

func Test_rectangle_area(t *testing.T) {
	type fields struct {
		length  float64
		breadth float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := &rectangle{
				length:  tt.fields.length,
				breadth: tt.fields.breadth,
			}
			if got := rec.area(); got != tt.want {
				t.Errorf("rectangle.area() = %v, want %v", got, tt.want)
			}
		})
	}
}
