package problem228

import (
	"reflect"
	"testing"
)

func Test_summaryRanges(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"T1", args{[]int{0, 1, 2, 4, 5, 7}}, []string{"0->2", "4->5", "7"}},
		{"T2", args{[]int{0, 2, 3, 4, 6, 8, 9}}, []string{"0", "2->4", "6", "8->9"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := summaryRanges(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("summaryRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
