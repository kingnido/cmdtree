package cmdtree

import (
	"errors"
	"fmt"
	"strings"
)

type Context struct {
	Params []string
}

type N interface {
	next(*Context, string) (N, error)
	run(*Context) error
}

type M map[string]N

type P struct {
	Name string
	Next N
}

type T struct {
	Func func(*Context) error
	Next N
}

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

func (p P) next(c *Context, s string) (N, error) {
	c.Params = append(c.Params, fmt.Sprintf("%s: %s", p.Name, s))
	return p.Next, nil
}

func (p P) run(c *Context) error {
	return errors.New("no runner for this type")
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

func Exec(n N, s string) {
	var err error
	c := &Context{}

	for _, s = range strings.Split(s, " ") {
		if s == "" {
			continue
		}

		if n == nil {
			fmt.Println("too many args")
			return
		}

		n, err = n.next(c, s)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println(n.run(c))
}
