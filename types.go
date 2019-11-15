package cmdtree

type M map[string]N

type P struct {
	Name string
	Next N
}

type L struct {
	Func func(*Context) error
}

type Context struct {
	Params []KV
}

type KV struct {
	K string
	V string
}
