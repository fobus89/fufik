package token

import (
	"fmt"
)

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Col     int
	Tokens  []Token
}

func (t Token) String() string {
	return fmt.Sprintf(
		"%s{ (%d:%d) %s }",
		t.Literal,
		t.Line,
		t.Col,
		t.Type,
	)
}

func NewEOF(line, col int) Token {
	return Token{
		Type:    EOF,
		Literal: "",
		Line:    line,
		Col:     col,
	}
}

func NewToken(t TokenType, value string, line, col int, tokens ...Token) Token {
	return Token{
		Type:    t,
		Literal: value,
		Line:    line,
		Col:     col,
		Tokens:  tokens,
	}
}
