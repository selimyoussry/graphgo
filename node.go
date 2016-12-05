package graphgo

// Node implements a graph node
type Node struct {
	Key   string                  `json:"key"`
	Props *map[string]interface{} `json:"props"`
	Out   *map[string]*Edge       `json:"out"`
	In    *map[string]*Edge       `json:"in"`
}

// NewNode instanciates
func NewNode(key string, props *map[string]interface{}) *Node {
	return &Node{
		Key:   key,
		Props: props,
		Out:   &(map[string]*Edge{}),
		In:    &(map[string]*Edge{}),
	}
}

// NewEmptyNode instanciates
func NewEmptyNode(key string) *Node {
	return NewNode(key, &map[string]interface{}{})
}

// SetProperty to a node
func (node *Node) SetProperty(key string, value interface{}) {
	(*node.Props)[key] = value
}

// AddOutEdge
func (node *Node) AddOutEdge(edge *Edge) {
	(*node.Out)[edge.Key] = edge
}

// AddInEdge
func (node *Node) AddInEdge(edge *Edge) {
	(*node.In)[edge.Key] = edge
}

// Get a property
func (node *Node) Get(key string) (interface{}, error) {
	value, ok := (*node.Props)[key]
	if !ok {
		return nil, errorNodePropNotFound(node.Key, key)
	}

	return value, nil
}

// OutE returns all the outgoing edges with given label
func (node *Node) OutE(label string) []*Edge {

	result := []*Edge{}

	for _, edge := range *node.Out {
		if edge.Label == label {
			result = append(result, edge)
		}
	}

	return result
}

// InE returns all the ingoing edges with given label
func (node *Node) InE(label string) []*Edge {

	result := []*Edge{}

	for _, edge := range *node.In {
		if edge.Label == label {
			result = append(result, edge)
		}
	}

	return result
}

// Out returns all the outgoing nodes with given label
func (node *Node) OutV(label string) []*Node {

	result := []*Node{}

	for _, edge := range *node.Out {
		if edge.Label == label {
			result = append(result, edge.EndNode)
		}
	}

	return result
}

// In returns all the ingoing nodes with given label
func (node *Node) InV(label string) []*Node {

	result := []*Node{}

	for _, edge := range *node.In {
		if edge.Label == label {
			result = append(result, edge.StartNode)
		}
	}

	return result
}
