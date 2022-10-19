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
	if o.Nodes[0].String() != "{type: normal, Value: root, Parent: <nil>}" {
		t.Errorf("1: not match header. got=%q", o.Nodes[0].String())
	}
	if o.Nodes[1].String() != "{type: normal, Value: text, Parent: *ast.Normal}" {
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
	if o.Nodes[0].String() != "{type: normal, Value: root, Parent: <nil>}" {
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

	// インラインの演算子は、before afterノードを取っておいて別で再帰評価しないとうまくいかない
	// front*bold*back の場合、*bold*を取り除いたあと、frontとback、boldを別に評価することが必要。
	// - root
	//   - normal "front"
	//   - bold
	//     - normal "bold"
	//   - normal "back"
	if len(o.Nodes) != 4 {
		t.Fatalf("program has not enough nodes. got=%d",
			len(o.Nodes))
	}

}
