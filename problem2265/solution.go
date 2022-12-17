// https://leetcode.com/problems/count-nodes-equal-to-average-of-subtree/

package problem2265

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
var globalRes = 0

func averageOfSubtree(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 1
	}
	globalRes = 0
	sumOfSubtree(root)
	return globalRes
}

type NodeInfo struct {
	num, sum int
}

func sumOfSubtree(root *TreeNode) *NodeInfo {

	ni := &NodeInfo{1, root.Val}
	var nic *NodeInfo
	if root.Left != nil {
		nic = sumOfSubtree(root.Left)
		ni.num += nic.num
		ni.sum += nic.sum
	}
	if root.Right != nil {
		nic = sumOfSubtree(root.Right)
		ni.num += nic.num
		ni.sum += nic.sum
	}
	if ni.sum/ni.num == root.Val {
		globalRes += 1
	}
	return ni
}
*/

func sumCount(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	a, b := sumCount(root.Left)
	a2, b2 := sumCount(root.Right)
	return root.Val + a + a2, 1 + b + b2
}

func averageOfSubtree(root *TreeNode) int {
	var result int

	if root == nil {
		return 0
	}

	result = averageOfSubtree(root.Left) + averageOfSubtree(root.Right)
	sum, count := sumCount(root)
	if (sum / count) == root.Val {
		return result + 1
	}
	return result
}

// Non-recursive
func treeFromList(l []int) TreeNode {

	if len(l) == 1 {
		return TreeNode{Val: l[0], Left: nil, Right: nil}
	}
	var treeList []*TreeNode
	treeList = append(treeList, &TreeNode{l[0], nil, nil})
	i, j := 1, 0
	for i < len(l) {
		if j >= len(l) {
			break
		}
		if l[i] >= 0 {
			treeList[j].Left = &TreeNode{Val: l[i]}
			treeList = append(treeList, treeList[j].Left)
		}
		i += 1
		if l[i] >= 0 {
			treeList[j].Right = &TreeNode{Val: l[i]}
			treeList = append(treeList, treeList[j].Right)
		}
		i += 1
		j += 1
	}
	return *treeList[0]
}
