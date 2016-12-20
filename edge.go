package graphgo

// Edge has a unique key, properties, start and end node
type Edge struct {
	Key   string                 `json:"key"`
	Label string                 `json:"label"`
	Props map[string]interface{} `json:"props"`
	Start string                 `json:"start"`
	End   string                 `json:"end"`
}

// NewEdge instanciates
func NewEdge(key, label string, start, end string, props map[string]interface{}) *Edge {
	return &Edge{
		Key:   key,
		Label: label,
		Props: props,
		Start: start,
		End:   end,
	}
}

// SetProperty one by one
func (edge *Edge) SetProperty(key string, value interface{}) {
	edge.Props[key] = value
}

// Get a property
func (edge *Edge) Get(key string) (interface{}, error) {
	value, ok := edge.Props[key]
	if !ok {
		return nil, errorEdgePropNotFound(edge.Key, key)
	}

	return value, nil
}
