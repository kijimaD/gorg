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
**invalid
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.NEWLINE, "\n"},
		{token.ASTERISK, "*"},
		{token.CONTENT, "h1"},
		{token.NEWLINE, "\n"},
		{token.CONTENT, "content"},
		{token.NEWLINE, "\n"},
		{token.ASTERISK, "*"},
		{token.ASTERISK, "*"},
		{token.CONTENT, "h2"},
		{token.NEWLINE, "\n"},
		{token.CONTENT, "foo"},
		{token.NEWLINE, "\n"},
		{token.ASTERISK, "*"},
		{token.ASTERISK, "*"},
		{token.ASTERISK, "*"},
		{token.CONTENT, "h3"},
		{token.NEWLINE, "\n"},
		{token.CONTENT, "bar"},
		{token.NEWLINE, "\n"},
		{token.ASTERISK, "*"},
		{token.ASTERISK, "*"},
		{token.CONTENT, "invalid"},
		{token.NEWLINE, "\n"},
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
