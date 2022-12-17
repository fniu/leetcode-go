// https://leetcode.com/problems/count-nodes-equal-to-average-of-subtree/

package problem2265

import "math"

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

	if root.Left != nil {
		nic := sumOfSubtree(root.Left)
		ni.num += nic.num
		ni.sum += nic.sum
	}
	if root.Right != nil {
		nic := sumOfSubtree(root.Right)
		ni.num += nic.num
		ni.sum += nic.sum
	}
	if int(math.Floor(float64(ni.sum)/float64(ni.num))) == root.Val {
		globalRes += 1
	}
	return ni
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
