package main

import (
	"fmt"
	"log"
	"math"
	"strconv"

	. "github.com/fobus89/fufik"
	"github.com/fobus89/fufik/internal/parser"
)

type FuncExpr struct {
	Value float64
}

func NewFuncExpr(value float64) *FuncExpr {
	return &FuncExpr{
		Value: value,
	}
}

func (b *FuncExpr) Eval() any {
	return b.Value
}

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
	Op    TokenType
	Right Expr
}

func NewBinaryExpr(op TokenType, left, right Expr) *BinaryExpr {
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
	case PLUS:
		return leftVal + rightVal

	case MINUS:
		return leftVal - rightVal

	case STAR:
		return leftVal * rightVal

	case SLASH:
		return leftVal / rightVal

	case PERCENT:
		return math.Mod(leftVal, rightVal)
	}

	return 0
}

func main() {
	p := parser.NewParser(`
			[]float64{1,2,3}
			[]float64{1,2,3,4,5,6,7,8,9,10,11,12}
				(1+2+3) * 3 -1 +2
		`)

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
	p.LedRegister(PLUS, parser.Additive, ledBinary)
	p.LedRegister(MINUS, parser.Additive, ledBinary)

	p.LedRegister(STAR, parser.Muptiplicative, ledBinary)
	p.LedRegister(SLASH, parser.Muptiplicative, ledBinary)
	p.LedRegister(PERCENT, parser.Muptiplicative, ledBinary)
	p.NudRegister(LBRACKET, nudArrayOrSlice)

	p.NudRegister(INT_LITERAL, nudIntLiteral)
	p.NudRegister(FLOAT_LITERAL, nudIntLiteral)
	p.NudRegister(LPARENT, nudGrouping)
}

func nudArrayOrSlice(p Parser) (Expr, error) {
	if !p.MatchNext(LBRACKET) {
		return nil, fmt.Errorf("expected LBRACKET, got %v", p.CurrentToken())
	}

	var elems []Expr

	//[]float64{1,2,3}
	if p.MatchAllNext(RBRACKET, Float64, LBRACE) {

		for !p.MatchNext(RBRACE) {

			expr, err := p.ParseExpr(parser.Lowest)
			{
				if err != nil {
					return nil, err
				}
			}

			elems = append(elems, expr)

			if !p.Match(RBRACE) && !p.MatchNext(COMMA) {
				return nil, fmt.Errorf("expected comma")
			}

		}

		return &SliceExpr{
			Type:     "float64",
			Size:     len(elems),
			Elements: elems,
		}, nil
	}

	for {
		expr, err := p.ParseExpr(parser.Lowest)
		{
			if err != nil {
				return nil, err
			}
		}

		elems = append(elems, expr)

		if p.Match(RBRACKET) {
			break
		}

		if !p.Match(COMMA) {
			return nil, fmt.Errorf("expected comma")
		}
	}

	return &ArrayExpr{Elements: elems}, nil
}

func nudGrouping(p Parser) (Expr, error) {
	if !p.MatchNext(LPARENT) {
		return nil, fmt.Errorf("expected LPARENT, got %v", p.CurrentToken())
	}

	expr, err := p.ParseExpr(parser.Lowest)
	{
		if err != nil {
			return nil, err
		}
	}

	if !p.MatchNext(RPARENT) {
		return nil, fmt.Errorf("expected RPARENT, got %v", p.CurrentToken())
	}

	return expr, nil
}

func nudIntLiteral(p Parser) (Expr, error) {
	literal := p.Next()

	numb, err := strconv.ParseFloat(literal.Literal, 64)
	{
		if err != nil {
			return nil, err
		}
	}

	return NewNumberExpr(numb), nil
}

func ledBinary(p Parser, left Expr, bp BindingPower) (Expr, error) {
	if !p.MatchAny(PLUS, MINUS, STAR, SLASH, PERCENT) {
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

// func parseType() Type {
// 	if match("[") {
// 		if isNumber(peek()) {
// 			size := parseNumber()
// 			expect("]")
// 			elem := parseType()
// 			return ArrayType{Size: size, Elem: elem}
// 		} else {
// 			expect("]")
// 			elem := parseType()
// 			return SliceType{Elem: elem}
// 		}
// 	}
//
// 	return parseBaseType() // int, float и т.д.
// }
//
// func ledIndex(p *Parser, left Expr, bp BindingPower) (Expr, error) {
//
// 	if p.Match(COLON) {
// 		// slice a[1:3]
// 		high, err := p.ParseExpr(bp)
// 		if err != nil {
// 			return nil, err
// 		}
//
// 		if !p.Match(RBRACKET) {
// 			return nil, fmt.Errorf("expected ]")
// 		}
//
// 		return &SliceExpr{
// 			Target: left,
// 			High:   high,
// 		}, nil
// 	}
//
// 	// index a[1]
// 	index, err := p.ParseExpr(bp)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	if !p.Match(RBRACKET) {
// 		return nil, fmt.Errorf("expected ]")
// 	}
//
// 	return &IndexExpr{
// 		Target: left,
// 		Index:  index,
// 	}, nil
// }
