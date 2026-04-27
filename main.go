package main

import (
	"fmt"

	"github.com/fobus89/fufik/internal/parser"
)

func main() {
	p := parser.NewParser("1+2")

	// p.LedRegister(kind token.TokenType, bp parser.BindingPower, ledHander parser.LedHandlerType)
	// p.NudRegister(kind token.TokenType, bp parser.BindingPower, ledHander parser.LedHandlerType)
	// p.StmtRegister(kind token.TokenType, bp parser.BindingPower, ledHander parser.LedHandlerType)

	fmt.Println(p.Parse())
}

func reg(p *parser.Parser) {
}

func ledBinary(p *Parser, left ast.Expr, bp BindingPower) (ast.Expr, error) {
	opToken := p.next()

	right, err := p.parseExpr(bp)
	{
		if err != nil {
			return nil, err
		}
	}

	return ast.NewBinaryExpr(left, opToken.Token, right), nil
}
