/* https://leetcode.com/problems/binary-search-tree-iterator/
Implement an iterator over a binary search tree (BST).
Your iterator will be initialized with the root node of a BST.

Calling next() will return the next smallest number in the BST.

Example:

https://assets.leetcode.com/uploads/2018/12/25/bst-tree.png

	BSTIterator iterator = new BSTIterator(root);
	iterator.next();    // return 3
	iterator.next();    // return 7
	iterator.hasNext(); // return true
	iterator.next();    // return 9
	iterator.hasNext(); // return true
	iterator.next();    // return 15
	iterator.hasNext(); // return true
	iterator.next();    // return 20
	iterator.hasNext(); // return false

Note:

	next() and hasNext() should run in average O(1) time and
	uses O(h) memory, where h is the height of the tree.
	You may assume that next() call will always be valid,
	that is, there will be at least a next smallest number in the BST
	 when next() is called.
*/

package ldesign

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
}

func BSTIteratorConstructor(root *TreeNode) BSTIterator {
	bst := BSTIterator{stack: []*TreeNode{}}

	bst.inOrder(root)
	return bst
}

func (this *BSTIterator) inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	this.stack = append(this.stack, root)
	this.inOrder(root.Left)
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	last := len(this.stack) - 1
	next := this.stack[last]
	this.stack = this.stack[:last]

	if next.Right != nil {
		this.inOrder(next.Right)
	}
	return next.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if len(this.stack) != 0 {
		return true
	}
	return false
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
