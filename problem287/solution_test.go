package problem287

import "testing"

func Test_findDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"T1", args{[]int{1, 3, 4, 2, 2}}, 2},
		{"T2", args{[]int{3, 1, 3, 4, 2}}, 3},
		{"T3", args{[]int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1}}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("findDuplicate() = %v, want %v", got, tt.want)
			}
			if got := findDuplicate2(tt.args.nums); got != tt.want {
				t.Errorf("findDuplicate2() = %v, want %v", got, tt.want)
			}
		})
	}
}
