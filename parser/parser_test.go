package parser

import (
	"gorg/lexer"
	"testing"
)

func TestHeaderStatements(t *testing.T) {
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
}
