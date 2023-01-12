package problem229

import (
	"reflect"
	"testing"
)

func Test_majorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"T0", args{[]int{3, 2, 3}}, []int{3}},
		{"T1", args{[]int{3, 2, 3, 5, 5}}, []int{3, 5}},
		{"T2", args{[]int{1}}, []int{1}},
		{"T3", args{[]int{1, 2}}, []int{1, 2}},
		{"T4", args{[]int{3, 2, 3, 5, 5, 1}}, []int{}},
		{"T5", args{[]int{2, 2}}, []int{2}},
		{"T6", args{[]int{2, 2, 2}}, []int{2}},
		{"T7", args{[]int{2, 3, 4}}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
