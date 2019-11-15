package cmdtree

type N interface {
	next(*Context, string) (N, error)
	run(*Context) error
}

type M map[string]N

type P struct {
	Name string
	Next N
}

type L struct {
	Func func(*Context) error
}

type R struct {
	Func func(*Context) error
	Next N
}

type Context struct {
	Params []KV
}

type KV struct {
	K string
	V string
}
