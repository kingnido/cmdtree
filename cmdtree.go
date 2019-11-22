package cmdtree

import (
	"errors"
	"strings"
)

type Handler interface {
	Handle(*Context, ...string) error
}

type HandlerFunc func(*Context, ...string) error

func (f HandlerFunc) Handle(c *Context, args ...string) error {
	return f(c, args...)
}

type M map[string]Handler

func (m M) Handle(c *Context, args ...string) error {
	if len(args) == 0 {
		return errors.New("missing arguments")
	}

	next, ok := m[args[0]]
	if !ok {
		return errors.New("invalid arguments")
	}

	return next.Handle(c, args[1:]...)
}

type P struct {
	Key  string
	Next Handler
}

func (p P) Handle(c *Context, args ...string) error {
	if len(args) == 0 {
		return errors.New("missing arguments")
	}

	c.KeyValues = append(c.KeyValues, KeyValue{p.Key, args[0]})

	return p.Next.Handle(c, args[1:]...)
}

type T struct {
	This Handler
	Next Handler
}

func (t T) Handle(c *Context, args ...string) error {
	if len(args) == 0 {
		return t.This.Handle(c, args...)
	}

	return t.Next.Handle(c, args...)
}

type KeyValue struct {
	Key   string
	Value string
}

type Context struct {
	Command   string
	Parts     []string
	KeyValues []KeyValue
}

func Exec(h Handler, s string) error {
	var ss []string
	for _, s := range strings.Split(s, " ") {
		if s != "" {
			ss = append(ss, s)
		}
	}

	c := &Context{
		Command:   s,
		Parts:     ss,
		KeyValues: []KeyValue{},
	}

	return h.Handle(c, ss...)
}
