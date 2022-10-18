package ast

import (
	"bytes"
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

type Header struct {
	Token token.Token
	Level int
	Value string // TODO: ここにも何かしらマークがある可能性があるので、あとでnodeを返すようにしたい
}

func (h *Header) TokenLiteral() string { return strings.Repeat(h.Token.Literal, h.Level) }
func (h *Header) String() string {
	var out bytes.Buffer

	out.WriteString(strings.Repeat(h.Token.Literal, h.Level))
	out.WriteString(" ")
	out.WriteString(h.Value)

	return out.String()
}
