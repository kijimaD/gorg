package parser

import (
	"testing"
)

func TestNormalNodes(t *testing.T) {
	input := `text`
	p := New(input)
	o := p.ParseOrg()

	if len(o.Nodes) != 1 {
		t.Fatalf("program has not enough nodes. got=%d",
			len(o.Nodes))
	}

	if o.Nodes[0].String() != "{type: normal, Value: text}" {
		t.Errorf("1: not match header")
	}
}

func TestHeaderNodes(t *testing.T) {
	input := `* header1
** header2`
	p := New(input)
	o := p.ParseOrg()

	if len(o.Nodes) != 4 {
		t.Fatalf("program has not enough nodes. got=%d",
			len(o.Nodes))
	}

	if o.Nodes[0].String() != "{type: header, Level: 1}" {
		t.Errorf("1: not match header")
	}
	if o.Nodes[1].String() != "{type: normal, Value: header1}" {
		t.Errorf("2: not match header")
	}
	if o.Nodes[2].String() != "{type: header, Level: 2}" {
		t.Errorf("3: not match header")
	}
	if o.Nodes[3].String() != "{type: normal, Value: header2}" {
		t.Errorf("4: not match header")
	}
}

func TestBoldNodes(t *testing.T) {
	input := `*bold*`
	p := New(input)
	o := p.ParseOrg()

	if len(o.Nodes) != 2 {
		t.Fatalf("program has not enough nodes. got=%d",
			len(o.Nodes))
	}
}
