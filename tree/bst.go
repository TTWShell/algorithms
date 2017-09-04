package tree

/*
二叉查找树（Binary Search Tree）是指一棵空树或者具有下列性质的二叉树：
	1. 若任意节点的左子树不空，则左子树上所有节点的值均小于它的根节点的值；
	2. 若任意节点的右子树不空，则右子树上所有节点的值均大于它的根节点的值；
	3. 任意节点的左、右子树也分别为二叉查找树；
	4. 没有键值相等的节点。
*/
type BST struct {
	root *TreeNode
}

// BST Construtor
func NewBST() *BST {
	return &BST{}
}

func insert(root *TreeNode, el int) bool {
	if root.Val == el {
		return false
	} else if root.Val > el {
		if root.Left == nil {
			root.Left = &TreeNode{Val: el}
		} else {
			return insert(root.Left, el)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{Val: el}
		} else {
			return insert(root.Right, el)
		}
	}
	return true
}

// 在二叉搜索树插入节点
func (b *BST) Insert(el int) bool {
	if b.root == nil {
		b.root = &TreeNode{Val: el}
		return true
	}
	return insert(b.root, el)
}

func search(root *TreeNode, el int) bool {
	if root == nil {
		return false
	}
	if root.Val == el {
		return true
	} else if root.Val > el {
		return search(root.Left, el)
	}
	return search(root.Right, el)
}

// 在二叉搜索树查找一个节点
func (b *BST) Search(el int) bool {
	if b.root == nil {
		return false
	}
	return search(b.root, el)
}
