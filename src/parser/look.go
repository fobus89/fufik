package parser

import (
	"github.com/fobus89/fufik/src/ast"
	"github.com/fobus89/fufik/src/token"
)

type BindingPower int

const (
	Lowest BindingPower = iota
	Comma
	Assigment
	Logical
	Relational
	Additive
	Muptiplicative
	Unary
	Call
	Member
	Primary
	Highest
)

type (
	StmtHandlerType = func(p *parser) (ast.Stmt, error)
	NudHandlerType  = func(p *parser) (ast.Expr, error)
	LedHandlerType  = func(p *parser, left ast.Expr, bp BindingPower) (ast.Expr, error)
)

type Handler[T comparable, E any] map[T]E

func (h Handler[T, E]) Get(key T) E {
	handler := h[key]
	return handler
}

func (h Handler[T, E]) GetOrNone(key T) (E, bool) {
	handler, ok := h[key]
	return handler, ok
}

func (h Handler[T, E]) Has(key T) bool {
	_, ok := h[key]
	return ok
}

func (h Handler[T, E]) Add(key T, val E) {
	h[key] = val
}

func (p *parser) NudRegister(kind token.TokenType, nudHander NudHandlerType) {
	p.nudLookup.Add(kind, nudHander)
}

func (p *parser) LedRegister(kind token.TokenType, bp BindingPower, ledHander LedHandlerType) {
	p.bpLookup.Add(kind, bp)
	p.ledLookup.Add(kind, ledHander)
}

func (p *parser) StmtRegister(kind token.TokenType, stmtHander StmtHandlerType) {
	p.bpLookup.Add(kind, Lowest)
	p.stmtLookup.Add(kind, stmtHander)
}
