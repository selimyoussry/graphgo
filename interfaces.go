package graphgo

// Graph needs to be able to find a node and an edge given their key
type IGraph interface {
	GetNode(key string) (INode, error)
	GetEdge(key string) (IEdge, error)
}

// Edge needs to be able to access its properties, start and end node, and label
// it has a unique key
type IEdge interface {
	Get(key string) (interface{}, error)
	Hop(graph IGraph, key string) (INode, error)
	StartN(graph IGraph) (INode, error)
	EndN(graph IGraph) (INode, error)
	GetLabel() string
	GetKey() string
	GetProps() map[string]interface{}
}

// Node needs to be able to access its properties,
// ingoing and outgoing edges
// it has a unique key
type INode interface {
	Get(key string) (interface{}, error)
	InE(graph IGraph, label string) (map[string]IEdge, error)
	OutE(graph IGraph, label string) (map[string]IEdge, error)
	GetKey() string
	GetProps() map[string]interface{}
}
