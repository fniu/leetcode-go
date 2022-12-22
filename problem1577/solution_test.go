package problem1577

import "testing"

func Test_numTriplets(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Test 1", args{[]int{7, 4}, []int{5, 2, 8, 9}}, 1},
		{"Test 2", args{[]int{1, 1}, []int{1, 1, 1}}, 9},
		{"Test 3", args{[]int{1, 9}, []int{3}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTriplets(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("numTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
