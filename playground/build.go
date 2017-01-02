package main

import "github.com/hippoai/graphgo"

func build() *graphgo.Graph {

	g := graphgo.NewEmptyGraph()

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

	// New son
	g.MergeNode("person.tim", map[string]interface{}{
		"name": "Tim",
		"age":  28,
	})

	g.MergeEdge(
		"tim.is_son_of.john", "IS_SON_OF",
		"person.tim", "person.john",
		map[string]interface{}{},
	)

	return g
}
