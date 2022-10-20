package parser

import (
	"testing"
)

func TestNormalNodes(t *testing.T) {
	input := `text`
	p := New(input)
	o := p.ParseOrg()

	if len(o.Nodes) != 2 {
		t.Fatalf("program has not enough nodes. got=%d",
			len(o.Nodes))
	}

	// root
	if o.Nodes[0].String() != "{type: root}" {
		t.Errorf("1: not match header. got=%q", o.Nodes[0].String())
	}
	if o.Nodes[1].String() != "{type: normal, Value: text, Parent: *ast.Root}" {
		t.Errorf("1: not match header. got=%q", o.Nodes[0].String())
	}
}

func TestHeaderNodes(t *testing.T) {
	input := `* header1
** header2`
	p := New(input)
	o := p.ParseOrg()

	if len(o.Nodes) != 5 {
		t.Fatalf("program has not enough nodes. got=%d",
			len(o.Nodes))
	}

	if o.Nodes[0].String() != "{type: root}" {
		t.Errorf("0: not match header")
	}
	if o.Nodes[1].String() != "{type: header, Level: 1}" {
		t.Errorf("1: not match header")
	}
	if o.Nodes[2].String() != "{type: normal, Value: header1, Parent: *ast.Header}" {
		t.Errorf("2: not match header")
	}
	if o.Nodes[3].String() != "{type: header, Level: 2}" {
		t.Errorf("3: not match header")
	}
	if o.Nodes[4].String() != "{type: normal, Value: header2, Parent: *ast.Header}" {
		t.Errorf("4: not match header")
	}
}

func TestBoldNodes(t *testing.T) {
	input := `front*bold*back`
	p := New(input)
	o := p.ParseOrg()

	// - root
	//   - normal "front"
	//   - bold
	//     - normal "bold"
	//   - normal "back"

	if len(o.Nodes) != 5 {
		t.Fatalf("program has not enough nodes. got=%d",
			len(o.Nodes))
	}

	if o.Nodes[0].String() != "{type: root}" {
		t.Errorf("0: not match header")
	}
	if o.Nodes[1].String() != "{type: normal, Value: front, Parent: *ast.Root}" {
		t.Errorf("1: not match header")
	}
	if o.Nodes[2].String() != "{type: bold, Parent: *ast.Root}" {
		t.Errorf("2: not match header")
	}
	if o.Nodes[3].String() != "{type: normal, Value: bold, Parent: *ast.Bold}" {
		t.Errorf("3: not match header")
	}
	if o.Nodes[4].String() != "{type: normal, Value: back, Parent: *ast.Normal}" {
		t.Errorf("4: not match header")
	}
}

func TestItalicNodes(t *testing.T) {
	input := `front/italic/back`
	p := New(input)
	o := p.ParseOrg()

	if len(o.Nodes) != 5 {
		t.Fatalf("program has not enough nodes. got=%d",
			len(o.Nodes))
	}

	if o.Nodes[0].String() != "{type: root}" {
		t.Errorf("0: not match header")
	}
	if o.Nodes[1].String() != "{type: normal, Value: front, Parent: *ast.Root}" {
		t.Errorf("1: not match header")
	}
	if o.Nodes[2].String() != "{type: italic, Parent: *ast.Root}" {
		t.Errorf("2: not match header")
	}
	if o.Nodes[3].String() != "{type: normal, Value: italic, Parent: *ast.Italic}" {
		t.Errorf("3: not match header")
	}
	if o.Nodes[4].String() != "{type: normal, Value: back, Parent: *ast.Normal}" {
		t.Errorf("4: not match header")
	}
}

func TestCommentNodes(t *testing.T) {
	// コメント直下のノードは再帰パースせずnormalにする
	input := `normal
# *comment*`
	p := New(input)
	o := p.ParseOrg()

	if len(o.Nodes) != 4 {
		t.Fatalf("program has not enough nodes. got=%d",
			len(o.Nodes))
	}

	if o.Nodes[0].String() != "{type: root}" {
		t.Errorf("0: not match header")
	}
	if o.Nodes[1].String() != "{type: normal, Value: normal, Parent: *ast.Root}" {
		t.Errorf("1: not match header")
	}
	if o.Nodes[2].String() != "{type: comment, Parent: *ast.Root}" {
		t.Errorf("2: not match header")
	}
	if o.Nodes[3].String() != "{type: normal, Value: *comment*, Parent: *ast.Comment}" {
		t.Errorf("3: not match header")
	}
}
