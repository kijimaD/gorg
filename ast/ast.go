package ast

import (
	"bytes"
	"fmt"
	"gorg/token"
	"strings"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Org struct {
	Nodes []Node
}

func (o *Org) TokenLiteral() string {
	if len(o.Nodes) > 0 {
		return o.Nodes[0].TokenLiteral()
	} else {
		return ""
	}
}

type Root struct{}

func (r *Root) TokenLiteral() string { return "" }
func (r *Root) String() string       { return "{type: root}" }

type Normal struct {
	Token  token.Token
	Value  string
	Parent Node
}

func (n *Normal) TokenLiteral() string { return n.Token.Literal }
func (h *Normal) String() string {
	var out bytes.Buffer
	text := fmt.Sprintf("{type: normal, Value: '%s', Parent: %T}", h.Value, h.Parent)
	out.WriteString(text)
	return out.String()
}

type Header struct {
	Level  int
	Parent Node
}

func (h *Header) TokenLiteral() string { return strings.Repeat(token.ASTERISK, h.Level) }
func (h *Header) String() string {
	var out bytes.Buffer
	text := fmt.Sprintf("{type: header, Level: %d}", h.Level)
	out.WriteString(text)
	return out.String()
}

type Bold struct {
	Parent Node
}

func (b *Bold) TokenLiteral() string { return token.ASTERISK }
func (b *Bold) String() string {
	var out bytes.Buffer
	text := fmt.Sprintf("{type: bold, Parent: %T}", b.Parent)
	out.WriteString(text)
	return out.String()
}

type Italic struct {
	Parent Node
}

func (i *Italic) TokenLiteral() string { return token.SLASH }
func (i *Italic) String() string {
	var out bytes.Buffer
	text := fmt.Sprintf("{type: italic, Parent: %T}", i.Parent)
	out.WriteString(text)
	return out.String()
}

type Comment struct {
	Parent Node
}

func (c *Comment) TokenLiteral() string { return token.NUMBER_SIGN }
func (c *Comment) String() string {
	var out bytes.Buffer
	text := fmt.Sprintf("{type: comment, Parent: %T}", c.Parent)
	out.WriteString(text)
	return out.String()
}
