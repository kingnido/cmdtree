package cmdtree

import (
	"fmt"
	"io"
	"strings"
)

type Handler interface {
	Handle(*Context, io.Writer, ...string)
}

type HandlerFunc func(*Context, io.Writer, ...string)

func (f HandlerFunc) Handle(c *Context, w io.Writer, args ...string) {
	f(c, w, args...)
}

type M map[string]Handler

func (m M) Handle(c *Context, w io.Writer, args ...string) {
	if len(args) == 0 {
		fmt.Fprintf(w, "missing arguments\n")
		return
	}

	next, ok := m[args[0]]
	if !ok {
		fmt.Fprintf(w, "invalid arguments\n")
		return
	}

	next.Handle(c, w, args[1:]...)
}

type P struct {
	Key  string
	Next Handler
}

func (p P) Handle(c *Context, w io.Writer, args ...string) {
	if len(args) == 0 {
		fmt.Fprintf(w, "missing arguments\n")
		return
	}

	c.KeyValues = append(c.KeyValues, KeyValue{p.Key, args[0]})

	p.Next.Handle(c, w, args[1:]...)
}

type T struct {
	This Handler
	Next Handler
}

func (t T) Handle(c *Context, w io.Writer, args ...string) {
	if len(args) == 0 {
		t.This.Handle(c, w, args...)
		return
	}

	t.Next.Handle(c, w, args...)
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

func Exec(h Handler, w io.Writer, s string) {
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

	h.Handle(c, w, ss...)
}
