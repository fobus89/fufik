package parse_statement

import (
	"fmt"
	"strconv"

	"github.com/fobus89/fufik"
	"github.com/fobus89/fufik/internal/parser"
	"github.com/fobus89/fufik/internal/token"
)

type Statement struct{}

func RegisterParser(p fufik.Parser) {
	p.NudRegister(fufik.INT_LITERAL, Statement{}.nudIntLiteral)
	p.NudRegister(fufik.FLOAT_LITERAL, Statement{}.nudIntLiteral)
	p.NudRegister(fufik.LPARENT, Statement{}.nudGrouping)
}

func (Statement) nudGrouping(p fufik.Parser) (fufik.Expr, error) {
	if !p.MatchNext(token.LPARENT) {
		return nil, fmt.Errorf("expected LPARENT, got %v", p.CurrentToken())
	}

	expr, err := p.ParseExpr(parser.Lowest)
	{
		if err != nil {
			return nil, err
		}
	}

	if !p.MatchNext(fufik.RPARENT) {
		return nil, fmt.Errorf("expected RPARENT, got %v", p.CurrentToken())
	}

	return expr, nil
}

func (Statement) nudIntLiteral(p fufik.Parser) (fufik.Expr, error) {
	literal := p.Next()

	numb, err := strconv.ParseFloat(literal.Literal, 64)
	{
		if err != nil {
			return nil, err
		}
	}

	return NewNumberExpr(numb), nil
}
