# Graphgo

Implements a **directed property graph**, along with a [Gremlin](https://tinkerpop.apache.org) like query language, in [Go](https://golang.org).

## Concepts

The graph is a collection of nodes and edges. Each node and edge is identified by a unique key, which helps us retrieve them in the graph. The graph essentially maintains a nodes map `map[string]*Node` and an edges map `map[string]*Edge`, using the node and edge keys as map keys.

## Install

`go get -u github.com/hippoai/graphgo.git`

## Getting started

Below, we list a few code snippets and use cases for this package. Make sure you import this package `import "github.com/hippoai/graphgo"`.

### 1. Create a graph

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

### 2 - Query the graph

Graphs are great to store relational data, because they provide an efficient data structure to explore these relationships. Asking questions to `graphgo`'s graphs is done using the `Query` structure, that wraps around the graph. The result will be stored in a nested object that will be built during the query (check the following examples to sense what's going on).

Let's create a new query, starting at the node "company.ups".
```go
query := graphgo.NewQuery(g, "company.ups")
```

Now, let's answer a few questions.

#### What are the names and ages of the people working for UPS?
```go
result := query.
  In("WORKS_IN").
  Save("name", "age").
  Return()
```
 will return
 ```javascript
{
  "person.john": {
    "name": "john",
    "age": 55
  },
  "person.patrick": {
    "name": "patrick",
    "age": 20
  }
}
 ```

#### What are the names and age of the UPS employees who have a son? Oh, and give me the son's name too.

Steps:
* Get all UPS employees
* Create a "Deep query" on this result, which will allow us to filter based on the employees properties and relationships (find the ones with a son).
* Save the son's name
* Return the fathers' names and ages.

```go
// Determine whether you have a son
exists := func (q * graphgo.Query) bool{
  return q.Size() > 0
}

result := query.
  In("WORKS_IN"). // Find the employees
  Deepen(). // Deep query
  In("IS_SON_OF"). // For each employee, get the sons
  Deepen(). // For each son, check if he works at UPS
  Out("WORKS_IN"). // Get the companies they work for
  DeepFilter(exists). // only keep the sons working for a company
  Flatten(). // return to the son level
  DeepFilter(exists). // Filter the ones with no son
  Save("name"). // Save the son's name in the deep cache
  DeepSave("sons"). // Save in the cache, under "sons" key
  Flatten(). // Return at the lower level
  Save("name", "age"). // Save the father's age
  Return()
```

will return

```javascript
{
  "person.john": {
    "age": 55,
    "name": "john",
    "sons": {
      "person.patrick": {
        "name": "patrick"
      }
    }
  }
}
```

#### Let us add another son and run the same query.

For this, we will first add a few nodes and relationships to the graph.

```go
g.MergeNode("person.tim", map[string]interface{}{
  "name": "Tim",
  "age":  28,
})

g.MergeEdge(
  "tim.is_son_of.john", "IS_SON_OF",
  "person.tim", "person.john",
  map[string]interface{}{},
)
```

we get

```javascript
{
  "person.john": {
    "age": 55,
    "name": "john",
    "sons": {
      "person.patrick": {
        "name": "patrick"
      },
      "person.tim": {
        "name": "Tim"
      }
    }
  }
}
```

#### Now we only want the sons who also work at UPS, so we need to add another filter somewhere.

```golang
result := query.
  Log("0. Start").
  In("WORKS_IN"). // Find the employees
  Log("1. Employees").
  Deepen(). // Deep query
  Log("1.55 Deepened").
  In("IS_SON_OF"). // For each employee, get the sons
  Log("1.6 Sons before filtering").
  DeepFilter(exists). // filter sons
  Log("1.7 Filter guys who have sons").
  Deepen(). // Deepen filter sons working at UPS
  Log("1.8 second depth").
  Out("WORKS_IN"). // get companies they work for
  Log("1.9 second depth > companies").
  DeepFilter(exists). // keep only sons who work for a company
  Log("1.95 second filter").
  Flatten().
  Save("name::sonName"). // son name
  DeepSave("sons").
  Log("2. Sons").
  Flatten().
  Log("3. Flattened").
  Save("name::fatherName").
  Log("1.5 Father name").
  Return()
```
