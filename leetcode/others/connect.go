package lothers

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	current := root
	for current != nil {
		dummy := &Node{}
		tail := dummy

		for node := current; node != nil; node = node.Next {
			if node.Left != nil {
				tail.Next = node.Left
				tail = tail.Next
			}
			if node.Right != nil {
				tail.Next = node.Right
				tail = tail.Next
			}
		}
		current = dummy.Next
	}

	return root
}
