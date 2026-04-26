package parser

import (
	"slices"

	"github.com/fobus89/fufik/src/ast"
	"github.com/fobus89/fufik/src/token"
)

type (
	StmtLookupType = Handler[token.TokenType, StmtHandlerType]
	NudLookupType  = Handler[token.TokenType, NudHandlerType]
	LedLookupType  = Handler[token.TokenType, LedHandlerType]
	BpLookupType   = Handler[token.TokenType, BindingPower]
)

type Parser interface {
	ParseExpr(bp BindingPower) (ast.Expr, error)
	ParseStmt() (ast.Stmt, error)
}

type parser struct {
	Nodes      []token.Token
	stmtLookup StmtLookupType
	nudLookup  NudLookupType
	ledLookup  LedLookupType
	bpLookup   BpLookupType
	pos        int
}

func (p *parser) Parse() (*ast.BlockStmt, error) {
	var body []ast.Stmt

	for p.HasToken() {
		stmt, err := p.parseStmt()
		{
			if err != nil {
				return nil, err
			}
		}

		body = append(body, stmt)
	}

	return ast.NewBlockStmt(body), nil
}

func (p *parser) CurrentToken() token.Token {
	return p.peek(0)
}

func (p *parser) peek(offset int) token.Token {
	pos := p.pos + offset

	if len(p.Nodes) < pos {
		return token.NewToken(token.EOF, "")
	}

	node := p.Nodes[p.pos]

	p.pos = pos

	return node
}

func (p *parser) Next() token.Token {
	return p.peek(1)
}

func (p *parser) CurrentTokenKind() token.TokenType {
	return p.CurrentToken().Type
}

func (p *parser) HasToken() bool {
	return p.pos < len(p.Nodes) && p.CurrentTokenKind() != token.EOF
}

func (p *parser) MatchAll(tokens ...token.TokenType) bool {
	for _, tok := range tokens {
		if !p.Match(tok) {
			return false
		}
	}

	return true
}

func (p *parser) MatchAllNext(tokens ...token.TokenType) bool {
	if !p.MatchAll(tokens...) {
		return false
	}

	p.pos += len(tokens)

	return true
}

func (p *parser) MatchAny(tokens ...token.TokenType) bool {
	return slices.ContainsFunc(tokens, p.Match)
}

func (p *parser) MatchAnyNext(tokens ...token.TokenType) bool {
	return slices.ContainsFunc(tokens, p.MatchNext)
}

func (p *parser) Match(tok token.TokenType) bool {
	return p.CurrentTokenKind() == tok
}

func (p *parser) MatchNext(tok token.TokenType) bool {
	return p.MatchNextN(tok, 1)
}

func (p *parser) MatchNextN(tok token.TokenType, n int) bool {
	if p.Match(tok) {
		p.peek(n)
		return true
	}

	return false
}
