package main

import (
	"testing"
)

func Test_findInvalid(t *testing.T) {
	type args struct {
		ns       []int
		preamble int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{[]int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}, 5}, 127},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findInvalid(tt.args.ns, tt.args.preamble); got != tt.want {
				t.Errorf("findInvalid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findContinuesSet(t *testing.T) {
	type args struct {
		ns      []int
		invalid int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{[]int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}, 127}, 62},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findContinuesSet(tt.args.ns, tt.args.invalid); got != tt.want {
				t.Errorf("findContinuesSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
