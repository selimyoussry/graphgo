package graphgo

// DeepQuery treats every node as its own Query
// Stores the queries in Queries
// Attributes a Key to the result of each query
type DeepQuery struct {
	Queries       map[string]*Query
	Key           string
	InitialResult map[string]*Node
}

// Out
func (dq *DeepQuery) Out(label string) *DeepQuery {
	for _, query := range dq.Queries {
		query = query.Out(label)
	}
	return dq
}

// In
func (dq *DeepQuery) In(label string) *DeepQuery {
	for _, query := range dq.Queries {
		query = query.In(label)
	}
	return dq
}

// FilterNodes
func (dq *DeepQuery) FilterNodes(predicate func(map[string]interface{}) bool) *DeepQuery {
	for _, query := range dq.Queries {
		query = query.FilterNodes(predicate)
	}
	return dq
}

// Get
func (dq *DeepQuery) Get(name string, keys ...string) *DeepQuery {
	for _, query := range dq.Queries {
		query = query.Get(name, keys...)
	}
	return dq
}

// GetOne
func (dq *DeepQuery) GetOne(name string, keys ...string) *DeepQuery {
	for _, query := range dq.Queries {
		query = query.GetOne(name, keys...)
	}
	return dq
}

// Flatten
func (dq *DeepQuery) Flatten() *Query {
	q := NewEmptyQuery()
	q.result = dq.InitialResult

	for nodeKey, query := range dq.Queries {
		q.Graph = query.Graph
		q.Cache[nodeKey] = query.Cache
	}

	return q
}
