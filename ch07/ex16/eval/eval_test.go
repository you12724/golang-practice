package eval

import "testing"

func TestCall2Eval(t *testing.T) {
	for _, test := range []struct {
		input  call2
		result float64
	}{
		{call2{"max", []Expr{literal(0), literal(-1)}}, 0},
		{call2{"max", []Expr{literal(-2), literal(-1)}}, -1},
		{call2{"max", []Expr{literal(1000), literal(10)}}, 1000},
		{call2{"max", []Expr{literal(0.33)}}, 0.33},
		{call2{"min", []Expr{literal(0), literal(-1)}}, -1},
		{call2{"min", []Expr{literal(2), literal(1)}}, 1},
		{call2{"min", []Expr{literal(1000), literal(10)}}, 10},
		{call2{"min", []Expr{literal(0.33)}}, 0.33},
	} {
		output := test.input.Eval(Env{})
		if output != test.result {
			t.Errorf("output is %f, expected %f", output, test.result)
		}
	}
}
