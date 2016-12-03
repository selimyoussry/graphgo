package graphgo

// Node implements a graph node
type Node struct {
	Key   string                  `json:"key"`
	Props *map[string]interface{} `json:"props"`
}

// NewNode instanciates
func NewNode(key string, props *map[string]interface{}) *Node {
	return &Node{
		Key:   key,
		Props: props,
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
