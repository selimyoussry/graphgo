package graphgo

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Query on top of a Graph instance
// Aims to have "functional" style
type Query struct {
	Graph  *Graph
	result map[string]*Node
	Cache  map[string](map[string]interface{})
	Path   map[string][]*Step

	Queries map[string]*Query
	IsDeep  bool
}

// Step, how did we get there
type Step struct {
	Node *Node
	Edge *Edge
}

// NewStep instanciates
func NewStep(node *Node, edge *Edge) *Step {
	return &Step{
		Node: node,
		Edge: edge,
	}
}

// RenameKey returns the wanted key
// if key == originalKey::newKey, return old key, new key
func RenameKey(key string) (string, string) {
	splitted := strings.Split(key, KEY_SEPARATOR)
	if len(splitted) == 1 {
		return splitted[0], splitted[0]
	}

	return splitted[0], splitted[1]
}

// NewQuery
func NewQuery(g *Graph, starts ...string) *Query {
	path := map[string][]*Step{}
	for _, start := range starts {
		path[start] = []*Step{}
	}

	return NewQueryWithPath(g, path, starts...)
}

// NewQuery instanciates
func NewQueryWithPath(g *Graph, path map[string][]*Step, starts ...string) *Query {
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
		Cache:  map[string](map[string]interface{}){},
		Path:   path,

		Queries: map[string]*Query{},
		IsDeep:  false,
	}

}

// IsDoubleDeep returns true if depth >= 2
func (q *Query) IsDoubleDeep() bool {
	if !q.IsDeep {
		return false
	}

	for _, nestedQuery := range q.Queries {
		return nestedQuery.IsDeep
	}

	return false
}

// Out returns outgoing nodes to this graph
func (q *Query) Out(label string, rememberPath bool) *Query {

	// Deep Calls
	if q.IsDeep {
		for _, nestedQuery := range q.Queries {
			nestedQuery.Out(label, rememberPath)
		}
		return q
	}

	newResult := map[string]*Node{}
	newPath := map[string][]*Step{}

	// Loop over all the nodes in the current result
	for nodeKey, node := range q.result {

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
				if rememberPath {
					newPath[endNode.Key] = append(q.Path[nodeKey], NewStep(node, edge))
				} else {
					newPath[endNode.Key] = q.Path[nodeKey]
				}

			}

		}

	}

	q.result = newResult
	q.Path = newPath

	return q
}

// In returns outgoing nodes to this graph
func (q *Query) In(label string, rememberPath bool) *Query {

	// Deep Calls
	if q.IsDeep {
		for _, nestedQuery := range q.Queries {
			nestedQuery.In(label, rememberPath)
		}
		return q
	}

	newResult := map[string]*Node{}
	newPath := map[string][]*Step{}

	// Loop over all the nodes in the current result
	for nodeKey, node := range q.result {

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
				if rememberPath {
					newPath[startNode.Key] = append(q.Path[nodeKey], NewStep(node, edge))
				} else {
					newPath[startNode.Key] = q.Path[nodeKey]
				}

			}

		}

	}

	q.result = newResult
	q.Path = newPath

	return q
}

// FilterNodes based on a predicate on their properties
func (q *Query) FilterNodes(predicate func(map[string]interface{}) bool) *Query {

	// Deep Calls
	if q.IsDeep {
		for _, nestedQuery := range q.Queries {
			nestedQuery.FilterNodes(predicate)
		}
		return q
	}

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
// If key is formatted "originalKey::newKey", we rename the key
func (q *Query) Save(keys ...string) *Query {

	// Deep Calls
	if q.IsDeep {
		for _, nestedQuery := range q.Queries {
			nestedQuery.Save(keys...)
		}
		return q
	}

	// Loop over every node in the result
	for nodeKey, node := range q.result {
		_, exists := q.Cache[nodeKey]
		if !exists {
			q.Cache[nodeKey] = map[string]interface{}{}
		}

		// Loop over every key we care about
		for _, key := range keys {
			oldKey, newKey := RenameKey(key)
			value, err := node.Get(oldKey)
			if err != nil {
				continue
			}
			q.Cache[nodeKey][newKey] = value
		}
	}
	return q
}

// DeepenQuery creates a new DeepQuery, from every node of a given Query
func (q *Query) Deepen() *Query {
	// Deep Calls
	if q.IsDeep {
		for _, nestedQuery := range q.Queries {
			nestedQuery.Deepen()
		}
		return q
	}

	queries := map[string]*Query{}
	for _, r := range q.result {

		// Use the node key as a query key
		queries[r.Key] = NewQueryWithPath(q.Graph, q.Path, r.Key)

	}

	q.Queries = queries
	q.IsDeep = true
	return q
}

// DeepSave flattens a query to the lower level
func (q *Query) DeepSave(name string) *Query {

	// Nothing to flatten
	if !q.IsDeep {
		return q
	}

	// If it's actually too deep, we keep going
	if q.IsDoubleDeep() {
		for _, nestedQuery := range q.Queries {
			nestedQuery.DeepSave(name)
		}
		return q
	}

	// Otherwise, this is the level before the lowest
	// We can flatten the cache
	for nodeKey, nestedQuery := range q.Queries {
		_, exists := q.Cache[nodeKey]
		if !exists {
			q.Cache[nodeKey] = map[string]interface{}{}
		}

		q.Cache[nodeKey][name] = nestedQuery.Cache
	}

	return q

}

// Flatten flattens a query to the lower level
func (q *Query) Flatten() *Query {

	// Nothing to flatten
	if !q.IsDeep {
		return q
	}

	// If it's actually too deep, we keep going
	if q.IsDoubleDeep() {
		for _, nestedQuery := range q.Queries {
			nestedQuery.Flatten()
		}
		return q
	}

	q.Queries = map[string]*Query{}
	q.IsDeep = false
	return q

}

// DeepFilter
func (q *Query) DeepFilter(keepQuery func(*Query, []*Step) bool) *Query {

	// Nothing to flatten
	if !q.IsDeep {
		return q
	}

	// If it's actually too deep, we keep going
	if q.IsDoubleDeep() {
		for _, nestedQuery := range q.Queries {
			nestedQuery.DeepFilter(keepQuery)
		}
		return q
	}

	// Otherwise, this is the level before the lowest
	// We can flatten the cache

	nodesToDiscard := []string{}
	for nodeKey, nestedQuery := range q.Queries {

		// if we need to filter this
		if !keepQuery(nestedQuery, q.Path[nodeKey]) {
			nodesToDiscard = append(nodesToDiscard, nodeKey)
		}

	}

	// Delete the nodes that have been filtered
	for _, nodeKey := range nodesToDiscard {
		delete(q.result, nodeKey)
		delete(q.Queries, nodeKey)
	}

	return q
}

// Return the cache value
func (q *Query) Return() map[string](map[string]interface{}) {
	return q.Cache
}

// Size returns how many nodes were found
func (q *Query) Size() int {
	return len(q.result)
}

func (q *Query) Result() map[string]*Node {
	return q.result
}

func (q *Query) Log(msgs ...string) *Query {
	for _, msg := range msgs {
		fmt.Println(msg)
	}

	fmt.Println("> Cache")
	q.LogCache()

	fmt.Println("> Result")
	q.LogResult()

	fmt.Println("> Path")
	q.LogPath()

	fmt.Println("--- --- ---")

	return q
}

func (q *Query) LogCache() *Query {

	b, _ := json.MarshalIndent(q.Cache, "", "  ")
	fmt.Println(string(b))

	return q
}

func (q *Query) LogResult() *Query {

	out := q.deepLog()

	b, _ := json.MarshalIndent(out, "", "  ")
	fmt.Println(string(b))
	return q

}

// deepLog is called on a deep the first time
func (q *Query) deepLog() map[string]interface{} {

	if !q.IsDeep {
		result := map[string]interface{}{}
		for nodeKey, _ := range q.result {
			result[nodeKey] = ""
		}
		return result
	}

	// Otherwise modify tmp
	result := map[string]interface{}{}
	for nodeKey, nestedQuery := range q.Queries {
		oneResult := nestedQuery.deepLog()
		result[nodeKey] = oneResult
	}
	return result

}

// LogPath -
func (q *Query) LogPath() *Query {

	for nodeKey, onePath := range q.Path {
		steps := []string{}
		for _, step := range onePath {
			steps = append(steps, fmt.Sprintf("%s (%s)", step.Node.Key, step.Edge.Key))
		}

		fmt.Printf("%s: %s \n", nodeKey, strings.Join(steps, " > "))
	}

	return q
}
