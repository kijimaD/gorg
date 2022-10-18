package token

type TokenType string

type Token struct {
	Id      int
	Type    TokenType
	Literal string
}

const (
	ASTERISK = "*"
	NEWLINE  = "NEWLINE"
	SPACE    = "SPACE"
	CONTENT  = "CONTENT"
	MINUS    = "MINUS"
	PLUS     = "PLUS"
	EOF      = "EOF"
)
