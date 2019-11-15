package cmdtree

import (
	"errors"
	"fmt"
)

func (m M) next(c *Context, s string) (N, error) {
	n, ok := m[s]
	if !ok {
		return nil, errors.New(fmt.Sprintf("%s not found", s))
	}

	return n, nil
}

func (m M) run(c *Context) error {
	return errors.New("no runner for this type")
}
