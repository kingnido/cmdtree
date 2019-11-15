package cmdtree

import (
	"errors"
	"strings"
)

type N interface {
	next(*Context, string) (N, error)
	run(*Context) error
}

func Exec(n N, s string) error {
	var err error
	c := &Context{}

	for _, s = range strings.Split(s, " ") {
		if s == "" {
			continue
		}

		if n == nil {
			return errors.New("too many args")
		}

		n, err = n.next(c, s)
		if err != nil {
			return err
		}
	}

	return n.run(c)
}
