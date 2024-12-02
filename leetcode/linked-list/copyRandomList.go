/* https://leetcode.com/problems/copy-list-with-random-pointer/?envType=study-plan-v2&envId=top-interview-150
 */

package lll

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// Step 1: Insert copy nodes into original list
	cur := head
	for cur != nil {
		newNode := &Node{Val: cur.Val, Next: cur.Next}
		cur.Next = newNode
		cur = newNode.Next
	}

	// Step 2: Assign random pointers for copied nodes
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	// Step 3: Separate original and copied lists
	oldList := head
	newList := head.Next
	newHead := newList
	for oldList != nil {
		oldList.Next = oldList.Next.Next
		if newList.Next != nil {
			newList.Next = newList.Next.Next
		}
		oldList = oldList.Next
		newList = newList.Next
	}

	return newHead
}
