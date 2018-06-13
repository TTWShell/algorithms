package tree

// PostOrderRecursion : 后序遍历的递归实现。
func PostOrderRecursion(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	res = append(res, PostOrderRecursion(root.Left)...)
	res = append(res, PostOrderRecursion(root.Right)...)
	res = append(res, root.Val)
	return res
}

// PostOrderStack : 后序遍历的非递归实现。
func PostOrderStack(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}

	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			res = append([]int{root.Val}, res...)
			root = root.Right
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = top.Left
	}
	return res
}

/*
func PostOrderStack(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{root}
	var prev, cur *TreeNode

	for len(stack) != 0 {
		cur = stack[len(stack)-1]
		if (cur.Right == nil && cur.Left == nil) ||
			(cur.Right != nil && cur.Right == prev) ||
			(cur.Right == nil && cur.Left == prev) {
			stack = stack[:len(stack)-1]
			res = append(res, cur.Val)
			prev = cur
		} else {
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
		}
	}
	return res
}
*/
