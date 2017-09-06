package tree

import (
	"fmt"
	"sync"
)

type avlNode struct {
	Val         int
	Height      int
	Left, Right *avlNode
}

func (t *avlNode) String() string {
	return fmt.Sprintf("<%s %d-%d %s>", t.Left, t.Val, t.Height, t.Right)
}

/*
AVL(Adelson-Velskii & Landis)树是带有平衡条件的二叉查找树(BST)，其每个节点的左子树和右子树的高度最多差1。
查找、插入和删除在平均和最坏情况下的时间复杂度都是O(log(n))。
*/
type AVL struct {
	sync.Mutex
	root *avlNode
}

// AVL Construtor
func NewAVL() *AVL {
	return &AVL{}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func height(root *avlNode) int {
	if root != nil {
		return root.Height
	}
	return 0
}

// https://upload.wikimedia.org/wikipedia/commons/c/c7/Tree_Rebalancing.png
// 左左情况，右旋
func rightRotate(root *avlNode) *avlNode {
	node := root.Left
	root.Left = node.Right
	node.Right = root

	root.Height = max(height(root.Left), height(root.Right)) + 1
	node.Height = max(height(node.Left), height(node.Right)) + 1
	return node
}

// 右右情况，左旋
func leftRotate(root *avlNode) *avlNode {
	node := root.Right
	root.Right = node.Left
	node.Left = root

	root.Height = max(height(root.Left), height(root.Right)) + 1
	node.Height = max(height(node.Left), height(node.Right)) + 1
	return node
}

// 左右情况，先左旋root.Left再右旋root
func leftRigthRotate(root *avlNode) *avlNode {
	root.Left = leftRotate(root.Left)
	root = rightRotate(root)
	return root
}

// 右左情况，先右旋root.Right再左旋root
func rightLeftRotate(root *avlNode) *avlNode {
	root.Right = rightRotate(root.Right)
	root = leftRotate(root)
	return root
}

func insertAVL(root *avlNode, el int) (res bool, node *avlNode) {
	switch {
	case root.Val == el:
		return false, nil
	case root.Val > el:
		if root.Left == nil {
			root.Left = &avlNode{Val: el}
		} else {
			res, node = insertAVL(root.Left, el)
			if !res {
				return false, nil
			}
			root.Left = node
			if height(root.Left)-height(root.Right) == 2 {
				root = rightRotate(root) // 左左
			} else {
				root = leftRigthRotate(root) // 左右
			}
		}
	default:
		if root.Right == nil {
			root.Right = &avlNode{Val: el}
		} else {
			res, node = insertAVL(root.Right, el)
			if !res {
				return false, nil
			}
			root.Right = node
			if height(root.Right)-height(root.Left) == 2 {
				if el > root.Right.Val {
					root = leftRotate(root) // 右右
				} else {
					root = rightLeftRotate(root) // 右左
				}
			}
		}
	}
	root.Height = max(height(root.Left), height(root.Right)) + 1
	return true, root
}

func (avl *AVL) Insert(el int) bool {
	avl.Lock()
	defer avl.Unlock()

	if avl.root == nil {
		avl.root = &avlNode{Val: el}
		return true
	}
	if res, root := insertAVL(avl.root, el); res {
		avl.root = root
		return true
	}
	return false
}
