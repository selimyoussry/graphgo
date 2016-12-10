package graphgo

// Graph stores nodes and edges in maps, using their unique keys
type Graph struct {
	Nodes map[string]*Node `json:"nodes"`
	Edges map[string]*Edge `json:"edges"`
}

// NewEmptyGraph instanciates
func NewEmptyGraph() *Graph {
	return &Graph{
		Nodes: map[string]*Node{},
		Edges: map[string]*Edge{},
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
func (graph *Graph) MergeNode(key string, props *map[string]interface{}) (*Node, error) {

	node, err := graph.GetNode(key)

	// If the node does not exist
	if err != nil {
		node = NewNode(key, props)
		graph.Nodes[key] = node

		return node, nil
	}

	if props == nil {
		return node, nil
	}

	// Otherwise, the node does not exist yet, merge the properties
	for k, v := range *props {
		node.SetProperty(k, v)
	}

	return node, nil
}

// GetEdge gets an existing edge or returns an error
func (graph *Graph) GetEdge(key string) (*Edge, error) {
	edge, ok := graph.Edges[key]
	if !ok {
		return nil, errEdgeNotFound(key)
	}
	return edge, nil
}

// MergeEdge adds an edge to the graph if it does not exist, merges its properties otherwise
func (graph *Graph) MergeEdge(key, label string, start, end string, props *map[string]interface{}) (*Edge, error) {
	edge, err := graph.GetEdge(key)

	// If the edge does not exist
	if err != nil {
		edge = NewEdge(key, label, start, end, props)
		graph.Edges[key] = edge

		startNode, err := graph.GetNode(start)
		if err != nil {
			startNode.AddOutEdge(edge.Key, label)
		}

		endNode, err := graph.GetNode(end)
		if err != nil {
			endNode.AddInEdge(edge.Key, label)
		}

		return edge, nil
	}

	if props == nil {
		return edge, nil
	}

	// Otherwise modify existing edge
	for k, v := range *props {
		edge.SetProperty(k, v)
	}

	return edge, nil
}
