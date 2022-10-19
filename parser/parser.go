package parser

import (
	"bufio"
	"bytes"
	"gorg/ast"
	"regexp"
	"strings"
)

const (
	HEADER1_REGEXP    = `^\* (.*)`
	HEADER2_REGEXP    = `^\*\* (.*)`
	BOLD_REGEXP       = `(.*)\*(.*)\*(.*)`
	BOLD_LEFT_REGEXP  = `(.*)\*(.*)\*`
	BOLD_RIGHT_REGEXP = `\*(.*)\*(.*)`
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
	org.Nodes = append(org.Nodes, &ast.Normal{Value: "root"})

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

		p.parseNode(o, strings.Replace(str, "* ", "", 1), header)
	} else if len(p.parseHeader(str, HEADER2_REGEXP)) > 0 {
		// header 2
		header := &ast.Header{Level: 2, Parent: parent}
		o.Nodes = append(o.Nodes, header)

		p.parseNode(o, strings.Replace(str, "** ", "", 1), header)
	} else if len(p.parseComment(str)) > 0 {
		comment := &ast.Comment{Parent: parent}
		o.Nodes = append(o.Nodes, comment)

		normal := &ast.Normal{Value: p.parseComment(str), Parent: comment}
		o.Nodes = append(o.Nodes, normal)
	} else if len(p.parseBold(str)) > 0 {
		// bold
		matches := p.parseBold(str)

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
		} else if len(matches) == 3 && p.boldLeftMatch(str) && !p.boldRightMatch(str) {
			left = matches[1]
			center = matches[2]
			right = ""
		} else if len(matches) == 3 && !p.boldLeftMatch(str) && p.boldRightMatch(str) {
			left = ""
			center = matches[1]
			right = matches[2]

		}
		p.parseNode(o, left, parent)

		bold := &ast.Bold{Parent: parent}
		o.Nodes = append(o.Nodes, bold)
		p.parseNode(o, center, bold)

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

func (p *Parser) parseBold(s string) []string {
	re := regexp.MustCompile(BOLD_REGEXP)
	matchs := re.FindStringSubmatch(s)
	return matchs
}
func (p *Parser) boldLeftMatch(s string) bool {
	re := regexp.MustCompile(BOLD_LEFT_REGEXP)
	return re.MatchString(s)
}

func (p *Parser) boldRightMatch(s string) bool {
	re := regexp.MustCompile(BOLD_RIGHT_REGEXP)
	return re.MatchString(s)
}

func (p *Parser) parseComment(s string) string {
	re := regexp.MustCompile(`^# (.*)`)
	matches := re.FindStringSubmatch(s)
	var match string
	if len(matches) > 0 {
		match = matches[1]
	} else {
		match = ""
	}
	return match
}
