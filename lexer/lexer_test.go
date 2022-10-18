package lexer

import (
	"testing"

	"gorg/token"
)

func TestNextToken(t *testing.T) {
	input := `
* h1
content
** h2
foo
*** h3
bar
`
	// **invalid

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.H1, "*"},
		{token.CONTENT, "h1"},
		{token.CONTENT, "content"},
		{token.H2, "**"},
		{token.CONTENT, "h2"},
		{token.CONTENT, "foo"},
		{token.H3, "***"},
		{token.CONTENT, "h3"},
		{token.CONTENT, "bar"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong: expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
