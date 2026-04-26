package ast

type BlockStmt struct {
	Body []Stmt
}

func (b *BlockStmt) Stmt() Expr {
	return nil
}

func NewBlockStmt(body []Stmt) *BlockStmt {
	return &BlockStmt{Body: body}
}
