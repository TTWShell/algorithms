/* https://leetcode.com/problems/n-ary-tree-postorder-traversal/
Given the root of an n-ary tree, return the postorder traversal of its nodes' values.

Nary-Tree input serialization is represented in their level order traversal. Each group of children is separated by the null value (See examples)

Example 1:

Input: root = [1,null,3,2,4,null,5,6]
Output: [5,6,3,2,4,1]
Example 2:

Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
Output: [2,6,14,11,7,3,12,8,4,13,9,10,5,1]

Constraints:

The number of nodes in the tree is in the range [0, 10^4].
0 <= Node.val <= 10^4
The height of the n-ary tree is less than or equal to 1000.
*/

package ltree

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	// var helper func(root *Node, ans *[]int)
	// helper = func(root *Node, ans *[]int) {
	// 	if root == nil {
	// 		return
	// 	}
	// 	for _, child := range root.Children {
	// 		helper(child, ans)
	// 	}
	// 	*ans = append(*ans, root.Val)
	// }

	// ans := []int{}
	// helper(root, &ans)
	// return ans

	ans := []int{}
	stack := []*Node{root}
	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		ans = append(ans, cur.Val)
		stack = stack[:len(stack)-1]
		for _, child := range cur.Children {
			stack = append(stack, child)
		}
	}

	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}

	return ans
}
