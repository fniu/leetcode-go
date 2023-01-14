// https://leetcode.com/problems/kth-smallest-element-in-a-bst/

package problem230

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
func kthSmallest(root *TreeNode, k int) int {

	s := 0
	return inOrderTraverse(root, &s, k)
}
func inOrderTraverse(root *TreeNode, s *int, k int) int {
	var val int
	if root.Left != nil {
		val = inOrderTraverse(root.Left, s, k)
	}
	if val != 0 {
		return val
	}
	*s++
	if *s == k {
		return root.Val
	}

	if root.Right != nil {
		val = inOrderTraverse(root.Right, s, k)
	}

	return val
}
