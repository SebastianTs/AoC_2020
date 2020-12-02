package main

import (
	"testing"
)

func Test_isValidPart1(t *testing.T) {
	type args struct {
		e entry
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"example1", args{entry{1, 3, 'a', "abcde"}}, true},
		{"example2", args{entry{1, 3, 'b', "cdefg"}}, false},
		{"example3", args{entry{2, 9, 'c', "cccccccc"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidPart1(tt.args.e); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidPart2(t *testing.T) {
	type args struct {
		e entry
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"example1", args{entry{1, 3, 'a', "abcde"}}, true},
		{"example2", args{entry{1, 3, 'b', "cdefg"}}, false},
		{"example3", args{entry{2, 9, 'c', "ccccccccc"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidPart2(tt.args.e); got != tt.want {
				t.Errorf("isValidPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
