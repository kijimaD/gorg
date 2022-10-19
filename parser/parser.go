package parser

import (
	"bufio"
	"bytes"
	"gorg/ast"
	"regexp"
)

const (
	HEADER1_REGEXP = `^\* (.*)`
	HEADER2_REGEXP = `^\*\* (.*)`
)

type Parser struct {
	input string
}

type (
	// ** Header
	prefixParseFn func() ast.Node
)

func New(input string) *Parser {
	p := &Parser{
		input: input,
	}

	return p
}

func (p *Parser) ParseOrg() *ast.Org {
	org := &ast.Org{}
	org.Nodes = []ast.Node{}

	buf := bytes.NewBufferString(p.input)
	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		p.parseNode(org, scanner.Text())
	}

	return org
}

/////////////
// private //
/////////////

func (p *Parser) parseNode(o *ast.Org, s string) {
	str := s

	if len(p.parseHeader(str, HEADER1_REGEXP)) > 0 {
		value := p.parseHeader(s, HEADER1_REGEXP)

		header := &ast.Header{Level: 1}
		normal := &ast.Normal{Value: value, Parent: header}
		o.Nodes = append(o.Nodes, header)
		o.Nodes = append(o.Nodes, normal)
	} else if len(p.parseHeader(str, HEADER2_REGEXP)) > 0 {
		value := p.parseHeader(s, HEADER2_REGEXP)

		header := &ast.Header{Level: 2}
		normal := &ast.Normal{Value: value, Parent: header}
		o.Nodes = append(o.Nodes, header)
		o.Nodes = append(o.Nodes, normal)
	}
}

func (p *Parser) parseHeader(s string, exp string) string {
	re := regexp.MustCompile(exp)
	ok := re.MatchString(s)

	var match string
	matchs := re.FindStringSubmatch(s)
	if ok {
		match = matchs[1]
	} else {
		match = ""
	}

	return match
}
