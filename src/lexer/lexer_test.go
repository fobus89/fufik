package lexer

import "testing"

func Test_Lexer(t *testing.T) {
	l := NewLexer(`
			" as { 1 + 2 + "asdsa" } bs "	
				"123"
		`)

	t.Log(l.Tokens())
}
