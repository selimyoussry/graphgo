// Package graphgo implements standard interfaces for libraries working on directed property graphs (see interfaces.go)
// It also provides a simple implementation for a directed property graph, compatible with these interfaces
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

// getNode in local format
func (graph *Graph) getNode(key string) (*Node, error) {
	node, ok := graph.Nodes[key]
	if !ok {
		return nil, errNodeNotFound(key)
	}
	return node, nil
}

// GetNode finds a node given its key
func (graph *Graph) GetNode(key string) (INode, error) {
	return graph.getNode(key)
}

// GetNodeProp finds a node prop
func (graph *Graph) GetNodeProp(key, prop string) (interface{}, error) {
	node, err := graph.getNode(key)
	if err != nil {
		return "", err
	}
	return node.Get(prop)
}

// MergeNode adds a node to the graph if it does not exist, or merges its properties ottherwise
func (graph *Graph) MergeNode(key string, props map[string]interface{}) (*Node, error) {

	node, err := graph.getNode(key)

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
	for k, v := range props {
		node.SetProperty(k, v)
	}

	return node, nil
}

// getEdge in local format
func (graph *Graph) getEdge(key string) (*Edge, error) {
	edge, ok := graph.Edges[key]
	if !ok {
		return nil, errEdgeNotFound(key)
	}
	return edge, nil
}

// GetEdge gets an existing edge or returns an error
func (graph *Graph) GetEdge(key string) (IEdge, error) {
	return graph.getEdge(key)
}

// GetEdgeProp finds a node prop
func (graph *Graph) GetEdgeProp(key, prop string) (interface{}, error) {
	edge, err := graph.getEdge(key)
	if err != nil {
		return "", err
	}
	return edge.Get(prop)
}

// MergeEdge adds an edge to the graph if it does not exist, merges its properties otherwise
func (graph *Graph) MergeEdge(edgeKey, label string, start, end string, props map[string]interface{}) (*Edge, error) {
	edge, err := graph.getEdge(edgeKey)

	// If the edge does not exist
	if err != nil {
		edge = NewEdge(edgeKey, label, start, end, props)

		startNode, err := graph.getNode(start)
		if err != nil {
			return nil, errNodeNotFound(start)
		}

		endNode, err := graph.getNode(end)
		if err != nil {
			return nil, errNodeNotFound(end)
		}

		graph.Edges[edgeKey] = edge
		startNode.AddOutEdge(edgeKey, label)
		endNode.AddInEdge(edgeKey, label)

		return edge, nil
	}

	if props == nil {
		return edge, nil
	}

	// Otherwise modify existing edge
	for k, v := range props {
		edge.SetProperty(k, v)
	}

	return edge, nil
}
