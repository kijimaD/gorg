package parser

import (
	"testing"
)

func TestNormalNodes(t *testing.T) {
	input := `text1

text2`
	p := New(input)
	o := p.ParseOrg()

	if len(o.Nodes) != 4 {
		t.Fatalf("program has not enough nodes. got=%d",
			len(o.Nodes))
	}

	tests := []struct {
		input string
	}{
		{"{type: root}"},
		{"{type: normal, Value: 'text1', Parent: *ast.Root}"},
		{"{type: normal, Value: '', Parent: *ast.Root}"},
		{"{type: normal, Value: 'text2', Parent: *ast.Root}"},
	}

	for i, tt := range tests {
		if o.Nodes[i].String() != tt.input {
			t.Errorf("%d: not match header got=%q", i, o.Nodes[i].String())
		}
	}
}

func TestHeaderNodes(t *testing.T) {
	input := `* header1
text1
** header2
text2`
	p := New(input)
	o := p.ParseOrg()

	if len(o.Nodes) != 7 {
		t.Fatalf("program has not enough nodes. got=%d",
			len(o.Nodes))
	}

	tests := []struct {
		input string
	}{
		{"{type: root}"},
		{"{type: header, Level: 1}"},
		{"{type: normal, Value: 'header1', Parent: *ast.Header}"},
		{"{type: normal, Value: 'text1', Parent: *ast.Root}"},
		{"{type: header, Level: 2}"},
		{"{type: normal, Value: 'header2', Parent: *ast.Header}"},
		{"{type: normal, Value: 'text2', Parent: *ast.Root}"},
	}

	for i, tt := range tests {
		if o.Nodes[i].String() != tt.input {
			t.Errorf("%d: not match header got=%q", i, o.Nodes[i].String())
		}
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

	tests := []struct {
		input string
	}{
		{"{type: root}"},
		{"{type: normal, Value: 'front', Parent: *ast.Root}"},
		{"{type: bold, Parent: *ast.Root}"},
		{"{type: normal, Value: 'bold', Parent: *ast.Bold}"},
		{"{type: normal, Value: 'back', Parent: *ast.Root}"},
	}

	for i, tt := range tests {
		if o.Nodes[i].String() != tt.input {
			t.Errorf("%d: not match header got=%q", i, o.Nodes[i].String())
		}
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

	tests := []struct {
		input string
	}{
		{"{type: root}"},
		{"{type: normal, Value: 'front', Parent: *ast.Root}"},
		{"{type: italic, Parent: *ast.Root}"},
		{"{type: normal, Value: 'italic', Parent: *ast.Italic}"},
		{"{type: normal, Value: 'back', Parent: *ast.Root}"},
	}

	for i, tt := range tests {
		if o.Nodes[i].String() != tt.input {
			t.Errorf("%d: not match header got=%q", i, o.Nodes[i].String())
		}
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

	tests := []struct {
		input string
	}{
		{"{type: root}"},
		{"{type: normal, Value: 'normal', Parent: *ast.Root}"},
		{"{type: comment, Parent: *ast.Root}"},
		{"{type: normal, Value: '*comment*', Parent: *ast.Comment}"},
	}

	for i, tt := range tests {
		if o.Nodes[i].String() != tt.input {
			t.Errorf("%d: not match header got=%q", i, o.Nodes[i].String())
		}
	}
}
