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

func insertBST(root *TreeNode, el int) bool {
	if root.Val == el {
		return false
	} else if root.Val > el {
		if root.Left == nil {
			root.Left = &TreeNode{Val: el}
		} else {
			return insertBST(root.Left, el)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{Val: el}
		} else {
			return insertBST(root.Right, el)
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
	return insertBST(b.root, el)
}

func searchBST(root *TreeNode, el int) bool {
	if root == nil {
		return false
	}
	if root.Val == el {
		return true
	} else if root.Val > el {
		return searchBST(root.Left, el)
	}
	return searchBST(root.Right, el)
}

// 在二叉搜索树查找一个节点
func (b *BST) Search(el int) bool {
	if b.root == nil {
		return false
	}
	return searchBST(b.root, el)
}

func deleteBST(node, parent *TreeNode, isLeft bool, el int) bool {
	switch {
	case node.Val == el:
		if node.Left != nil && node.Right != nil {
			// 左右子树均不空
			q, tmp := node, node.Left
			for tmp.Right != nil {
				q, tmp = tmp, tmp.Right
			}
			node.Val = tmp.Val
			if q.Val != node.Val {
				q.Right = tmp.Left
			} else {
				q.Left = tmp.Left
			}
			return true
		}

		var tmpRes *TreeNode
		// 该节点为叶子节点，直接删除
		if node.Left == nil && node.Right == nil {
			tmpRes = nil
		} else if node.Left == nil {
			// 左子树空只需重接它的右子树
			tmpRes = node.Right
		} else if node.Right == nil {
			// 右子树空则只需重接它的左子树
			tmpRes = node.Left
		}

		if isLeft {
			parent.Left = tmpRes
		} else {
			parent.Right = tmpRes
		}
		return true
	case node.Val > el:
		if node.Left != nil {
			return deleteBST(node.Left, node, true, el)
		}
	case node.Val < el:
		if node.Right != nil {
			return deleteBST(node.Right, node, false, el)
		}
	}
	return false
}

// 在二叉查找树删除结点
func (b *BST) Delete(el int) bool {
	if b.root == nil {
		return false
	}
	if b.root.Val == el {
		tmp := &TreeNode{Val: 0}
		tmp.Left = b.root
		res := deleteBST(b.root, tmp, true, el)
		b.root = tmp.Left
		return res
	} else if b.root.Val > el {
		return deleteBST(b.root.Left, b.root, true, el)
	}
	return deleteBST(b.root.Right, b.root, false, el)
}
