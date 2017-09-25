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
	sync.RWMutex
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
	return -1
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
				if el < root.Left.Val {
					root = rightRotate(root) // 左左
				} else {
					root = leftRigthRotate(root) // 左右
				}
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

func searchAVL(root *avlNode, el int) bool {
	if root == nil {
		return false
	}
	if root.Val == el {
		return true
	} else if root.Val > el {
		return searchAVL(root.Left, el)
	}
	return searchAVL(root.Right, el)
}

func (avl *AVL) Search(el int) bool {
	avl.RLock()
	defer avl.RUnlock()

	if avl.root == nil {
		return false
	}
	return searchAVL(avl.root, el)
}

// 找到最小的叶子节点
func minLeaf(root *avlNode) (leaf *avlNode) {
	leaf = root
	for leaf.Left != nil || leaf.Right != nil {
		if leaf.Left != nil {
			leaf = leaf.Left
		} else if leaf.Right.Left != nil {
			leaf = leaf.Right
		} else {
			break
		}
	}
	return
}

// 找到最大的叶子节点
func maxLeaf(root *avlNode) (leaf *avlNode) {
	leaf = root
	for leaf.Left != nil || leaf.Right != nil {
		if leaf.Right != nil {
			leaf = leaf.Right
		} else if leaf.Left.Right != nil {
			leaf = leaf.Left
		} else {
			break
		}
	}
	return
}

func deleteAVL(root *avlNode, el int) (res bool, node *avlNode) {
	if root == nil {
		return false, root
	}

	var isSwap bool
	if root.Val == el {
		if root.Left == nil && root.Right == nil {
			return true, nil
		}
		var leaf *avlNode
		isSwap = true
		if height(root.Left) > height(root.Right) {
			leaf = root.Left
			// find max in Left
			leaf = maxLeaf(root.Left)
		} else {
			// find min in Right
			leaf = minLeaf(root.Right)
		}
		el = leaf.Val
	}
	switch {
	case root.Val > el:
		if isSwap {
			root.Val = el
		}
		res, node = deleteAVL(root.Left, el)
		if !res {
			return false, root
		}
		root.Left = node
		if height(root.Right)-height(root.Left) == 2 {
			if height(root.Right.Left) < height(root.Right.Right) {
				root = leftRotate(root) // 右右
			} else {
				root = rightLeftRotate(root) // 右左
			}
		}
	case root.Val < el:
		if isSwap {
			root.Val = el
		}
		res, node = deleteAVL(root.Right, el)
		if !res {
			return false, root
		}
		root.Right = node
		if height(root.Left)-height(root.Right) == 2 {
			if height(root.Left.Left) > height(root.Left.Right) {
				root = rightRotate(root) // 左左
			} else {
				root = leftRigthRotate(root) // 左右
			}
		}
	}
	root.Height = max(height(root.Left), height(root.Right)) + 1
	return true, root
}

func (avl *AVL) Delete(el int) bool {
	avl.Lock()
	defer avl.Unlock()

	if avl.root == nil {
		return false
	}
	res, root := deleteAVL(avl.root, el)
	if res {
		avl.root = root
	}
	return res
}
