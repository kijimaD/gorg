package lexer

import (
	"gorg/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// 現在の1文字を読み込んでトークンを返す
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '*':
		if l.peekChar() == '*' {
			l.readChar()
			if l.peekChar() == '*' {
				l.readChar()
				tok = token.Token{Type: token.H3, Literal: "***"}
			} else {
				tok = token.Token{Type: token.H2, Literal: "**"}
			}
		} else {
			tok = newToken(token.H1, l.ch)
		}
	default:
		tok.Type = token.CONTENT
		tok.Literal = l.readString()
	}

	l.readChar()
	return tok
}

/////////////////
// Private     //
/////////////////

// 1文字分解読位置を進める
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // NUL
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// のぞき見(peek)。readChar()の、文字解析器を進めずないバージョン。先読みだけを行う
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition] // 次の位置を返す
	}
}

// 文字列を読み取る
func (l *Lexer) readString() string {
	initial_position := l.position
	for {
		l.readChar()
		if l.ch == '\n' {
			break
		}
	}
	return l.input[initial_position:l.position]
}

// 半角スペースを読み飛ばす
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// トークンを生成する
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
