package main

import (
	"errors"
	"testing"
)

func Test_findSum(t *testing.T) {
	type args struct {
		exps []int
		sum  int
	}
	type result struct {
		result int
		err    error
	}
	tests := []struct {
		name string
		args args
		want result
	}{
		{"example1", args{[]int{1721, 979, 366, 299, 675, 1456}, 2020}, result{514579, nil}},
		{"example2", args{[]int{}, 2020}, result{0, errors.New("No result found")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := result{}
			if got.result, got.err = findSum(tt.args.exps, tt.args.sum); got.result != tt.want.result {
				t.Errorf("findSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSumPart2(t *testing.T) {
	type args struct {
		exps []int
		sum  int
	}
	type result struct {
		result int
		err    error
	}
	tests := []struct {
		name string
		args args
		want result
	}{
		{"example1", args{[]int{1721, 979, 366, 299, 675, 1456}, 2020}, result{241861950, nil}},
		{"example2", args{[]int{}, 2020}, result{0, errors.New("No Result found")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := result{}
			if got.result, got.err = findSumPart2(tt.args.exps, tt.args.sum); got.result != tt.want.result {
				t.Errorf("findSumPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
