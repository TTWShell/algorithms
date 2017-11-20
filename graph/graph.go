package graph

import "sync"

// GraphType values.
const (
	Undirected int = iota
	Directed
)

// Node is a Node(node) in Graph.
type Node struct {
	Value interface{}
}

// Graph .
type Graph struct {
	sync.RWMutex
	// Auto set type of graph.
	kind int
	// Bool is used for graph traversal.
	nodes map[*Node]bool
	// Edge connects two Nodes in a graph, float32 is weight of edge.
	edges map[*Node]map[*Node]float32
}

// New creates and returns an empty graph.
func New() *Graph {
	return &Graph{
		kind:  Undirected,
		nodes: make(map[*Node]bool),
		edges: make(map[*Node]map[*Node]float32),
	}
}

// MakeNode creates a node, adds it to the graph and returns the new node.
func (g *Graph) MakeNode(value interface{}) *Node {
	g.Lock()
	defer g.Unlock()

	newNode := &Node{Value: value}
	if _, ok := g.nodes[newNode]; !ok {
		g.nodes[newNode] = false
		g.edges[newNode] = make(map[*Node]float32)
	}
	return newNode
}
