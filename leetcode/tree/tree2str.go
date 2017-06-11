/* https://leetcode.com/problems/construct-string-from-binary-tree/#/description
You need to construct a string consists of parenthesis and integers from a binary tree with the preorder traversing way.

The null node needs to be represented by empty parenthesis pair "()".
And you need to omit all the empty parenthesis pairs that don't affect the one-to-one mapping relationship between the string and the original binary tree.

Example 1:
    Input: Binary tree: [1,2,3,4]
           1
         /   \
        2     3
       /
      4

    Output: "1(2(4))(3)"

    Explanation: Originallay it needs to be "1(2(4)())(3()())",
    but you need to omit all the unnecessary empty parenthesis pairs.
    And it will be "1(2(4))(3)".

Example 2:
    Input: Binary tree: [1,2,3,null,4]
           1
         /   \
        2     3
         \
          4

    Output: "1(2()(4))(3)"

    Explanation: Almost the same as the first example,
    except we can't omit the first parenthesis pair to break the one-to-one mapping relationship between the input and the output.
*/

package leetcode

import (
	"bytes"
	"strconv"
)

func tree2str(t *TreeNode) string {
	buffer := new(bytes.Buffer)

	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}

		buffer.WriteString(strconv.Itoa(node.Val))
		if node.Left == nil && node.Right == nil {
			return
		}

		buffer.WriteByte('(')
		helper(node.Left)
		buffer.WriteByte(')')

		if node.Right != nil {
			buffer.WriteByte('(')
			helper(node.Right)
			buffer.WriteByte(')')
		}
	}

	helper(t)

	return buffer.String()
}

/*
import "fmt"

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}

	if t.Left == nil && t.Right == nil {
		return fmt.Sprintf("%d", t.Val)
	}

	left, right := "", ""
	left = fmt.Sprintf("(%s)", tree2str(t.Left))

	if t.Right != nil {
		right = fmt.Sprintf("(%s)", tree2str(t.Right))
	}

	return fmt.Sprintf("%d%s%s", t.Val, left, right)
}
*/
