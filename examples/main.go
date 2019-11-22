package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/kingnido/cmdtree"
)

func dumpContextHandler() Handler {
	return HandlerFunc(func(c *Context, args ...string) error {
		fmt.Fprintln(os.Stdout, c, args)
		return nil
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
	return HandlerFunc(func(c *Context, args ...string) error {
		p.Name = c.KeyValues[0].Value
		return nil
	})
}

func getNameHandler(p *Player) Handler {
	return HandlerFunc(func(c *Context, args ...string) error {
		fmt.Fprintln(os.Stdout, p.Name)
		return nil
	})
}

func moveHandler(p *Player, dx int, dy int) Handler {
	return HandlerFunc(func(c *Context, args ...string) error {
		p.Move(dx, dy)
		return nil
	})
}

func whereHandler(p *Player) Handler {
	return HandlerFunc(func(c *Context, args ...string) error {
		fmt.Fprintf(os.Stdout, "x: %d, y: %d\n", p.X, p.Y)
		return nil
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
		if err := Exec(cmdtree, scanner.Text()); err != nil {
			fmt.Fprintf(os.Stdin, "error: %v\n", err)
		}
	}
	fmt.Printf("\n")
}
