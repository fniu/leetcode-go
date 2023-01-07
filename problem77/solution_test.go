package problem77

import (
	"reflect"
	"testing"
)

func Test_combinestack(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"T1", args{4, 2}, [][]int{{3, 4}, {2, 3}, {2, 4}, {1, 2}, {1, 3}, {1, 4}}},
		{"T2", args{3, 1}, [][]int{{1}, {2}, {3}}},
		{"T3", args{1, 1}, [][]int{{1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combine(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combine() = %v, want %v", got, tt.want)
			}
		})
	}
}
