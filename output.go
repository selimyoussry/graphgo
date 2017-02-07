package graphgo

type Delete struct {
	Nodes       []string `json:"nodes"`
	Edges       []string `json:"edges"`
	LegacyNodes []string `json:"legacyNodes"`
	LegacyEdges []string `json:"legacyEdges"`
}

func NewEmptyDelete() *Delete {
	return &Delete{
		Nodes:       []string{},
		Edges:       []string{},
		LegacyNodes: []string{},
		LegacyEdges: []string{},
	}
}

// Delete then Merge into new graph to augment it
type Output struct {
	Merge  *Graph  `json:"$merge"`
	Delete *Delete `json:"$delete"`
}

func NewOutput() *Output {
	return &Output{
		Merge:  NewEmptyGraph(),
		Delete: NewEmptyDelete(),
	}
}
