package tree

// PreOrderRecursion : 前序遍历的递归实现。
func PreOrderRecursion(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	res = append(res, root.Val)
	res = append(res, PreOrderRecursion(root.Left)...)
	res = append(res, PreOrderRecursion(root.Right)...)
	return res
}

// PreOrderStack : 前序遍历的非递归实现。
func PreOrderStack(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}

	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			res = append(res, root.Val)
			root = root.Left
		}

		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = top.Right
	}
	return res
}
