package ast

type Expr interface {
	Eval() (any, string)
	Type() string
	Out(any) error
}
