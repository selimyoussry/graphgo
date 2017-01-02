package graphgo

// Error implements our custom errors
type Err struct {
	code  string
	Props map[string]interface{}
}

// Error to implement the error interface
func (e *Err) Error() string {
	return e.code
}

// NewErr instanciates
func NewErr(code string, props map[string]interface{}) *Err {
	return &Err{
		code:  code,
		Props: props,
	}
}

// errNodeNotFound when a node is not found in the graph, given its key
func errNodeNotFound(key string) *Err {
	return NewErr(ERR_NODE_NOT_FOUND, map[string]interface{}{
		"key": key,
	})
}

// errEdgeNotFound when an edge is not found in the graph, given its key
func errEdgeNotFound(key string) *Err {
	return NewErr(ERR_EDGE_NOT_FOUND, map[string]interface{}{
		"key": key,
	})
}

// errorNodePropNotFound when a property is not found on a node
func errorNodePropNotFound(nodeKey, prop string) *Err {
	return NewErr(ERR_NODE_PROP_NOT_FOUND, map[string]interface{}{
		"nodeKey": nodeKey,
		"prop":    prop,
	})
}

// errorEdgePropNotFound when a property is not found on a node
func errorEdgePropNotFound(edgeKey, prop string) *Err {
	return NewErr(ERR_EDGE_PROP_NOT_FOUND, map[string]interface{}{
		"edgeKey": edgeKey,
		"prop":    prop,
	})
}
