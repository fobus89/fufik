package parse_literal

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
