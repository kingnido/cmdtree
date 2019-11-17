package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	. "github.com/kingnido/cmdtree"
)

func dumpContextHandler() Handler {
	return HandlerFunc(func(c *Context, w io.Writer, args ...string) {
		fmt.Fprintln(w, c, args)
	})
}

type Player struct {
	Name string
	X    int
	Y    int
}

func (p *Player) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

func setNameHandler(p *Player) Handler {
	return HandlerFunc(func(c *Context, w io.Writer, args ...string) {
		p.Name = c.KeyValues[0].Value
	})
}

func getNameHandler(p *Player) Handler {
	return HandlerFunc(func(c *Context, w io.Writer, args ...string) {
		fmt.Fprintln(w, p.Name)
	})
}

func moveHandler(p *Player, dx int, dy int) Handler {
	return HandlerFunc(func(c *Context, w io.Writer, args ...string) {
		p.Move(dx, dy)
	})
}

func whereHandler(p *Player) Handler {
	return HandlerFunc(func(c *Context, w io.Writer, args ...string) {
		fmt.Fprintf(w, "x: %d, y: %d\n", p.X, p.Y)
	})
}

func main() {
	p := &Player{"alice", 0, 0}

	cmdtree := M{
		"move": M{
			"north": moveHandler(p, 0, -1),
			"south": moveHandler(p, 0, 1),
			"west":  moveHandler(p, -1, 0),
			"east":  moveHandler(p, 1, 0),
		},
		"whereami": whereHandler(p),
		"name": T{getNameHandler(p), P{
			"name", setNameHandler(p)}},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for fmt.Print("> "); scanner.Scan(); fmt.Print("> ") {
		Exec(cmdtree, os.Stdout, scanner.Text())
	}
	fmt.Printf("\n")
}
