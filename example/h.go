package main

// // array
// type ArrayExpr struct {
// 	Type     string
// 	Size     int
// 	Elements []ast.Expr
// }

// func NewArrayExpr(size int, ty string, elements []ast.Expr) *ArrayExpr {
// 	return &ArrayExpr{
// 		Type:     ty,
// 		Size:     size,
// 		Elements: elements,
// 	}
// }

// func (a *ArrayExpr) Eval() any {
// 	res := make([]any, len(a.Elements))

// 	for i, e := range a.Elements {
// 		res[i] = e.Eval()
// 	}

// 	return res
// }

// // slice a[1:3]
// type SliceExpr struct {
// 	Type     string
// 	Size     int
// 	Elements []ast.Expr
// 	Low      ast.Expr // optional
// 	High     ast.Expr // optional
// }

// func NewSliceExpr(target, low, high ast.Expr) *SliceExpr {
// 	return &SliceExpr{}
// }

// func (s *SliceExpr) Eval() any {
// 	var slice []any

// 	for _, el := range s.Elements {
// 		slice = append(slice, el.Eval())
// 	}

// 	return slice
// 	// arr := s.Target.Eval().([]any)
// 	//
// 	// low := 0
// 	// high := len(arr)
// 	//
// 	// if s.Low != nil {
// 	// 	low = int(s.Low.Eval().(float64))
// 	// }
// 	// if s.High != nil {
// 	// 	high = int(s.High.Eval().(float64))
// 	// }
// 	//
// 	// return arr[low:high]
// }

// type IndexExpr struct {
// 	Target ast.Expr
// 	Index  ast.Expr
// }
