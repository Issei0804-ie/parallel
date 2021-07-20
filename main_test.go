package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		i       interface{}
		weights int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"test", args{
			i: map[int]map[int]int{
				0: {1: 1},
				1: {1: 1},
			},

			weights: 0,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
