package main

import "testing"

type args struct {
	items  []Item
	weight uint64
}

type test struct {
	name     string
	args     args
	wantW    uint64
	wantV    uint64
	wantBits uint64
}

var tests []test

func TestMain(m *testing.M) {
	tests = []test{
		{
			name: "n=10",
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
			name: "n=20",
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
			name: "n=25",
			args: args{
				items: []Item{
					{Value: 3, Weight: 28},
					{Value: 10, Weight: 15},
					{Value: 9, Weight: 27},
					{Value: 5, Weight: 14},
					{Value: 7, Weight: 27},
					{Value: 8, Weight: 18},
					{Value: 8, Weight: 6},
					{Value: 5, Weight: 13},
					{Value: 9, Weight: 17},
					{Value: 4, Weight: 21},
					{Value: 7, Weight: 18},
					{Value: 8, Weight: 6},
					{Value: 8, Weight: 26},
					{Value: 3, Weight: 8},
					{Value: 9, Weight: 6},
					{Value: 10, Weight: 28},
					{Value: 9, Weight: 12},
					{Value: 9, Weight: 16},
					{Value: 4, Weight: 29},
					{Value: 7, Weight: 19},
					{Value: 6, Weight: 5},
					{Value: 3, Weight: 22},
					{Value: 3, Weight: 26},
					{Value: 4, Weight: 28},
					{Value: 10, Weight: 15},
				},
				weight: 400,
			},
			wantW:    0,
			wantV:    0,
			wantBits: 0,
		},
	}

	m.Run()
}

func Test_solve(t *testing.T) {
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
