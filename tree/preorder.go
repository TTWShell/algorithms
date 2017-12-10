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
