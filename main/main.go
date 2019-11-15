package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func dumpContext(c *Context) error {
	fmt.Println(c.ss)
	return nil
}

func main() {
	var cmdtree N

	cmdtree = M{
		"give": P{"item", M{"to": P{"player", T{dumpContext, nil}}}},
		"turn": M{
			"left":  T{dumpContext, nil},
			"right": T{dumpContext, nil}},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for fmt.Print("> "); scanner.Scan(); fmt.Print("> ") {
		Exec(cmdtree, scanner.Text())
	}

}
