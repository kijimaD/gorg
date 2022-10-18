package token

type TokenType string

type Token struct {
	Id      int
	Type    TokenType
	Literal string
}

const (
	ASTERISK = "*"
	NEWLINE  = "NL"
	CONTENT  = "CONTENT"
)
