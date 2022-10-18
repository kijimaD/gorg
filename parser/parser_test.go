package parser

import (
	"gorg/ast"
	"gorg/lexer"
	"strings"
	"testing"
)

func TestHeaderNodes(t *testing.T) {
	input := `
* header1
** header2
`
	l := lexer.New(input)
	o := New(l)
	org := o.ParseOrg()

	if len(org.Nodes) != 2 {
		t.Fatalf("program has not enough nodes. got=%d",
			len(org.Nodes))
	}

	for _, node := range org.Nodes {
		headerNode, ok := node.(*ast.Header)
		if !ok {
			t.Errorf("headerNode not *ast.Header. got=%T", node)
			continue
		}
		if headerNode.TokenLiteral() != strings.Repeat("*", headerNode.Level) {
			t.Errorf("headerNode.TokenLiteral not 'header'. got %q",
				node.TokenLiteral())
		}
	}
}
