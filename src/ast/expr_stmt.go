package ast

type ExprStmt struct {
	Expr Expr
}

func (e ExprStmt) Stmt() Expr {
	return e.Expr
}

func NewExprStmt(expr Expr) *ExprStmt {
	return &ExprStmt{Expr: expr}
}
