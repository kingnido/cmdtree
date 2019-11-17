# cmdtree

A simple library to parse a command string and map to a handler. Kind of based
on net/http's handlers approach.

## Example

```go
func moveHandler(direction string) Handler {
    return HandlerFunc(func(c *cmdtree.Context, w io.Writer, args ...string) {
        fmt.Fprintf(w, "going %s\n", direction)
    })
}

func setNameHandler(name *string) Handler {
    return HandlerFunc(func(c *cmdtree.Context, w io.Writer, args ...string) {
        // c.KeyValues[0].Key == "newname"
        *name = c.KeyValues[0].Value
    })
}

func getNameHandler(name *string) Handler {
    return HandlerFunc(func(c *cmdtree.Context, w io.Writer, args ...string) {
        fmt.Fprintf(w, "name: %s\n", *name)
    })
}
...
// 
name := "alice"

tree := cmdtree.M{
    "move": cmdtree.M{
        "north": moveHandler("north"),
        "south": moveHandler("south"),
        "west":  moveHandler("west"),
        "east":  moveHandler("east"),
    },
    "name": cmdtree.T{getNameHandler(&name), cmdtree.P{
        "newname", setNameHandler(&name)}},
}

scanner := bufio.NewScanner(os.Stdin)
for fmt.Print("> "); scanner.Scan(); fmt.Print("> ") {
    cmdtree.Exec(tree, os.Stdout, scanner.Text())
}
fmt.Printf("\n")
// possible commands:
// move north
// move south
// name
// name bob
//
// invalid commands
// move up
// move north east
// asd
// name bob smith
```

M stands for Map. The argument at this point must be a key in this map.

T stands for Terminal. If there's no argument left to be consumed, the left
handler will be selected

P stands for Parameter. The argument at this point will be save to the
context's slice with the key defined in the first field.
