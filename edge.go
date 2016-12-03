package graphgo

// Edge has a unique key, properties, start and end node
type Edge struct {
	Key   string                 `json:"key"`
	Label string                 `json:"label"`
	Props map[string]interface{} `json:"props"`
	Start string                 `json:"start"`
	End   string                 `json:"end"`
}
