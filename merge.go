package graphgo

// Merge will merge a serialized graph into the current graph
func (graph *Graph) Merge(o *Output) error {

	// Remove edges
	for _, edgeKey := range o.Delete.Edges {
		graph.DeleteEdge(edgeKey)
	}

	// Remove edges from legacy index
	for _, legacyEdgeKey := range o.Delete.LegacyEdges {
		graph.DeleteEdgeFromLegacyIndex(legacyEdgeKey)
	}

	// Remove nodes
	for _, nodeKey := range o.Delete.Nodes {
		graph.DeleteNode(nodeKey)
	}

	// Remove nodes from legacy index
	for _, legacyNodeKey := range o.Delete.LegacyNodes {
		graph.DeleteNodeFromLegacyIndex(legacyNodeKey)
	}

	// Add new nodes
	for nodeKey, node := range o.Merge.Nodes {
		graph.MergeNode(nodeKey, node.Props)
	}

	// Add new edges
	for edgeKey, edge := range o.Merge.Edges {
		graph.MergeEdge(edgeKey, edge.Label, edge.Start, edge.End, edge.Props)
	}

	// Add nodes legacy index
	for legacyNodeIndex, nodeKey := range o.Merge.LegacyIndex.Nodes {
		graph.AddNodeLegacyIndex(legacyNodeIndex, nodeKey)
	}

	// Add edges legacy index
	for legacyEdgeIndex, edgeKey := range o.Merge.LegacyIndex.Edges {
		graph.AddEdgeLegacyIndex(legacyEdgeIndex, edgeKey)
	}
	return nil

}
