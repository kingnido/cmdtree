package cmdtree

func (r R) next(c *Context, s string) (N, error) {
	return r.Next.next(c, s)
}

func (r R) run(c *Context) error {
	return r.Func(c)
}
