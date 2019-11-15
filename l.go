package cmdtree

import "errors"

func (l L) next(c *Context, s string) (N, error) {
	return nil, errors.New("too many arguments")
}

func (l L) run(c *Context) error {
	return l.Func(c)
}
