package parse_literal

import (
	"strconv"

	"github.com/fobus89/fufik"
)

type Statement struct{}

func RegisterParser(p fufik.Parser) {
	p.NudRegister(fufik.INT_LITERAL, nudIntLiteral)
	p.NudRegister(fufik.STRING_LITERAL, nudStringLiteral)
	p.NudRegister(fufik.FLOAT_LITERAL, nudFloat64Literal)
	p.NudRegister(fufik.IDENT, nudIdentLiteral)
}

func nudIntLiteral(p fufik.Parser) (fufik.Expr, error) {
	literal := p.Next()

	numb, err := strconv.Atoi(literal.Literal)
	{
		if err != nil {
			return nil, err
		}
	}

	return NewIntExpr(numb), nil
}

func nudFloat64Literal(p fufik.Parser) (fufik.Expr, error) {
	literal := p.Next()

	numb, err := strconv.ParseFloat(literal.Literal, 64)
	{
		if err != nil {
			return nil, err
		}
	}

	return NewFloat64Expr(numb), nil
}

func nudStringLiteral(p fufik.Parser) (fufik.Expr, error) {
	literal := p.Next()
	return NewStringExpr(literal.Literal), nil
}

func nudIdentLiteral(p fufik.Parser) (fufik.Expr, error) {
	literal := p.Next()
	return NewStringExpr(literal.Literal), nil
}
