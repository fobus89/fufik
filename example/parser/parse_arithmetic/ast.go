package parse_arithmetic

import (
	"math"

	"github.com/fobus89/fufik"
)

var _ fufik.Expr = (*BinaryExpr)(nil)

type BinaryExpr struct {
	Left  fufik.Expr
	Op    fufik.TokenType
	Right fufik.Expr
}

func NewBinaryExpr(op fufik.TokenType, left, right fufik.Expr) *BinaryExpr {
	return &BinaryExpr{
		Left:  left,
		Op:    op,
		Right: right,
	}
}

func (b *BinaryExpr) Eval() (any, string) {

	leftExpr := b.Left
	rightExpr := b.Right

	var leftVal float64
	{
		if err := leftExpr.Out(&leftExpr); err != nil {

		}
	}

	var rightVal float64
	{
		if err := rightExpr.Out(&rightVal); err != nil {

		}
	}

	switch b.Op {
	case fufik.PLUS:
		return leftVal + rightVal, "Binary"

	case fufik.MINUS:
		return leftVal - rightVal, "Binary"

	case fufik.STAR:
		return leftVal * rightVal, "Binary"

	case fufik.SLASH:
		return leftVal / rightVal, "Binary"

	case fufik.PERCENT:
		return math.Mod(leftVal, rightVal), "Binary"
	}

	return 0, "Binary"
}

// Out implements [ast.Expr].
func (b *BinaryExpr) Out(any) error {
	panic("unimplemented")
}

// Type implements [ast.Expr].
func (b *BinaryExpr) Type() string {
	return "Binary"
}
