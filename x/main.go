package main

import (
	"fmt"

	"github.com/hippoai/goerr"
)

type NodeI interface {
	Get() error
}

type X interface {
	Get() error
}

type Node struct {
}

func (node *Node) Get() *goerr.Err {
	return goerr.New("my err", map[string]interface{}{})
}

func get(e error) string {
	return e.Error()
}

func main() {

	e := goerr.New("my e", map[string]interface{}{})
	fmt.Println(get(e))

}
