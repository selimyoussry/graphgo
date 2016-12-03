package graphgo

// Graph stores nodes and edges in maps, using their unique keys
type Graph struct {
	Nodes         map[string]*Node `json:"nodes"`
	DirectEdges   map[string]*Edge `json:"directEdges"`
	IndirectEdges map[string]*Edge `json:"indirectEdges"`
}

// NewEmptyGraph instanciates
func NewEmptyGraph() *Graph {
	return &Graph{
		Nodes:         map[string]*Node{},
		DirectEdges:   map[string]*Edge{},
		IndirectEdges: map[string]*Edge{},
	}
}

// GetNode finds a node given its key
func (graph *Graph) GetNode(key string) (*Node, error) {
	node, ok := graph.Nodes[key]
	if !ok {
		return nil, errNodeNotFound(key)
	}
	return node, nil
}

// MergeNode adds a node to the graph if it does not exist, or merges its properties ottherwise
func (graph *Graph) MergeNode(key string, props *map[string]interface{}) error {
	node, err := graph.GetNode(key)

	// If the node does not exist
	if err != nil {
		graph.Nodes[key] = NewNode(key, props)
		return nil
	}

	// Otherwise, the node does not exist yet, merge the properties
	for k, v := range *props {
		node.SetProperty(k, v)
	}

	return nil
}
