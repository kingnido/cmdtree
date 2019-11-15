package cmdtree

import (
	"errors"
	"fmt"
)

type P struct {
	Name string
	Next N
}

func (p P) next(c *Context, s string) (N, error) {
	c.Params = append(c.Params, fmt.Sprintf("%s: %s", p.Name, s))
	return p.Next, nil
}

func (p P) run(c *Context) error {
	return errors.New("no runner for this type")
}
