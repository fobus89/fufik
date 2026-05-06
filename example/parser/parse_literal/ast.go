package parse_literal

import "github.com/fobus89/fufik"

// Eval() (any, string)
// Type() string
// Out(any) error

type Float64 float64

var _ fufik.Expr = (Float64)(0)

func NewFloat64Expr(value float64) Float64 {
	return Float64(value)
}

func (f Float64) Eval() (any, string) {
	return f, "float64"
}

func (f Float64) Out(any) error {
	panic("unimplemented")
}

func (f Float64) Type() string {
	return "float64"
}

var _ fufik.Expr = (Int)(0)

type Int int

func NewIntExpr(value int) Int {
	return Int(value)
}

func (i Int) Eval() (any, string) {
	return i, "int"
}

func (i Int) Out(any) error {
	panic("unimplemented")
}

func (i Int) Type() string {
	return "int"
}

var _ fufik.Expr = (String)(0)

type String string

func NewStringExpr(value string) String {
	return String(value)
}

func (s String) Eval() (any, string) {
	return s, "string"
}

func (s String) Out(any) error {
	panic("unimplemented")
}

func (s String) Type() string {
	return "string"
}

type NumberExpr struct {
	Value float64
}

func NewNumberExpr(value float64) *NumberExpr {
	return &NumberExpr{
		Value: value,
	}
}

func (b *NumberExpr) Eval() any {
	return b.Value
}
