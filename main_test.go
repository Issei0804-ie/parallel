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
					{Value: 546, Weight: 2712},
					{Value: 966, Weight: 1384}, //
					{Value: 302, Weight: 1321},
					{Value: 948, Weight: 2890},
					{Value: 594, Weight: 2541},
					{Value: 677, Weight: 2411},
					{Value: 789, Weight: 2621}, //
					{Value: 666, Weight: 2997},
					{Value: 930, Weight: 438}, //
					{Value: 470, Weight: 1177},
				},
				weight: 4700,
			},
			wantW:    4443,
			wantV:    2685,
			wantBits: 322,
		},
		{
			name: "n=20",
			args: args{
				items: []Item{
					{Value: 668, Weight: 1888},
					{Value: 539, Weight: 1212},
					{Value: 363, Weight: 2458},
					{Value: 658, Weight: 838},
					{Value: 620, Weight: 2362},
					{Value: 906, Weight: 2133},
					{Value: 623, Weight: 2891},
					{Value: 472, Weight: 2296},
					{Value: 446, Weight: 2464},
					{Value: 674, Weight: 1094},
					{Value: 631, Weight: 1444},
					{Value: 726, Weight: 2708},
					{Value: 311, Weight: 2869},
					{Value: 771, Weight: 689},
					{Value: 548, Weight: 1464},
					{Value: 542, Weight: 668},
					{Value: 625, Weight: 2269},
					{Value: 534, Weight: 1680},
					{Value: 495, Weight: 2597},
					{Value: 957, Weight: 1535},
				},
				weight: 4700,
			},
			wantW:    4501,
			wantV:    3184,
			wantBits: 41482,
		},
		{
			name: "n=25",
			args: args{
				items: []Item{
					{Value: 577, Weight: 2139},
					{Value: 693, Weight: 1651},
					{Value: 715, Weight: 2620},
					{Value: 313, Weight: 2682},
					{Value: 732, Weight: 609},
					{Value: 320, Weight: 2115},
					{Value: 639, Weight: 1908},
					{Value: 889, Weight: 2802},
					{Value: 438, Weight: 2010},
					{Value: 491, Weight: 2081},
					{Value: 359, Weight: 1392},
					{Value: 963, Weight: 2728},
					{Value: 353, Weight: 2362},
					{Value: 930, Weight: 590},
					{Value: 820, Weight: 904},
					{Value: 927, Weight: 1048},
					{Value: 520, Weight: 1665},
					{Value: 960, Weight: 2951},
					{Value: 930, Weight: 2908},
					{Value: 763, Weight: 2124},
					{Value: 304, Weight: 2786},
					{Value: 326, Weight: 1821},
					{Value: 506, Weight: 1014},
					{Value: 812, Weight: 2836},
					{Value: 382, Weight: 2442},
				},
				weight: 4700,
			},
			wantW:    4165,
			wantV:    3915,
			wantBits: 4251664,
		},
		{
			name: "n=26",
			args: args{
				items: []Item{
					{Value: 532, Weight: 2287},
					{Value: 730, Weight: 1799},
					{Value: 386, Weight: 576},
					{Value: 928, Weight: 1269},
					{Value: 974, Weight: 2279},
					{Value: 750, Weight: 2609},
					{Value: 346, Weight: 1468},
					{Value: 643, Weight: 2350},
					{Value: 946, Weight: 538},
					{Value: 916, Weight: 2375},
					{Value: 622, Weight: 2774},
					{Value: 376, Weight: 2172},
					{Value: 533, Weight: 1481},
					{Value: 420, Weight: 2249},
					{Value: 825, Weight: 2513},
					{Value: 639, Weight: 2220},
					{Value: 785, Weight: 1507},
					{Value: 329, Weight: 779},
					{Value: 380, Weight: 313},
					{Value: 963, Weight: 2884},
					{Value: 910, Weight: 2200},
					{Value: 658, Weight: 577},
					{Value: 659, Weight: 2154},
					{Value: 476, Weight: 2207},
					{Value: 786, Weight: 1606},
					{Value: 375, Weight: 2093},
				},
				weight: 4700,
			},
			wantW:    4566,
			wantV:    3704,
			wantBits: 18874636,
		},
		{
			name: "n=27",
			args: args{
				items: []Item{
					{Value: 794, Weight: 691},
					{Value: 509, Weight: 586},
					{Value: 907, Weight: 1458},
					{Value: 899, Weight: 1546},
					{Value: 690, Weight: 1488},
					{Value: 793, Weight: 1406},
					{Value: 455, Weight: 1303},
					{Value: 912, Weight: 509},
					{Value: 666, Weight: 2906},
					{Value: 484, Weight: 1822},
					{Value: 787, Weight: 695},
					{Value: 514, Weight: 414},
					{Value: 756, Weight: 1869},
					{Value: 756, Weight: 739},
					{Value: 300, Weight: 2931},
					{Value: 921, Weight: 2023},
					{Value: 940, Weight: 2444},
					{Value: 982, Weight: 2568},
					{Value: 688, Weight: 2836},
					{Value: 428, Weight: 482},
					{Value: 568, Weight: 2811},
					{Value: 711, Weight: 2841},
					{Value: 475, Weight: 2743},
					{Value: 474, Weight: 1356},
					{Value: 854, Weight: 1286},
					{Value: 609, Weight: 1967},
					{Value: 898, Weight: 2025},
				},
				weight: 4700,
			},
			wantW:    4663,
			wantV:    4798,
			wantBits: 17304707,
		},
		{
			name: "n=28",
			args: args{
				items: []Item{
					{Value: 380, Weight: 2911},
					{Value: 599, Weight: 2403},
					{Value: 971, Weight: 1838},
					{Value: 837, Weight: 1697},
					{Value: 329, Weight: 1221},
					{Value: 351, Weight: 1030},
					{Value: 913, Weight: 1156},
					{Value: 604, Weight: 1362},
					{Value: 735, Weight: 2433},
					{Value: 441, Weight: 1120},
					{Value: 732, Weight: 2679},
					{Value: 889, Weight: 1588},
					{Value: 908, Weight: 1156},
					{Value: 791, Weight: 2252},
					{Value: 436, Weight: 412},
					{Value: 685, Weight: 439},
					{Value: 339, Weight: 2582},
					{Value: 990, Weight: 1082},
					{Value: 774, Weight: 2225},
					{Value: 328, Weight: 1573},
					{Value: 636, Weight: 1111},
					{Value: 900, Weight: 1522},
					{Value: 703, Weight: 1259},
					{Value: 818, Weight: 1580},
					{Value: 664, Weight: 2001},
					{Value: 417, Weight: 1287},
					{Value: 982, Weight: 1262},
					{Value: 857, Weight: 1879},
				},
				weight: 4700,
			},
			wantW:    4351,
			wantV:    4006,
			wantBits: 67289152,
		},
		{
			name: "n=29",
			args: args{
				items: []Item{
					{Value: 554, Weight: 1986},
					{Value: 321, Weight: 1492},
					{Value: 861, Weight: 1097},
					{Value: 395, Weight: 2355},
					{Value: 631, Weight: 1747},
					{Value: 365, Weight: 676},
					{Value: 773, Weight: 309},
					{Value: 599, Weight: 805},
					{Value: 887, Weight: 1941},
					{Value: 617, Weight: 2153},
					{Value: 434, Weight: 700},
					{Value: 590, Weight: 1358},
					{Value: 853, Weight: 2086},
					{Value: 760, Weight: 436},
					{Value: 645, Weight: 1000},
					{Value: 855, Weight: 1620},
					{Value: 572, Weight: 2315},
					{Value: 655, Weight: 872},
					{Value: 821, Weight: 2773},
					{Value: 862, Weight: 1805},
					{Value: 984, Weight: 1009},
					{Value: 629, Weight: 2140},
					{Value: 503, Weight: 1663},
					{Value: 403, Weight: 2550},
					{Value: 579, Weight: 2721},
					{Value: 422, Weight: 2115},
					{Value: 831, Weight: 2664},
					{Value: 981, Weight: 1077},
					{Value: 377, Weight: 1738},
				},
				weight: 4700,
			},
			wantW:    4628,
			wantV:    4793,
			wantBits: 135275588,
		},
		{
			name: "n=30",
			args: args{
				items: []Item{
					{Value: 788, Weight: 1971},
					{Value: 664, Weight: 1437},
					{Value: 861, Weight: 348},
					{Value: 304, Weight: 1895},
					{Value: 991, Weight: 1135},
					{Value: 738, Weight: 1458},
					{Value: 797, Weight: 1827},
					{Value: 567, Weight: 1537},
					{Value: 968, Weight: 2169},
					{Value: 550, Weight: 1926},
					{Value: 776, Weight: 1342},
					{Value: 432, Weight: 2405},
					{Value: 447, Weight: 2201},
					{Value: 1000, Weight: 1668},
					{Value: 441, Weight: 2359},
					{Value: 400, Weight: 1847},
					{Value: 477, Weight: 395},
					{Value: 754, Weight: 1301},
					{Value: 547, Weight: 861},
					{Value: 480, Weight: 1926},
					{Value: 954, Weight: 1479},
					{Value: 334, Weight: 1451},
					{Value: 495, Weight: 2375},
					{Value: 703, Weight: 1243},
					{Value: 891, Weight: 1388},
					{Value: 676, Weight: 1630},
					{Value: 537, Weight: 1900},
					{Value: 796, Weight: 1198},
					{Value: 593, Weight: 699},
					{Value: 363, Weight: 662},
				},
				weight: 4700,
			},
			wantW:    4636,
			wantV:    4265,
			wantBits: 402980884,
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
