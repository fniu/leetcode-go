// https://leetcode.com/problems/count-nodes-equal-to-average-of-subtree/

package problem2265

import (
	"reflect"
	"testing"
)

func Test_averageOfSubtree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Test 1", args{&TreeNode{4, nil, nil}}, 1},
		{"Test 2", args{&TreeNode{4, &TreeNode{8, &TreeNode{Val: 0}, &TreeNode{Val: 1}}, &TreeNode{5, nil, &TreeNode{Val: 6}}}}, 5},
		{"Test 3", args{&TreeNode{4, &TreeNode{Val: 1}, &TreeNode{Val: 7}}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageOfSubtree(tt.args.root); got != tt.want {
				t.Errorf("averageOfSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_treeFromList(t *testing.T) {
	type args struct {
		l []int
	}
	tests := []struct {
		name string
		args args
		want TreeNode
	}{
		// TODO: Add test cases.
		{"Test 1", args{[]int{4, 8, 5}}, TreeNode{4, &TreeNode{8, nil, nil}, &TreeNode{5, nil, nil}}},
		{"Test 2", args{[]int{4, 8, 5, 0, 1, 50, -1}}, TreeNode{4, &TreeNode{8, &TreeNode{0, nil, nil}, &TreeNode{1, nil, nil}}, &TreeNode{5, &TreeNode{Val: 50}, nil}}},
		{"Test 3", args{[]int{4, 8, -1, 0, 1}}, TreeNode{4, &TreeNode{8, &TreeNode{0, nil, nil}, &TreeNode{1, nil, nil}}, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := treeFromList(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("treeFromList() = %v, want %v", got, tt.want)
			}
		})
	}
}
