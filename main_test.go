package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		items  []Item
		weight uint64
	}
	tests := []struct {
		name     string
		args     args
		wantW    uint64
		wantV    uint64
		wantBits uint64
	}{
		// TODO: Add test cases.
		{
			name: "n=10 & solve",
			args: args{
				items: []Item{
					{Value: 3, Weight: 15},
					{Value: 7, Weight: 3},
					{Value: 3, Weight: 12},
					{Value: 10, Weight: 26},
					{Value: 6, Weight: 14},
					{Value: 3, Weight: 21},
					{Value: 4, Weight: 6},
					{Value: 4, Weight: 18},
					{Value: 6, Weight: 13},
					{Value: 9, Weight: 3},
				},
				weight: 50,
			},
			wantW:    50,
			wantV:    20,
			wantBits: 42,
		},
		{
			name: "n=20 & solve",
			args: args{
				items: []Item{
					{Value: 6, Weight: 11},
					{Value: 3, Weight: 28},
					{Value: 7, Weight: 6},
					{Value: 4, Weight: 26},
					{Value: 9, Weight: 3},
					{Value: 6, Weight: 11},
					{Value: 5, Weight: 10},
					{Value: 5, Weight: 30},
					{Value: 3, Weight: 25},
					{Value: 5, Weight: 29},
					{Value: 4, Weight: 16},
					{Value: 10, Weight: 27},
					{Value: 4, Weight: 25},
					{Value: 7, Weight: 15},
					{Value: 3, Weight: 24},
					{Value: 9, Weight: 30},
					{Value: 4, Weight: 13},
					{Value: 5, Weight: 7},
					{Value: 8, Weight: 15},
					{Value: 3, Weight: 30},
				},
				weight: 60,
			},
			wantW:    60,
			wantV:    14,
			wantBits: 14,
		},
		{
			name: "n=40 & solve",
			args: args{
				items: []Item{
					{Value: 3, Weight: 10},
					{Value: 8, Weight: 17},
					{Value: 5, Weight: 26},
					{Value: 6, Weight: 10},
					{Value: 10, Weight: 13},
					{Value: 7, Weight: 3},
					{Value: 8, Weight: 29},
					{Value: 9, Weight: 18},
					{Value: 5, Weight: 18},
					{Value: 7, Weight: 14},
					{Value: 3, Weight: 14},
					{Value: 8, Weight: 29},
					{Value: 7, Weight: 25},
					{Value: 6, Weight: 23},
					{Value: 4, Weight: 19},
					{Value: 4, Weight: 14},
					{Value: 10, Weight: 4},
					{Value: 7, Weight: 16},
					{Value: 9, Weight: 22},
					{Value: 10, Weight: 23},
					{Value: 3, Weight: 18},
					{Value: 3, Weight: 17},
					{Value: 4, Weight: 5},
					{Value: 5, Weight: 5},
					{Value: 9, Weight: 9},
					{Value: 3, Weight: 16},
					{Value: 7, Weight: 8},
					{Value: 6, Weight: 17},
					{Value: 10, Weight: 16},
					{Value: 9, Weight: 21},
					{Value: 10, Weight: 20},
					{Value: 10, Weight: 3},
					{Value: 3, Weight: 26},
					{Value: 9, Weight: 8},
					{Value: 4, Weight: 5},
					{Value: 4, Weight: 30},
					{Value: 8, Weight: 22},
					{Value: 6, Weight: 3},
					{Value: 8, Weight: 5},
					{Value: 9, Weight: 18},
				},
				weight: 500,
			},
			wantW:    0,
			wantV:    0,
			wantBits: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotW, gotV, gotBits := solve(tt.args.items, tt.args.weight)

			if gotW != tt.wantW {
				t.Errorf("solve() gotW = %v, want %v", gotW, tt.wantW)
			}
			if gotV != tt.wantV {
				t.Errorf("solve() gotV = %v, want %v", gotV, tt.wantV)
			}
			if gotBits != tt.wantBits {
				t.Errorf("solve() gotBits = %v, want %v", gotBits, tt.wantBits)
			}
		})
	}
}

func Test_parallelSolve(t *testing.T) {
	type args struct {
		items  []Item
		weight uint64
	}
	tests := []struct {
		name     string
		args     args
		wantW    uint64
		wantV    uint64
		wantBits uint64
	}{
		// TODO: Add test cases.
		{
			name: "n=10 & parallelSolve",
			args: args{
				items: []Item{
					{Value: 3, Weight: 15},
					{Value: 7, Weight: 3},
					{Value: 3, Weight: 12},
					{Value: 10, Weight: 26},
					{Value: 6, Weight: 14},
					{Value: 3, Weight: 21},
					{Value: 4, Weight: 6},
					{Value: 4, Weight: 18},
					{Value: 6, Weight: 13},
					{Value: 9, Weight: 3},
				},
				weight: 50,
			},
			wantW:    50,
			wantV:    20,
			wantBits: 42,
		},
		{
			name: "n=20 & parallelSolve",
			args: args{
				items: []Item{
					{Value: 6, Weight: 11},
					{Value: 3, Weight: 28},
					{Value: 7, Weight: 6},
					{Value: 4, Weight: 26},
					{Value: 9, Weight: 3},
					{Value: 6, Weight: 11},
					{Value: 5, Weight: 10},
					{Value: 5, Weight: 30},
					{Value: 3, Weight: 25},
					{Value: 5, Weight: 29},
					{Value: 4, Weight: 16},
					{Value: 10, Weight: 27},
					{Value: 4, Weight: 25},
					{Value: 7, Weight: 15},
					{Value: 3, Weight: 24},
					{Value: 9, Weight: 30},
					{Value: 4, Weight: 13},
					{Value: 5, Weight: 7},
					{Value: 8, Weight: 15},
					{Value: 3, Weight: 30},
				},
				weight: 60,
			},
			wantW:    60,
			wantV:    14,
			wantBits: 14,
		},
		{
			name: "n=40 & parallelSolve",
			args: args{
				items: []Item{
					{Value: 3, Weight: 10},
					{Value: 8, Weight: 17},
					{Value: 5, Weight: 26},
					{Value: 6, Weight: 10},
					{Value: 10, Weight: 13},
					{Value: 7, Weight: 3},
					{Value: 8, Weight: 29},
					{Value: 9, Weight: 18},
					{Value: 5, Weight: 18},
					{Value: 7, Weight: 14},
					{Value: 3, Weight: 14},
					{Value: 8, Weight: 29},
					{Value: 7, Weight: 25},
					{Value: 6, Weight: 23},
					{Value: 4, Weight: 19},
					{Value: 4, Weight: 14},
					{Value: 10, Weight: 4},
					{Value: 7, Weight: 16},
					{Value: 9, Weight: 22},
					{Value: 10, Weight: 23},
					{Value: 3, Weight: 18},
					{Value: 3, Weight: 17},
					{Value: 4, Weight: 5},
					{Value: 5, Weight: 5},
					{Value: 9, Weight: 9},
					{Value: 3, Weight: 16},
					{Value: 7, Weight: 8},
					{Value: 6, Weight: 17},
					{Value: 10, Weight: 16},
					{Value: 9, Weight: 21},
					{Value: 10, Weight: 20},
					{Value: 10, Weight: 3},
					{Value: 3, Weight: 26},
					{Value: 9, Weight: 8},
					{Value: 4, Weight: 5},
					{Value: 4, Weight: 30},
					{Value: 8, Weight: 22},
					{Value: 6, Weight: 3},
					{Value: 8, Weight: 5},
					{Value: 9, Weight: 18},
				},
				weight: 500,
			},
			wantW:    0,
			wantV:    0,
			wantBits: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotW, gotV, gotBits := parallelSolve(tt.args.items, tt.args.weight)

			if gotW != tt.wantW {
				t.Errorf("parallelSolve() gotW = %v, want %v", gotW, tt.wantW)
			}
			if gotV != tt.wantV {
				t.Errorf("parallelSolve() gotV = %v, want %v", gotV, tt.wantV)
			}
			if gotBits != tt.wantBits {
				t.Errorf("parallelSolve() gotBits = %v, want %v", gotBits, tt.wantBits)
			}
		})
	}
}
