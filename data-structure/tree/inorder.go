package tree

// InOrderRecursion : 中序遍历的递归实现。
func InOrderRecursion(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	res = append(res, InOrderRecursion(root.Left)...)
	res = append(res, root.Val)
	res = append(res, InOrderRecursion(root.Right)...)
	return res
}

// InOrderStack : 中序遍历的非递归实现。
func InOrderStack(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}

	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, top.Val)
		root = top.Right

	}
	return res
}
