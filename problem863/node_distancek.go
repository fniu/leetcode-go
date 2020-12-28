package problem863

// TreeNode is a struct type of a tree node.
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
func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	solution := []int{}
	// K steps downwards starting from 'target'.
	dfsDepthValues(target, K, &solution)
	if target == root {
		return solution
	}
	nodesToTarget := dfsToTarget(root, target, []*TreeNode{})
	index := len(nodesToTarget) - 1
	for i := 1; i <= K; i++ {
		index--
		if nodesToTarget[index].Left != nil && nodesToTarget[index].Left != nodesToTarget[index+1] {
			dfsDepthValues(nodesToTarget[index].Left, K-i-1, &solution)
		}
		if nodesToTarget[index].Right != nil && nodesToTarget[index].Right != nodesToTarget[index+1] {
			dfsDepthValues(nodesToTarget[index].Right, K-i-1, &solution)
		}
		if i == K {
			solution = append(solution, nodesToTarget[index].Val)
		}
		if index == 0 {
			break
		}
	}
	return solution
}

// DFS to find target in the tree.
func dfsToTarget(root, target *TreeNode, parents []*TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	parents = append(parents, root)
	if root.Left == target {
		return append(parents, root.Left)
	}
	if root.Right == target {
		return append(parents, root.Right)
	}
	left := dfsToTarget(root.Left, target, parents)
	if left != nil {
		return left
	}
	return dfsToTarget(root.Right, target, parents)
}

// DFS to find the node values at depth.
func dfsDepthValues(root *TreeNode, depth int, nodes *[]int) {
	if root == nil {
		return
	}
	if depth == 0 {
		*nodes = append(*nodes, root.Val)
		return
	}
	dfsDepthValues(root.Left, depth-1, nodes)
	dfsDepthValues(root.Right, depth-1, nodes)
}
