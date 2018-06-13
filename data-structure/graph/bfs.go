package graph

// BFS Breadth-First-Search.
func (g *Graph) BFS(node *Node) []*Node {
	visited := make(map[*Node]bool, len(g.nodes))
	res := make([]*Node, 0, len(g.nodes))

	queue := []*Node{node}
	for len(queue) != 0 {
		newQueue := []*Node{}
		for _, cur := range queue {
			if _, ok := visited[cur]; ok {
				continue
			}
			res = append(res, cur)
			visited[cur] = true
			for _, next := range g.Neighbors(cur) {
				if _, ok := visited[next]; !ok {
					newQueue = append(newQueue, next)
				}
			}
		}
		queue = newQueue
	}
	return res
}
