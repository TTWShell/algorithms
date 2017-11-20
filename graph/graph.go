package graph

import (
	"errors"
	"sync"
)

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

// MakeEdge creates  an edge in the graph with a corresponding weight.
// It returns an error if either of the nodes do not belong in the graph.
func (g *Graph) MakeEdge(from, to *Node, weight float32) error {
	g.Lock()
	defer g.Unlock()

	if _, ok := g.nodes[from]; !ok {
		return errors.New("`from` node does not belong to this graph")
	}
	if _, ok := g.nodes[to]; !ok {
		return errors.New("`to` node does not belong to this graph")
	}

	g.edges[from][to] = weight
	return nil
}
