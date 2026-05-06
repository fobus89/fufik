package parse_arithmetic

import (
	"fmt"

	"github.com/fobus89/fufik"
	"github.com/fobus89/fufik/internal/parser"
	"github.com/fobus89/fufik/internal/token"
)

type Statement struct{}

func RegisterParser(p fufik.Parser) {

	p.NudRegister(fufik.LPARENT, nudGrouping)

	p.LedRegister(fufik.PLUS, fufik.Additive, ledBinary)
	p.LedRegister(fufik.MINUS, fufik.Additive, ledBinary)

	p.LedRegister(fufik.STAR, fufik.Muptiplicative, ledBinary)
	p.LedRegister(fufik.SLASH, fufik.Muptiplicative, ledBinary)
	p.LedRegister(fufik.PERCENT, fufik.Muptiplicative, ledBinary)
}

func nudGrouping(p fufik.Parser) (fufik.Expr, error) {
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

func ledBinary(p fufik.Parser, left fufik.Expr, bp fufik.BindingPower) (fufik.Expr, error) {
	if !p.MatchAny(fufik.PLUS, fufik.MINUS, fufik.STAR, fufik.SLASH, fufik.PERCENT) {
		return nil, fmt.Errorf("expected PLUS, MINUS, STAR, SLASH, PERCEN got %v", p.CurrentToken())
	}

	opToken := p.Next()

	right, err := p.ParseExpr(bp)
	{
		if err != nil {
			return nil, err
		}
	}

	return NewBinaryExpr(opToken.Type, left, right), nil
}
