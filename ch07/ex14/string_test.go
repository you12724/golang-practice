package eval

import "testing"

func TestStringVar(t *testing.T) {
	for _, test := range []struct {
		input  Var
		result string
	}{
		{"1", "1"},
		{"", ""},
		{"100", "100"},
	} {
		if test.input.String() != test.result {
			t.Errorf("output is %s, expected %s", test.input.String(), test.result)
		}
	}
}

func TestStringLiteral(t *testing.T) {
	for _, test := range []struct {
		input  literal
		result string
	}{
		{1, "1"},
		{0, "0"},
		{10.000001, "10.000001"},
	} {
		if test.input.String() != test.result {
			t.Errorf("output is %s, expected %s", test.input.String(), test.result)
		}
	}
}

func TestStringUnary(t *testing.T) {
	for _, test := range []struct {
		input  unary
		result string
	}{
		{unary{'+', Var("1")}, "+ 1"},
		{unary{'-', Var("100")}, "- 100"},
		{unary{'+', Var("1.1134")}, "+ 1.1134"},
	} {
		if test.input.String() != test.result {
			t.Errorf("output is %s, expected %s", test.input.String(), test.result)
		}
	}
}

func TestStringBinary(t *testing.T) {
	for _, test := range []struct {
		input  binary
		result string
	}{
		{binary{'/', Var("100"), Var("10")}, "100 / 10"},
		{binary{'-', Var("100"), Var("1")}, "100 - 1"},
		{binary{'+', Var("1.1134"), Var("0.3")}, "1.1134 + 0.3"},
		{binary{'*', Var("1"), Var("0.3")}, "1 * 0.3"},
	} {
		if test.input.String() != test.result {
			t.Errorf("output is %s, expected %s", test.input.String(), test.result)
		}
	}
}

func TestStringCall(t *testing.T) {
	for _, test := range []struct {
		input  call
		result string
	}{
		{call{"pow", []Expr{Var("2"), Var("2")}}, "pow(2, 2)"},
		{call{"sin", []Expr{Var("1.22")}}, "sin(1.22)"},
		{call{"sqrt", []Expr{Var("0.33")}}, "sqrt(0.33)"},
	} {
		if test.input.String() != test.result {
			t.Errorf("output is %s, expected %s", test.input.String(), test.result)
		}
	}
}

func TestParse(t *testing.T) {
	for _, test := range []struct {
		input Expr
	}{
		{call{"pow", []Expr{Var("2"), Var("2")}}},
		{call{"sin", []Expr{Var("1.22")}}},
		{call{"sqrt", []Expr{Var("0.33")}}},
		{call2{"max", []Expr{Var("0.33"), Var("10")}}},
		{binary{'-', Var("100"), Var("1")}},
		{unary{'+', Var("1")}},
	} {
		expr, _ := Parse(test.input.String())
		if expr.String() != test.input.String() {
			t.Errorf("cant parse, result: %s -> %s", test.input.String(), expr.String())
		}
	}
}
