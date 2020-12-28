package problem863

func serializeToTarget(root, target *TreeNode) []*TreeNode {
	if target == nil {
		return []*TreeNode{}
	}
	if root == target {
		return []*TreeNode{root}
	}
	s := []*TreeNode{}
	queue := []*TreeNode{root}
	var head *TreeNode
	for len(queue) > 0 {
		head = queue[0]
		s = append(s, head)
		if head != nil {
			if head == target {
				return s
			}
			queue = append(queue, head.Left)
			queue = append(queue, head.Right)
		}
		queue = queue[1:]
	}

	return []*TreeNode{}

}
