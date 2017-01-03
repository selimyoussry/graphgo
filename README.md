# Graphgo

Implements standard interfaces for libraries working on directed property graph. It also comes with a simple implementation of directed property graph, compatible with these interfaces. It is compatible with the traversal library [AskGo](https://github.com/hippoai/askgo.git).

## Concepts

The graph is a collection of nodes and edges. Each node and edge is identified by a unique key, which helps us retrieve them in the graph. The graph essentially maintains a nodes map `map[string]*Node` and an edges map `map[string]*Edge`, using the node and edge keys as map keys.

The standard interfaces mentioned above are the following:

```go
// Graph needs to be able to find a node and an edge given their key
type IGraph interface {
	GetNode(key string) (INode, error)
	GetEdge(key string) (IEdge, error)
}

// Edge needs to be able to access its properties, start and end node, and label
// it has a unique key
type IEdge interface {
	Get(key string) (interface{}, error)
	Hop(graph IGraph, key string) (INode, error)
	StartN(graph IGraph) (INode, error)
	EndN(graph IGraph) (INode, error)
	GetLabel() string
	GetKey() string
}

// Node needs to be able to access its properties,
// ingoing and outgoing edges
// it has a unique key
type INode interface {
	Get(key string) (interface{}, error)
	InE(graph IGraph, label string) (map[string]IEdge, error)
	OutE(graph IGraph, label string) (map[string]IEdge, error)
	GetKey() string
}
```

## Install

`go get -u github.com/hippoai/graphgo.git`

## Getting started

Below, we list a few code snippets and use cases for this package. Make sure you import this package `import "github.com/hippoai/graphgo"`.

### Create a graph and add a bunch of nodes and edges

```go
// Instanciate an empty graph
g := graphgo.NewEmptyGraph()

// Add a few nodes, identified by their unique key
g.MergeNode("company.ups", map[string]interface{}{
  "name":     "ups",
  "location": "america",
})

g.MergeNode("person.patrick", map[string]interface{}{
  "name": "patrick",
  "age":  20,
})

g.MergeNode("person.john", map[string]interface{}{
  "name": "john",
  "age":  55,
})

// Now, add the edges (relationships) between these nodes
g.MergeEdge(
  "patrick.worksin.ups", "WORKS_IN",
  "person.patrick", "company.ups",
  map[string]interface{}{},
)
g.MergeEdge(
  "john.worksin.ups", "WORKS_IN",
  "person.john", "company.ups",
  map[string]interface{}{},
)
g.MergeEdge(
  "patrick.is_son_of.john", "IS_SON_OF",
  "person.patrick", "person.john",
  map[string]interface{}{},
)
```

### Once you've got your graph, traverse it
with [AskGo](https://github.com/hippoai/askgo.git)
