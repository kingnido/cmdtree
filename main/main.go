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

func main() {
	var cmdtree N

	cmdtree = M{
		"give": P{
			"item", M{
				"to": P{
					"player", L{dumpContext}}}},
		"turn": M{
			"left":  L{dumpContext},
			"right": L{dumpContext}},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for fmt.Print("> "); scanner.Scan(); fmt.Print("> ") {
		Exec(cmdtree, scanner.Text())
	}
}
