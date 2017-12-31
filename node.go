package graphgo

import "github.com/hippoai/goerr"

// Node implements a graph node
type Node struct {
	Key   string                 `json:"key"`
	Props map[string]interface{} `json:"props"`
	Out   map[string]string      `json:"out"`
	In    map[string]string      `json:"in"`
}

// NewNode instanciates
func NewNode(key string, props map[string]interface{}) *Node {
	return &Node{
		Key:   key,
		Props: props,
		Out:   map[string]string{},
		In:    map[string]string{},
	}
}

// NewEmptyNode instanciates
func NewEmptyNode(key string) *Node {
	return NewNode(key, map[string]interface{}{})
}

// SetProperty to a node
func (node *Node) SetProperty(key string, value interface{}) {
	node.Props[key] = value
}

// AddOutEdge
func (node *Node) AddOutEdge(edge, label string) {
	node.Out[edge] = label
}

// AddInEdge
func (node *Node) AddInEdge(edge, label string) {
	node.In[edge] = label
}

// Get a property
func (node *Node) Get(key string) (interface{}, error) {
	value, ok := node.Props[key]
	if !ok {
		return nil, errorNodePropNotFound(node.Key, key)
	}

	return value, nil
}

// Copy returns a fresh node with same properties
func (node *Node) Copy() *Node {
	props := map[string]interface{}{}
	for key, value := range node.Props {
		props[key] = value
	}

	out := map[string]string{}
	for key, value := range out {
		out[key] = value
	}

	in := map[string]string{}
	for key, value := range out {
		in[key] = value
	}

	return &Node{
		Key:   node.Key,
		Props: props,
		Out:   out,
		In:    in,
	}
}

// InE returns the incoming edges
func (node *Node) InE(graph IGraph, label string) (map[string]IEdge, error) {

	result := map[string]IEdge{}
	missingEdges := []string{}

	// Loop over the edges
	for edgeKey, edgeLabel := range node.In {
		if edgeLabel != label {
			continue
		}

		edge, err := graph.GetEdge(edgeKey)
		if err != nil {
			missingEdges = append(missingEdges, edgeKey)
			continue
		}

		result[edge.GetKey()] = edge

	}

	// Return the result, along with missing edges
	if len(missingEdges) > 0 {
		return result, goerr.New(ERR_NODE_INE_MISSING_EDGES, map[string]interface{}{
			"missingEdges": missingEdges,
			"nodeKey":      node.Key,
		})
	}

	return result, nil
}

// OutE returns the outgoing edges
func (node *Node) OutE(graph IGraph, label string) (map[string]IEdge, error) {

	result := map[string]IEdge{}
	missingEdges := []string{}

	// Loop over the edges
	for edgeKey, edgeLabel := range node.Out {
		if edgeLabel != label {
			continue
		}

		edge, err := graph.GetEdge(edgeKey)
		if err != nil {
			missingEdges = append(missingEdges, edgeKey)
			continue
		}

		result[edge.GetKey()] = edge

	}

	// Return the result, along with missing edges
	if len(missingEdges) > 0 {
		return result, goerr.New(ERR_NODE_OUTE_MISSING_EDGES, map[string]interface{}{
			"missingEdges": missingEdges,
			"nodeKey":      node.Key,
		})
	}

	return result, nil
}

// GetKey returns the key, to implement askgo interface
func (node *Node) GetKey() string {
	return node.Key
}

// GetProps returns all properties
func (node *Node) GetProps() map[string]interface{} {
	return node.Props
}
