package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hippoai/graphgo"
)

func main() {

	g := build()
	query := graphgo.NewQuery(g, "company.ups")

	exists := func(q *graphgo.Query) bool {
		return q.Size() > 0
	}
	log.Println(exists)

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

	b, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println("ret", string(b))

}
