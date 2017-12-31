package graphgo

import "github.com/hippoai/goerr"

// errNodeNotFound when a node is not found in the graph, given its key
func errNodeNotFound(key string) error {
	return goerr.New(ERR_NODE_NOT_FOUND, map[string]interface{}{
		"key": key,
	})
}

// errEdgeNotFound when an edge is not found in the graph, given its key
func errEdgeNotFound(key string) error {
	return goerr.New(ERR_EDGE_NOT_FOUND, map[string]interface{}{
		"key": key,
	})
}

// errorNodePropNotFound when a property is not found on a node
func errorNodePropNotFound(nodeKey, prop string) error {
	return goerr.New(ERR_NODE_PROP_NOT_FOUND, map[string]interface{}{
		"nodeKey": nodeKey,
		"prop":    prop,
	})
}

// errorEdgePropNotFound when a property is not found on a node
func errorEdgePropNotFound(edgeKey, prop string) error {
	return goerr.New(ERR_EDGE_PROP_NOT_FOUND, map[string]interface{}{
		"edgeKey": edgeKey,
		"prop":    prop,
	})
}

func errConnectedNode(nodeKey string) error {
	return goerr.New(ERR_CONNECTED_NODE, map[string]interface{}{
		"nodeKey": nodeKey,
	})
}

func errNoNodeLegacyRecord(nodeKey string) error {
	return goerr.New(ERR_NO_NODE_LEGACY_RECORD, map[string]interface{}{
		"nodeKey": nodeKey,
	})
}

func errNoEdgeLegacyRecord(edgeKey string) error {
	return goerr.New(ERR_NO_EDGE_LEGACY_RECORD, map[string]interface{}{
		"edgeKey": edgeKey,
	})
}
