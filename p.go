package cmdtree

import (
	"errors"
)

func (p P) next(c *Context, s string) (N, error) {
	c.Params = append(c.Params, KV{p.Name, s})
	return p.Next, nil
}

func (p P) run(c *Context) error {
	return errors.New("no runner for this type")
}
