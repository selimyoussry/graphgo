package graphgo

import "testing"

func TestAddNode(t *testing.T) {

	g := NewEmptyGraph()

	nodeProps := map[string]interface{}{
		"prop1": 120,
		"prop2": "hello",
	}
	g.MergeNode("node.1", nodeProps)

	node, err := g.getNode("node.1")
	if err != nil {
		t.Errorf(err.Error())
	}

	for k, v := range node.Props {
		if v != nodeProps[k] {
			t.Errorf("Did not match %s", k)
		}
	}

}
