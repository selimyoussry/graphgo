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

	exists := func(q *graphgo.Query, path []*graphgo.Step) bool {
		return q.Size() > 0
	}

	hasSameCompanyName := func(q *graphgo.Query, path []*graphgo.Step) bool {
		if q.Size() == 0 {
			return false
		}

		company := path[0]
		companyName, err := company.Node.Get("name")
		if err != nil {
			return false
		}
		for _, sonCompany := range q.Result() {
			sonCompanyName, err := sonCompany.Get("name")
			if err != nil {
				continue
			}
			log.Println("Company name", companyName.(string), sonCompanyName.(string))
			if sonCompanyName.(string) != sonCompanyName.(string) {
				return true
			}
		}

		return false

	}

	result := query.
		Log("0. Start").
		In("WORKS_IN", true). // Find the employees
		Log("1. Employees").
		Deepen(). // Deep query
		Log("1.55 Deepened").
		In("IS_SON_OF", false). // For each employee, get the sons
		Log("1.6 Sons before filtering").
		DeepFilter(exists). // filter sons
		Log("1.7 Filter guys who have sons").
		Deepen(). // Deepen filter sons working at UPS
		Log("1.8 second depth").
		Out("WORKS_IN", false). // get companies they work for
		Log("1.9 second depth > companies").
		DeepFilter(hasSameCompanyName). // keep only sons who work for a company
		Log("1.95 second filter").
		Flatten().
		Log("1.96 Re-flatten").
		Save("name::sonName"). // son name
		DeepSave("sons").
		Log("2. Sons").
		Flatten().
		Log("3. Flattened").
		Save("name::fatherName", "age").
		Log("1.5 Father name").
		Return()

	b, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println("ret", string(b))

}
