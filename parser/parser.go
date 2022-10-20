package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"gorg/ast"
	"gorg/token"
	"regexp"
)

const (
	HEADER1_REGEXP = `^\* (.*)`
	HEADER2_REGEXP = `^\*\* (.*)`
	COMMENT_REGEXP = `^# (.*)`
)

type Parser struct {
	input string
}

func New(input string) *Parser {
	p := &Parser{
		input: input,
	}

	return p
}

func (p *Parser) ParseOrg() *ast.Org {
	org := &ast.Org{}
	org.Nodes = []ast.Node{}
	org.Nodes = append(org.Nodes, &ast.Root{})

	buf := bytes.NewBufferString(p.input)
	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		p.parseNode(org, scanner.Text(), org.Nodes[0]) // Parent: root
	}

	return org
}

/////////////
// private //
/////////////

func (p *Parser) parseNode(o *ast.Org, str string, parent ast.Node) {
	if len(p.parseHeader(str, HEADER1_REGEXP)) > 0 {
		// header 1
		header := &ast.Header{Level: 1, Parent: parent}
		o.Nodes = append(o.Nodes, header)

		p.parseNode(o, p.parseHeader(str, HEADER1_REGEXP), header)
	} else if len(p.parseHeader(str, HEADER2_REGEXP)) > 0 {
		// header 2
		header := &ast.Header{Level: 2, Parent: parent}
		o.Nodes = append(o.Nodes, header)

		p.parseNode(o, p.parseHeader(str, HEADER2_REGEXP), header)
	} else if len(p.parseComment(str)) > 0 {
		// comment
		comment := &ast.Comment{Parent: parent}
		o.Nodes = append(o.Nodes, comment)

		normal := &ast.Normal{Value: p.parseComment(str), Parent: comment}
		o.Nodes = append(o.Nodes, normal)
	} else if len(p.matchInfixTag(str, token.ASTERISK)) > 0 {
		// bold
		left, center, right := p.parseInfixTag(str, token.ASTERISK)

		p.parseNode(o, left, parent)

		bold := &ast.Bold{Parent: parent}
		o.Nodes = append(o.Nodes, bold)
		p.parseNode(o, center, bold)

		p.parseNode(o, right, parent)
	} else if len(p.matchInfixTag(str, token.SLASH)) > 0 {
		// italic
		left, center, right := p.parseInfixTag(str, token.SLASH)

		p.parseNode(o, left, parent)

		italic := &ast.Italic{Parent: parent}
		o.Nodes = append(o.Nodes, italic)
		p.parseNode(o, center, italic)

		p.parseNode(o, right, parent)
	} else if str != "" {
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

func (p *Parser) parseInfixTag(s string, tag string) (string, string, string) {
	matches := p.matchInfixTag(s, tag)

	var left string
	var center string
	var right string
	if len(matches) == 4 {
		left = matches[1]
		center = matches[2]
		right = matches[3]
	} else if len(matches) == 2 {
		left = ""
		center = matches[1]
		right = ""
	} else if len(matches) == 3 && p.matchLeftTag(s, tag) && !p.matchRightTag(s, tag) {
		left = matches[1]
		center = matches[2]
		right = ""
	} else if len(matches) == 3 && !p.matchLeftTag(s, tag) && p.matchRightTag(s, tag) {
		left = ""
		center = matches[1]
		right = matches[2]

	}

	return left, center, right
}

func (p *Parser) matchInfixTag(s string, tag string) []string {
	re := regexp.MustCompile(fmt.Sprintf(`(.*)\%s(.*)\%s(.*)`, tag, tag))
	matchs := re.FindStringSubmatch(s)
	return matchs
}
func (p *Parser) matchLeftTag(s string, tag string) bool {
	re := regexp.MustCompile(fmt.Sprintf(`(.*)\%s(.*)\%s`, tag, tag))
	return re.MatchString(s)
}

func (p *Parser) matchRightTag(s string, tag string) bool {
	re := regexp.MustCompile(fmt.Sprintf(`\%s(.*)\%s(.*)`, tag, tag))
	return re.MatchString(s)
}

func (p *Parser) parseComment(s string) string {
	re := regexp.MustCompile(COMMENT_REGEXP)
	matches := re.FindStringSubmatch(s)
	var match string
	if len(matches) > 0 {
		match = matches[1]
	} else {
		match = ""
	}
	return match
}
