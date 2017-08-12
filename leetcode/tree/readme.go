package tree

import "log"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DFSPreOrder(root *TreeNode) {
	if root != nil {
		log.Printf("%d ", root.Val)
		DFSPreOrder(root.Left)
		DFSPreOrder(root.Right)
	}
}

func DFSInOrder(root *TreeNode) {
	if root != nil {
		DFSInOrder(root.Left)
		log.Printf("%d ", root.Val)
		DFSInOrder(root.Right)
	}
}
