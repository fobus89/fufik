package parser

import (
	"fmt"

	"github.com/fobus89/fufik/src/ast"
)

func (p *parser) parseStmt() (ast.Stmt, error) {
	if stmtHandler, ok := p.stmtLookup.GetOrNone(p.CurrentTokenKind()); ok {
		return stmtHandler(p)
	}

	expr, err := p.parseExpr(Lowest)
	{
		if err != nil {
			return nil, err
		}
	}

	return ast.NewExprStmt(expr), nil
}

func (p *parser) parseExpr(bp BindingPower) (ast.Expr, error) {
	tokKind := p.CurrentTokenKind()

	nudHandler, ok := p.nudLookup.GetOrNone(tokKind)
	{
		if !ok {
			panic(fmt.Sprintf("%s not found", tokKind))
		}
	}

	left, err := nudHandler(p)
	{
		if err != nil {
			return nil, err
		}
	}

	for {

		tokKind = p.CurrentTokenKind()

		curBp := p.bpLookup.Get(tokKind)
		{
			if curBp <= bp {
				break
			}
		}

		ledHandler, ok := p.ledLookup.GetOrNone(tokKind)
		{
			if !ok {
				panic(fmt.Sprintf("%s not found", tokKind))
			}
		}

		left, err = ledHandler(p, left, curBp)
		{
			if err != nil {
				return nil, err
			}
		}
	}

	return left, nil
}
