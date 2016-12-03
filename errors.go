package graphgo

import (
	"errors"
	"fmt"
)

func errNodeNotFound(key string) error {
	return errors.New(fmt.Sprintf("Could not find node %s", key))
}
