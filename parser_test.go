package lawparser

import (
	//	"fmt"
	"testing"
)

var input = "hola como estás"
var testHeader = "Este es un preambulo chafa\ncapítulo 10\n"

func TestPeekItem(t *testing.T) {
	inputB := []byte(input)
	p := parser{
		lexer: lex(inputB),
		pos:   Pos(-1),
	}
	p.peekItem()
	if p.item.typ != itemWord || p.item.val != "hola" {
		t.Error("expected hola got", p.item)
	}
	p.peekItem()
	if p.item.typ != itemWord || p.item.val != "como" {
		t.Error("expected como got", p.item)
	}
	p.peekItem()
	if p.item.typ != itemWord || p.item.val != "estás" {
	}
	p.peekItem()
	if p.item.typ != itemEOF {
		t.Error("expected itemEOF", p.item)
	}
}

func TestResetItem(t *testing.T) {
	inputB := []byte(input)
	p := parser{
		lexer: lex(inputB),
		pos:   Pos(-1),
	}
	p.peekItem()
	backup := p.pos
	p.peekItem()
	p.resetItem(backup)
	if p.item.typ != itemWord || p.item.val != "hola" {
		t.Error("expected hola got", p.item)
	}
	p.peekItem()
	if p.item.typ != itemWord || p.item.val != "como" {
		t.Error("expected como got", p.item)
	}
	p.peekItem()
	p.peekItem()
	p.resetItem(backup)
	if p.item.typ != itemWord || p.item.val != "hola" {
		t.Error("expected hola got", p.item)
	}
}

func TestConfirmItem(t *testing.T) {
	inputB := []byte(input)
	p := parser{
		lexer: lex(inputB),
		pos:   Pos(-1),
	}
	p.peekItem()
	p.peekItem()
	p.peekItem()
	str := p.confirmItem()
	if str != "hola como" || p.pos != Pos(0) ||
		p.item.typ != itemWord || p.item.val != "estás" {
		t.Error(str, p.pos, p.item)
	}
}
func TestRender(t *testing.T) {
	t.Error("Nothing was wrote yet")
}
func TestParseHeader(t *testing.T) {
	testHeaderB := []byte(testHeader)
	p := &parser{
		lexer: lex(testHeaderB),
		pos:   Pos(-1),
	}
	parseHeader(p)
	if p.item.typ != itemChapter || p.item.val != "10" {
		t.Error("expected itemChapter with val 10, got ", p.item)
	}
}
