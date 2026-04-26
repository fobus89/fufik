package ast

type Stmt interface {
	Stmt() Expr
}

type Expr interface {
	Expr() Expr
}
