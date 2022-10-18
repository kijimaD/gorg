package lexer

import (
	"testing"

	"gorg/token"
)

func TestNextToken(t *testing.T) {
	input := `
* h1
** h2
*** h3
**** h4
***** h5
content
**invalid
- list1
- list2
+ list1
+ list2
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.NEWLINE, "\n"},
		{token.ASTERISK, "*"},
		{token.SPACE, " "},
		{token.CONTENT, "h1"},
		{token.NEWLINE, "\n"},

		{token.ASTERISK, "*"},
		{token.ASTERISK, "*"},
		{token.SPACE, " "},
		{token.CONTENT, "h2"},
		{token.NEWLINE, "\n"},

		{token.ASTERISK, "*"},
		{token.ASTERISK, "*"},
		{token.ASTERISK, "*"},
		{token.SPACE, " "},
		{token.CONTENT, "h3"},
		{token.NEWLINE, "\n"},

		{token.ASTERISK, "*"},
		{token.ASTERISK, "*"},
		{token.ASTERISK, "*"},
		{token.ASTERISK, "*"},
		{token.SPACE, " "},
		{token.CONTENT, "h4"},
		{token.NEWLINE, "\n"},

		{token.ASTERISK, "*"},
		{token.ASTERISK, "*"},
		{token.ASTERISK, "*"},
		{token.ASTERISK, "*"},
		{token.ASTERISK, "*"},
		{token.SPACE, " "},
		{token.CONTENT, "h5"},
		{token.NEWLINE, "\n"},

		{token.CONTENT, "content"},
		{token.NEWLINE, "\n"},

		{token.ASTERISK, "*"},
		{token.ASTERISK, "*"},
		{token.CONTENT, "invalid"},
		{token.NEWLINE, "\n"},

		{token.MINUS, "-"},
		{token.SPACE, " "},
		{token.CONTENT, "list1"},
		{token.NEWLINE, "\n"},

		{token.MINUS, "-"},
		{token.SPACE, " "},
		{token.CONTENT, "list2"},
		{token.NEWLINE, "\n"},

		{token.PLUS, "+"},
		{token.SPACE, " "},
		{token.CONTENT, "list1"},
		{token.NEWLINE, "\n"},

		{token.PLUS, "+"},
		{token.SPACE, " "},
		{token.CONTENT, "list2"},
		{token.NEWLINE, "\n"},

		{token.EOF, ""},
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
