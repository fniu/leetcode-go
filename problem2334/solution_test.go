package problem2334

import "testing"

func Test_validSubarraySize(t *testing.T) {
	type args struct {
		nums      []int
		threshold int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"example 0", args{[]int{3, 3, 3, 3, 3}, 4}, 2},
		{"example 1", args{[]int{1, 3, 4, 3, 1}, 6}, 3},
		{"example 2", args{[]int{6, 5, 6, 5, 8}, 7}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validSubarraySize(tt.args.nums, tt.args.threshold); got != tt.want {
				t.Errorf("validSubarraySize() = %v, want %v", got, tt.want)
			}
		})
	}
}
