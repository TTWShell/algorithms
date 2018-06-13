package tree

import (
	"fmt"
	"log"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	return fmt.Sprintf("<%s %d %s>", t.Left, t.Val, t.Right)
}

// 广度优先遍历
func BFS(root *TreeNode) {
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		tmp := []*TreeNode{}
		for _, node := range stack {
			log.Printf("%d ", node.Val)
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		stack = tmp
	}
}
