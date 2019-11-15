package cmdtree

import "errors"

type T struct {
	Func func(*Context) error
	Next N
}

func (t T) next(c *Context, s string) (N, error) {
	if t.Next != nil {
		return t.Next.next(c, s)
	}

	return nil, errors.New("too many arguments")
}

func (t T) run(c *Context) error {
	return t.Func(c)
}
