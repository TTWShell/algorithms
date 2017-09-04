package tree

import "sync"

/*
AVL(Adelson-Velskii & Landis)树是带有平衡条件的二叉查找树(BST)，其每个节点的左子树和右子树的高度最多差1。
*/
type AVL struct {
	sync.Mutex
	root *TreeNode
}
