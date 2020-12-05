package main

import "testing"

func Test_calculateID(t *testing.T) {
	type args struct {
		pass string
	}
	tests := []struct {
		name   string
		args   args
		wantID int
	}{
		{"example 1", args{"FBFBBFFRLR"}, 357},
		{"example 2", args{"BFFFBBFRRR"}, 567},
		{"example 3", args{"FFFBBBFRRR"}, 119},
		{"example 4", args{"BBFFBBFRLL"}, 820},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotID := calculateID(tt.args.pass); gotID != tt.wantID {
				t.Errorf("calculateID() = %v, want %v", gotID, tt.wantID)
			}
		})
	}
}
