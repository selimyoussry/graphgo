package graphgo

import (
	"errors"
	"fmt"
)

func errNodeNotFound(key string) error {
	return errors.New(fmt.Sprintf("Could not find node %s", key))
}

func errEdgeNotFound(key string) error {
	return errors.New(fmt.Sprintf("Could not find edge %s", key))
}

func errorNodePropNotFound(nodeKey, key string) error {
	return errors.New(fmt.Sprintf("Could not find node (%s)'s' property %s", nodeKey, key))
}

func errorEdgePropNotFound(nodeKey, key string) error {
	return errors.New(fmt.Sprintf("Could not find node (%s)'s' property %s", nodeKey, key))
}
