package graphgo

import "github.com/hippoai/goerr"

// Edge has a unique key, properties, start and end node
type Edge struct {
	Key   string                 `json:"key"`
	Label string                 `json:"label"`
	Props map[string]interface{} `json:"props"`
	Start string                 `json:"start"`
	End   string                 `json:"end"`
}

// NewEdge instanciates
func NewEdge(key, label string, start, end string, props map[string]interface{}) *Edge {
	return &Edge{
		Key:   key,
		Label: label,
		Props: props,
		Start: start,
		End:   end,
	}
}

// SetProperty one by one
func (edge *Edge) SetProperty(key string, value interface{}) {
	edge.Props[key] = value
}

// Get a property
func (edge *Edge) Get(key string) (interface{}, *goerr.Err) {
	value, ok := edge.Props[key]
	if !ok {
		return nil, errorEdgePropNotFound(edge.Key, key)
	}

	return value, nil
}

// Hop returns the other node
func (edge *Edge) Hop(graph *Graph, key string) (*Node, *goerr.Err) {

	otherNodeKey := edge.Start
	if otherNodeKey == key {
		otherNodeKey = edge.End
	}

	node, err := graph.GetNode(otherNodeKey)
	if err != nil {
		return nil, err
	}

	return node, nil

}

// StartN returns the start node
func (edge *Edge) StartN(graph *Graph) (*Node, *goerr.Err) {
	return edge.Hop(graph, edge.End)
}

// EndN returns the end node
func (edge *Edge) EndN(graph *Graph) (*Node, *goerr.Err) {
	return edge.Hop(graph, edge.Start)
}

// GetLabel returns the label
func (edge *Edge) GetLabel() string {
	return edge.Label
}

// GetKey returns the key
func (edge *Edge) GetKey() string {
	return edge.Key
}
