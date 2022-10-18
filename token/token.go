package token

type TokenType string

type Token struct {
	Id      int
	Type    TokenType
	Literal string
}

const (
	H1      = "*"
	H2      = "**"
	H3      = "***"
	CONTENT = "CONTENT"
)
