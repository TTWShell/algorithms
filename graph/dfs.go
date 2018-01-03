package graph

// DFS Depth-First-Search.
func (g *Graph) DFS(node *Node) []*Node {
	visited := make(map[*Node]bool, len(g.nodes))
	res := make([]*Node, 0, len(g.nodes))

	var helper func(node *Node)
	helper = func(node *Node) {
		if _, ok := visited[node]; ok {
			return
		}
		res = append(res, node)
		visited[node] = true
		for _, cur := range g.Neighbors(node) {
			helper(cur)
		}
	}

	helper(node)
	return res
}
