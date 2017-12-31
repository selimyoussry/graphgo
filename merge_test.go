package graphgo

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/hippoai/goerr"
	"github.com/hippoai/goutil"
)

func loadData(fileName string) (*Output, error) {

	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var o Output
	err = json.Unmarshal(b, &o)
	if err != nil {
		return nil, err
	}

	return &o, nil

}

func TestMerge(t *testing.T) {

	output1, err := loadData("./data/data1.json")
	if err != nil {
		t.Errorf(goutil.Pretty(err))
	}

	// Create a graph and merge this data into it
	g := NewEmptyGraph()
	g.Merge(output1)

	var nodeKey string

	// should find node from legacy index
	nodeKey, err = g.FindNodeFromLegacyIndex("5")
	if err != nil {
		t.Errorf(goutil.Pretty(err))
	}
	if nodeKey != "person.jimbo" {
		t.Errorf(goutil.Pretty(goerr.NewS("CANT_FIND_NODE")))
	}

	// should find node prop
	vItf, err := g.GetNodeProp("person.elliott", "age")
	if err != nil {
		t.Errorf(goutil.Pretty(err))
	}
	v, ok := vItf.(float64)
	if !ok {
		t.Errorf(goutil.Pretty(goerr.NewS("WRONG_FORMAT")))
	}
	if v != 80 {
		t.Errorf(goutil.Pretty(goerr.New("WRONG_VALUE", map[string]interface{}{
			"age": v,
		})))
	}

	// should find edge from legacy index
	edgeKey, err := g.FindEdgeFromLegacyIndex("9")
	if err != nil {
		t.Errorf(goutil.Pretty(err))
	}
	if edgeKey != "elliott.is_father_of.john" {
		t.Errorf(goutil.Pretty(goerr.NewS("CANT_FIND_EDGE")))
	}

	// Delete nodes and edges
	output2, err := loadData("./data/data2.json")
	if err != nil {
		t.Errorf(goutil.Pretty(err))
	}

	// Create a graph and merge this data into it
	g.Merge(output2)

	// Should not have a node person.patrick anymore
	if g.HasNode("person.patrick") {
		t.Errorf(goutil.Pretty(goerr.NewS("STILL_HAS_NODE")))
	}

	// should not have an edge patrick.worksin.OnePlus anymore
	if g.HasEdge("patrick.worksin.OnePlus") {
		t.Errorf(goutil.Pretty(goerr.NewS("STILL_HAS_EDGE")))
	}

	// should not have a node person.tim anymore
	if g.HasNode("person.tim") {
		t.Errorf(goutil.Pretty(goerr.NewS("STILL_HAS_NODE")))
	}

	// should not have an edge tim.is_father_of.jimbo anymore
	if g.HasEdge("tim.is_father_of.jimbo") {
		t.Errorf(goutil.Pretty(goerr.NewS("STILL_HAS_EDGE")))
	}

	// Should still have a node person.patrick
	if !g.HasNode("company.OnePlus") {
		t.Errorf(goutil.Pretty(goerr.NewS("LOST_NODE")))
	}

	// Should have a node person.luisa
	if !g.HasNode("person.luisa") {
		t.Errorf(goutil.Pretty(goerr.NewS("LOST_NODE")))
	}

	// should have an edge clara.worksin.OnePlus
	if !g.HasEdge("clara.worksin.OnePlus") {
		t.Errorf(goutil.Pretty(goerr.NewS("STILL_HAS_EDGE")))
	}

	// should still have an edge luisa.worksin.OnePlus
	if !g.HasEdge("luisa.worksin.OnePlus") {
		t.Errorf(goutil.Pretty(goerr.NewS("STILL_HAS_EDGE")))
	}

}
