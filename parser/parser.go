package parser

import (
	"bufio"
	"bytes"
	"gorg/ast"
	"regexp"
	"strings"
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
	org.Nodes = append(org.Nodes, &ast.Normal{Value: "root"})

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

func (p *Parser) parseNode(o *ast.Org, str string) {
	if len(p.parseHeader(str, HEADER1_REGEXP)) > 0 {
		// header 1
		header := &ast.Header{Level: 1, Parent: o.Nodes[len(o.Nodes)-1]}
		o.Nodes = append(o.Nodes, header)

		p.parseNode(o, strings.Replace(str, "* ", "", 1))
	} else if len(p.parseHeader(str, HEADER2_REGEXP)) > 0 {
		// header 2
		header := &ast.Header{Level: 2, Parent: o.Nodes[len(o.Nodes)-1]}
		o.Nodes = append(o.Nodes, header)

		p.parseNode(o, strings.Replace(str, "** ", "", 1))
	} else if len(p.parseBold(str)) > 0 {
		// bold
		value := p.parseBold(str)
		bold := &ast.Bold{Parent: nil} // TODO: 後で入れる
		normal := &ast.Normal{Value: value, Parent: bold}
		o.Nodes = append(o.Nodes, bold)
		o.Nodes = append(o.Nodes, normal)

		str = strings.Replace(str, "*"+value+"*", "", 1)
	} else {
		// normal
		normal := &ast.Normal{Value: str, Parent: o.Nodes[len(o.Nodes)-1]}
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

func (p *Parser) parseBold(s string) string {
	re := regexp.MustCompile(`\*(.*)\*`)
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
