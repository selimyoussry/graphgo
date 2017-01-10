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
func (graph *Graph) AddNodeLegacyIndex(legacyIndex, index string) {
	li := fmt.Sprintf("legacy.%s", legacyIndex)
	graph.LegacyIndex.Nodes[li] = index
	graph.LegacyIndex.Nodes[index] = li
}

// AddEdgeIndex
func (graph *Graph) AddEdgeLegacyIndex(legacyIndex, index string) {
	li := fmt.Sprintf("legacy.%s", legacyIndex)
	graph.LegacyIndex.Edges[li] = index
	graph.LegacyIndex.Edges[index] = li
}

// FindNodeFromLegacy
func (graph *Graph) FindNodeFromLegacyIndex(legacyIndex string) (string, error) {

	li := fmt.Sprintf("legacy.%s", legacyIndex)
	nodeKey, ok := graph.LegacyIndex.Nodes[li]
	if !ok {
		return "", errNoNodeLegacyRecord(legacyIndex)
	}

	return nodeKey, nil
}

// FindEdgeFromLegacyIndex
func (graph *Graph) FindEdgeFromLegacyIndex(legacyIndex string) (string, error) {

	li := fmt.Sprintf("legacy.%s", legacyIndex)
	edgeKey, ok := graph.LegacyIndex.Edges[li]
	if !ok {
		return "", errNoEdgeLegacyRecord(legacyIndex)
	}

	return edgeKey, nil
}
