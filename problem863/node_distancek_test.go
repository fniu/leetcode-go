package problem863

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestSolution(t *testing.T) {
	var r *TreeNode

	r = BuildTreeNodeFromArray(
		[]int{3,
			2, 0,
			6, 99, 100, 101,
			-1, 7, -1, 22, -1, 33, 23, -1})

	assertEquals(
		t,
		[]int{100, 101},
		distanceK(r, r.Left.Right, 4))

}

func assertEquals(t *testing.T, expected, got []int) {
	t.Helper()
	sort.Ints(expected)
	sort.Ints(got)
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected %v, got=%v", expected, got)
	}
}

func TestTreeNodeBuilder(t *testing.T) {

	var a []int

	a = []int{
		1,
		2, 0,
		6, 1, 100, -1,
		-1, 7, -1, 22, -1, 33, -1, -1}

	root := BuildTreeNodeFromArray(a)
	fmt.Printf("Tree Node: %v", root)
	fmt.Printf("Tree Node: %v", root.Left)
	fmt.Printf("Tree Node: %v\n", root.Right)
	fmt.Printf("Tree Node: %v", root.Left.Left)
	fmt.Printf("Tree Node: %v", root.Left.Right)
	fmt.Printf("Tree Node: %v\n", root.Right.Left)
	fmt.Printf("Tree Node: %v", root.Right.Left.Left)
	fmt.Printf("Tree Node: %v", root.Left.Left.Right)
	fmt.Printf("Tree Node: %v\n", root.Left.Right.Right)
}

func BuildTreeNodeFromArray(treeArr []int) *TreeNode {
	n := len(treeArr)

	if n < 3 {
		panic(fmt.Sprintf("The array is too short, %v", treeArr))
	}
	root := &TreeNode{Val: treeArr[0]}
	var queue []*TreeNode
	queue = append(queue, root)
	var left, right *TreeNode
	for i := 0; i < n; i++ {
		if queue[0] == nil {
			queue = queue[1:]
			continue
		}
		if 2*(i+1) < n {
			if treeArr[2*(i+1)-1] == -1 {
				left = nil
			} else {
				left = &TreeNode{Val: treeArr[2*(i+1)-1]}
			}
			if treeArr[2*(i+1)] == -1 {
				right = nil
			} else {
				right = &TreeNode{Val: treeArr[2*(i+1)]}
			}
			queue = append(queue, left)
			queue = append(queue, right)
		} else {
			break
		}

		queue[0].Left = left
		queue[0].Right = right
		queue = queue[1:]
	}
	return root
}

func TestSerialize(t *testing.T) {
	var a []int
	var root *TreeNode
	var s []*TreeNode
	a = []int{
		1,
		2, 0,
		6, 1, 100, -1,
		-1, 7, -1, 22, -1, 33, -1, -1}

	root = BuildTreeNodeFromArray(a)

	s = serializeToTarget(root, root.Left.Right.Right)
	for _, val := range s {
		if val != nil {
			fmt.Print(val.Val)
		} else {
			fmt.Print(nil)
		}
		fmt.Print(" ")
	}
	fmt.Println()

}

func TestDfsToTarget(t *testing.T) {

	var a []int
	var root *TreeNode
	a = []int{
		1,
		2, 0,
		6, 1, 100, -1,
		-1, 7, -1, 22, -1, 33, -1, -1}

	root = BuildTreeNodeFromArray(a)
	pathToTarget := dfsToTarget(root, root.Left.Right, []*TreeNode{})
	for _, v := range pathToTarget {
		t.Logf("%v\n", v)
	}
}

func TestDfsDepthValues(t *testing.T) {

	var a []int
	var root *TreeNode
	a = []int{
		1,
		2, 0,
		6, 1, 100, -1,
		13, 7, -1, 22, -1, 33, -1, -1}

	root = BuildTreeNodeFromArray(a)
	nodes := []int{}
	dfsDepthValues(root.Left, 2, &nodes)
	t.Log(nodes)
}
