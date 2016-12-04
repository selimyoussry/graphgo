package graphgo

import (
	"encoding/json"
	"testing"
)

func TestEdge(t *testing.T) {

	g := NewEmptyGraph()

	nodeProps := &map[string]interface{}{
		"prop1": 120,
		"prop2": "hello",
	}
	g.MergeNode("node.1", nodeProps)

	node2Props := &map[string]interface{}{
		"prop1": 120,
		"prop2": "hello",
	}
	g.MergeNode("node.2", node2Props)

	node1, err := g.GetNode("node.1")
	if err != nil {
		t.Errorf("Could not find node1")
	}

	node2, err := g.GetNode("node.2")
	if err != nil {
		t.Errorf("Could not find node2")
	}

	edgeProps := &map[string]interface{}{
		"prop3": true,
	}
	g.MergeEdge("a.b", "is_friend", node1, node2, edgeProps)

	// Now modify node1's property
	node1.SetProperty("prop3", 199)
	node1.SetProperty("prop2", "world")

	// Find the edge
	edge, err := g.GetEdge("a.b")
	if err != nil {
		t.Errorf("Could not find edge")
	}

	if edge.Start != node1.Key {
		t.Errorf("Wrong start node")
	}
	if edge.End != node2.Key {
		t.Errorf("Wrong end node")
	}
	prop2, err := edge.StartNode.Get("prop2")
	if err != nil {
		t.Errorf("Could not get prop2")
	}
	if prop2.(string) != "world" {
		t.Errorf("Update of node prop %s is not propagated", "prop2")
	}
	prop3, err := edge.StartNode.Get("prop3")
	if err != nil {
		t.Errorf("Could not get prop3")
	}
	if prop3.(int) != 199 {
		t.Errorf("Update of node prop %s is not propagated", "prop3")
	}

	out, err := json.Marshal(g)
	if err != nil {
		t.Errorf("Could not marshal the graph")
	}
	t.Logf(string(out))

}
