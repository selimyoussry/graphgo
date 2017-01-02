package graphgo

import "github.com/hippoai/goerr"

// errNodeNotFound when a node is not found in the graph, given its key
func errNodeNotFound(key string) *goerr.Err {
	return goerr.New(ERR_NODE_NOT_FOUND, map[string]interface{}{
		"key": key,
	})
}

// errEdgeNotFound when an edge is not found in the graph, given its key
func errEdgeNotFound(key string) *goerr.Err {
	return goerr.New(ERR_EDGE_NOT_FOUND, map[string]interface{}{
		"key": key,
	})
}

// errorNodePropNotFound when a property is not found on a node
func errorNodePropNotFound(nodeKey, prop string) *goerr.Err {
	return goerr.New(ERR_NODE_PROP_NOT_FOUND, map[string]interface{}{
		"nodeKey": nodeKey,
		"prop":    prop,
	})
}

// errorEdgePropNotFound when a property is not found on a node
func errorEdgePropNotFound(edgeKey, prop string) *goerr.Err {
	return goerr.New(ERR_EDGE_PROP_NOT_FOUND, map[string]interface{}{
		"edgeKey": edgeKey,
		"prop":    prop,
	})
}
