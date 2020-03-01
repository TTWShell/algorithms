package unionfind

type ufNode struct {
	value            interface{}
	ancestor, parent *ufNode
	children         []*ufNode
}

// UF : Union-Find.
type UF struct {
	// map all value to node and, for find method
	nodes map[interface{}]*ufNode
	// save trees, for count method, int value is count of nodes in set
	ufSets map[*ufNode]int
}

// New : create a empty Union-Find.
func New() *UF {
	return &UF{
		nodes:  map[interface{}]*ufNode{},
		ufSets: map[*ufNode]int{},
	}
}

// Find : return ancestor's value.
func (uf *UF) Find(value interface{}) interface{} {
	if ufN, ok := uf.nodes[value]; ok {
		return ufN.ancestor.value
	}
	return nil
}

// Connected : return true if p ∈ setA && q ∈ setA else false.
func (uf *UF) Connected(p, q interface{}) bool {
	pAnc, qAnc := uf.Find(p), uf.Find(q)
	if pAnc != nil && qAnc != nil && pAnc == qAnc {
		return true
	}
	return false
}

// Count : count of disjoint sets.
func (uf *UF) Count() int {
	return len(uf.ufSets)
}

// SetCount : value
func (uf *UF) SetCount() map[interface{}]int {
	res := make(map[interface{}]int, uf.Count())
	for node, count := range uf.ufSets {
		res[node.value] = count
	}
	return res
}

// transposition: node --> root & set ancestor.
func (uf *UF) transposition(node, newAnc *ufNode) {
	ago, cur, anc := node, node, node.ancestor
	queue := []*ufNode{anc}

	for len(queue) != 0 {
		nextQueue := []*ufNode{}
		for _, p := range queue {
			p.ancestor = newAnc
			nextQueue = append(nextQueue, p.children...)
		}
		queue = nextQueue
	}

	for cur != anc {
		cIdx, parent := -1, cur.parent
		for i, n := range parent.children {
			if n == cur {
				cIdx = i
				break
			}
		}
		parent.children = append(
			append([]*ufNode{}, parent.children[0:cIdx]...),
			parent.children[cIdx+1:]...)
		cur.children = append(cur.children, parent)
		ago, cur.parent, cur = cur, ago, parent
	}

	cur.parent = ago
}

// assume qNode merge to pNode.
func (uf *UF) union(pNode, qNode *ufNode) {
	pAnc, qAnc := pNode.ancestor, qNode.ancestor
	uf.ufSets[pAnc] += uf.ufSets[qAnc]
	delete(uf.ufSets, qAnc)

	// transposition qNode tree.
	uf.transposition(qNode, pAnc)

	// connect pNode and qNode.
	pNode.children = append(pNode.children, qNode)
	qNode.parent = pNode
}

func (uf *UF) getOrCreateNode(value interface{}) *ufNode {
	node, ok := uf.nodes[value]
	if !ok {
		node = &ufNode{value: value, children: []*ufNode{}}
		node.ancestor = node
		node.parent = node
		uf.nodes[value] = node
		uf.ufSets[node] = 1
	}
	return node
}

// Union : add a connection for p and q.
func (uf *UF) Union(p, q interface{}) {
	if uf.Connected(p, q) {
		return
	}

	if p == q {
		uf.getOrCreateNode(p)
		return
	}

	pNode, qNode := uf.getOrCreateNode(p), uf.getOrCreateNode(q)
	countP, _ := uf.ufSets[pNode.ancestor]
	countQ, _ := uf.ufSets[qNode.ancestor]
	if countP >= countQ {
		uf.union(pNode, qNode)
	} else {
		uf.union(qNode, pNode)
	}
}
