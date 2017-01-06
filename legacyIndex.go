package graphgo

import "fmt"

// LegacyIndex
// Stores maps for Neo4j indexes
type LegacyIndex struct {
	Nodes map[string]string `json:"nodes"`
	Edges map[string]string `json:"edges"`
}

// NewLegacyIndex instanciates
func NewLegacyIndex() *LegacyIndex {
	return &LegacyIndex{
		Nodes: map[string]string{},
		Edges: map[string]string{},
	}
}

// AddNodeIndex
func (graph *Graph) AddNodeIndex(legacyIndex, index string) {
	li := fmt.Sprintf("legacy.%s", legacyIndex)
	graph.LegacyIndex.Nodes[li] = index
	graph.LegacyIndex.Nodes[index] = li
}

// AddEdgeIndex
func (graph *Graph) AddEdgeIndex(legacyIndex, index string) {
	li := fmt.Sprintf("legacy.%s", legacyIndex)
	graph.LegacyIndex.Edges[li] = index
	graph.LegacyIndex.Edges[index] = li
}
