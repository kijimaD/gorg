package parser

import (
	"gorg/ast"
	"gorg/lexer"
	"gorg/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

type (
	// ** Header
	prefixParseFn func() ast.Node
)

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
	}

	// 2つトークンを読み込む。curTokenとpeekTokenの両方がセットされる
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) ParseOrg() *ast.Org {
	org := &ast.Org{}
	org.Nodes = []ast.Node{}

	for p.curToken.Type != token.EOF {
		node := p.parseNode()

		if node != nil {
			org.Nodes = append(org.Nodes, node)
		}
		p.nextToken()
	}

	return org
}

/////////////
// private //
/////////////

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) parseNode() ast.Node {
	switch p.curToken.Type {
	case token.ASTERISK:
		return p.parseAsterisk()
	default:
		return nil
	}
}

func (p *Parser) parseAsterisk() ast.Node {
	header := &ast.Header{Token: p.curToken}

	header.Level = 1

	for p.expectPeek(token.ASTERISK) {
		header.Level += 1
	}

	if !p.expectPeek(token.SPACE) {
		return nil
	}

	p.nextToken()
	header.Value = p.curToken.Literal
	return header
}

// peekTokenの型をチェックし、その型が正しい場合に限ってnextTokenを読んで、トークンを進める
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}

// 次のトークンと引数の型を比較する
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}
