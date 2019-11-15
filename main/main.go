package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/kingnido/cmdtree"
)

func dumpContext(c *Context) error {
	fmt.Println(c.Params)
	return nil
}

func giveAll(c *Context) error {
	fmt.Printf("giving %s for all\n", c.Params[0].V)
	return nil
}

func giveOne(c *Context) error {
	fmt.Printf("giving %s for %s\n", c.Params[0].V, c.Params[1].V)
	return nil
}

type Person struct {
	Name string
}

func (p *Person) giveFunc() func(*Context) error {
	return func(c *Context) error {
		target := "everyone"

		if len(c.Params) == 2 {
			target = c.Params[1].V
		}

		fmt.Printf("%s gave %s to %s\n", p.Name, c.Params[0].V, target)
		return nil
	}
}

func main() {
	var cmdtree N

	p := Person{"lokinho"}

	cmdtree = M{
		"give": P{
			"item", R{p.giveFunc(), M{
				"to": P{
					"player", L{p.giveFunc()}}}}},
		"turn": M{
			"left":  L{dumpContext},
			"right": L{dumpContext}},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for fmt.Print("> "); scanner.Scan(); fmt.Print("> ") {
		Exec(cmdtree, scanner.Text())
	}
}
