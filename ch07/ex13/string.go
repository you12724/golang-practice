package eval

import (
	"fmt"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%g", l)
}

func (u unary) String() string {
	return fmt.Sprintf("%c %s", u.op, u.x.String())
}

func (b binary) String() string {
	return fmt.Sprintf("%s %c %s", b.x.String(), b.op, b.y.String())
}

func (c call) String() string {
	switch c.fn {
	case "pow":
		return fmt.Sprintf("%s(%s, %s)", c.fn, c.args[0], c.args[1])
	default:
		return fmt.Sprintf("%s(%s)", c.fn, c.args[0])
	}
}
