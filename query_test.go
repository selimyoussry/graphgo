package graphgo

import (
	"encoding/json"
	"testing"
)

func TestQuery(t *testing.T) {

	g := NewEmptyGraph()

	// Add nodes to the graph
	g.MergeNode("company.ups", map[string]interface{}{
		"name":     "ups",
		"location": "america",
	})

	g.MergeNode("employee.patrick", map[string]interface{}{
		"name": "patrick",
		"age":  20,
	})

	g.MergeNode("employee.john", map[string]interface{}{
		"name": "john",
		"age":  "30",
	})

	g.MergeNode("celebrity.travolta", map[string]interface{}{
		"name":    "john travolta",
		"twitter": "@travolta",
		"age":     60,
	})

	g.MergeNode("celebrity.obama", map[string]interface{}{
		"name":    "barack obama",
		"twitter": "@potus",
		"age":     55,
	})

	g.MergeNode("celebrity.bocelli", map[string]interface{}{
		"name":    "andrea bocelli",
		"twitter": "@bocelli",
		"age":     40,
	})

	// Add edges

	// Who works in the company
	g.MergeEdge("john.worksin.ups", "WORKS_IN", "employee.john", "company.ups", map[string]interface{}{})
	g.MergeEdge("patrick.worksin.ups", "WORKS_IN", "employee.patrick", "company.ups", map[string]interface{}{})

	// Who follows who
	g.MergeEdge("john.follows.travolta", "FOLLOWS", "employee.john", "celebrity.travolta", map[string]interface{}{})
	g.MergeEdge("john.follows.obama", "FOLLOWS", "employee.john", "celebrity.obama", map[string]interface{}{})

	g.MergeEdge("patrick.follows.bocelli", "FOLLOWS", "employee.patrick", "celebrity.bocelli", map[string]interface{}{})

	// Query
	// We want to know all the followed celebrities in the company
	query := NewQuery(g, "company.ups")
	result := query.
		In("WORKS_IN").
		Out("FOLLOWS").
		Get("followed_celebrities", "name", "twitter").
		Flatten()

	t.Log(PrettyPrint(result))

	// Now get all the company's followed celebrities
	result = NewQuery(g, "company.ups").
		Deepen("employees_following_bocelli").
		In("WORKS_IN").Out("FOLLOWS").Get("celebrity", "name").
		Flatten().
		Flatten()

	t.Log(PrettyPrint(result))

	// For every company, get all the employees following Bocelli
	// TO DO

	// Print graph
	// gBytes, _ := json.MarshalIndent(g, "", "  ")
	// t.Logf(string(gBytes))
}

func PrettyPrint(v interface{}) string {
	b, _ := json.MarshalIndent(v, "", "    ")
	return string(b)
}
