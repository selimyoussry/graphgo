package graphgo

// Query on top of a Graph instance
// Aims to have "functional" style
type Query struct {
	Graph  *Graph
	result map[string]*Node
	Cache  map[string]interface{}
}

// NewEmptyQuery instanciates
func NewEmptyQuery() *Query {
	return &Query{
		Graph:  nil,
		result: map[string]*Node{},
		Cache:  map[string]interface{}{},
	}
}

// NewQuery instanciates
func NewQuery(g *Graph, starts ...string) *Query {
	result := map[string]*Node{}

	for _, start := range starts {
		node, err := g.GetNode(start)
		if err != nil {
			continue
		}
		// At this point. we found the start node, and add it to the result graph
		result[node.Key] = node
	}

	return &Query{
		Graph:  g,
		result: result,
		Cache:  map[string]interface{}{},
	}

}

// Out returns outgoing nodes to this graph
func (q *Query) Out(label string) *Query {
	newResult := map[string]*Node{}

	// Loop over all the nodes in the current result
	for _, node := range q.result {

		// Loop over all relationships for this node
		for edgeKey, edgeLabel := range node.Out {

			// Only keep the ones with given label
			if edgeLabel == label {

				edge, err := q.Graph.GetEdge(edgeKey)
				if err != nil {
					continue
				}

				endNode, err := q.Graph.GetNode(edge.End)
				if err != nil {
					continue
				}

				newResult[endNode.Key] = endNode

			}

		}

	}

	q.result = newResult

	return q
}

// In returns outgoing nodes to this graph
func (q *Query) In(label string) *Query {
	newResult := map[string]*Node{}

	// Loop over all the nodes in the current result
	for _, node := range q.result {

		// Loop over all relationships for this node
		for edgeKey, edgeLabel := range node.In {

			// Only keep the ones with given label
			if edgeLabel == label {

				edge, err := q.Graph.GetEdge(edgeKey)
				if err != nil {
					continue
				}

				startNode, err := q.Graph.GetNode(edge.Start)
				if err != nil {
					continue
				}

				newResult[startNode.Key] = startNode

			}

		}

	}

	q.result = newResult

	return q
}

// FilterNodes based on a predicate on their properties
func (q *Query) FilterNodes(predicate func(map[string]interface{}) bool) *Query {

	newResult := map[string]*Node{}

	// Loop over all the nodes in the current result
	for nodeKey, node := range q.result {

		if predicate(node.Props) {
			newResult[nodeKey] = node
		}

	}

	q.result = newResult
	return q

}

// Flatten function
// Get an iterable of all the keys, per node
func (q *Query) Get(name string, keys ...string) *Query {
	out := map[string](map[string]interface{}){}

	// Loop over every node in the result
	for nodeKey, node := range q.result {
		m := map[string]interface{}{}

		// Loop over every key we care about
		for _, key := range keys {
			value, err := node.Get(key)
			if err != nil {
				continue
			}
			m[key] = value
		}

		out[nodeKey] = m

	}

	q.Cache[name] = out
	return q
}

// GetOne returns a map of the keys and their values
// works ONLY if there is only one node in the result
func (q *Query) GetOne(name string, keys ...string) *Query {
	// If the result holds more than one node, throw error
	if len(q.result) != 1 {
		return q
	}

	// Extract the unique node
	var node *Node
	for _, _node := range q.result {
		node = _node
	}
	out := map[string]interface{}{}

	// Loop over every node in the result
	// Loop over every key we care about
	for _, key := range keys {
		value, err := node.Get(key)
		if err != nil {
			continue
		}
		out[key] = value
	}

	q.Cache[name] = out

	return q
}

// // Output just returns the result
// func (q *Query) Output() map[string]*Node {
// 	copy := map[string]*Node{}
//
// 	for nodeKey, node := range q.result {
// 		copy[nodeKey] = node.Copy()
// 	}
//
// 	return copy
// }

// DeepenQuery creates a new DeepQuery, from every node of a given Query
func (query *Query) Deepen(key string) *DeepQuery {

	b := map[string]*Query{}
	for _, r := range query.result {

		// Use the node key as a query key
		b[r.Key] = NewQuery(query.Graph, r.Key)

	}

	return &DeepQuery{
		Queries:       b,
		Key:           key,
		InitialResult: query.result,
	}

}

// Flatten returns the cache
func (q *Query) Flatten() map[string]interface{} {
	return q.Cache
}
