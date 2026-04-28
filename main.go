package main

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/fobus89/fufik/internal/ast"
	"github.com/fobus89/fufik/internal/parser"
	"github.com/fobus89/fufik/internal/token"
)

type (
	Parser       = parser.Parser
	Expr         = ast.Expr
	BindingPower = parser.BindingPower
)

type NumberExpr struct {
	Value float64
}

func NewNumberExpr(value float64) *NumberExpr {
	return &NumberExpr{
		Value: value,
	}
}

func (b *NumberExpr) Eval() any {
	return b.Value
}

type BinaryExpr struct {
	Left  Expr
	Op    token.TokenType
	Right Expr
}

func NewBinaryExpr(op token.TokenType, left, right Expr) *BinaryExpr {
	return &BinaryExpr{
		Left:  left,
		Op:    op,
		Right: right,
	}
}

func (b *BinaryExpr) Eval() any {
	leftVal := b.Left.Eval().(float64)
	rightVal := b.Right.Eval().(float64)

	switch b.Op {
	case token.PLUS:
		return leftVal + rightVal

	case token.MINUS:
		return leftVal - rightVal

	case token.STAR:
		return leftVal * rightVal

	case token.SLASH:
		return leftVal / rightVal

	case token.PERCENT:
		return math.Mod(leftVal, rightVal)
	}

	return 0
}

func main() {
	p := parser.NewParser("( 1+2 )*( 3*2 )")

	register(p)

	exprs, err := p.Parse()
	{
		if err != nil {
			log.Fatalln(err)
		}
	}

	for _, expr := range exprs {
		fmt.Println(expr.Eval())
	}
}

func register(p parser.Parser) {
	p.LedRegister(token.PLUS, parser.Additive, ledBinary)
	p.LedRegister(token.MINUS, parser.Additive, ledBinary)

	p.LedRegister(token.STAR, parser.Muptiplicative, ledBinary)
	p.LedRegister(token.SLASH, parser.Muptiplicative, ledBinary)
	p.LedRegister(token.PERCENT, parser.Muptiplicative, ledBinary)

	p.NudRegister(token.INT_LITERAL, nudIntLiteral)
	p.NudRegister(token.FLOAT_LITERAL, nudIntLiteral)

	p.NudRegister(token.LPARENT, nudGrouping)
}

func nudGrouping(p Parser) (ast.Expr, error) {
	if !p.MatchNext(token.LPARENT) {
		return nil, fmt.Errorf("expected LPARENT, got %v", p.CurrentToken())
	}

	expr, err := p.ParseExpr(parser.Lowest)
	{
		if err != nil {
			return nil, err
		}
	}

	if !p.MatchNext(token.RPARENT) {
		return nil, fmt.Errorf("expected LPARENT, got %v", p.CurrentToken())
	}

	return expr, nil
}

func nudIntLiteral(p Parser) (ast.Expr, error) {
	numb, err := strconv.ParseFloat(p.Next().Literal, 64)
	{
		if err != nil {
			return nil, err
		}
	}
	return NewNumberExpr(numb), nil
}

func ledBinary(p Parser, left ast.Expr, bp BindingPower) (ast.Expr, error) {
	opToken := p.Next()

	right, err := p.ParseExpr(bp)
	{
		if err != nil {
			return nil, err
		}
	}

	return NewBinaryExpr(opToken.Type, left, right), nil
}
