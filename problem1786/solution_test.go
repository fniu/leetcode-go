package problem1786

import "testing"

func Test_countRestrictedPaths(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Test 1", args{5, [][]int{{1, 2, 3}, {1, 3, 3}, {2, 3, 1}, {1, 4, 2}, {5, 2, 2}, {3, 5, 1}, {5, 4, 10}}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countRestrictedPaths(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("countRestrictedPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
