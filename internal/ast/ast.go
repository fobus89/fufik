package ast

type Expr interface {
	Eval() any
}
