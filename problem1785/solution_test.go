// https://leetcode.com/problems/minimum-elements-to-add-to-form-a-given-sum/

package problem1785

import "testing"

func Test_minElements(t *testing.T) {
	type args struct {
		nums  []int
		limit int
		goal  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Test 1", args{[]int{1, -1, 1}, 3, -4}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minElements(tt.args.nums, tt.args.limit, tt.args.goal); got != tt.want {
				t.Errorf("minElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
