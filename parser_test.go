package lawparser

import (
	"testing"
)

var input = "hola como estás"
var testHeader = "Este es un preambulo chafa\ncapítulo 10\n"
var testContainers = "\nCapítulo 10\nDe los perritos\nArtículo 10.-Los perritos no comerán solitos en la calle bajo ninguna circunstancia\nI.-No sé que estoy haciendo\na)pero lo estoy haciendo\nb)y tú no\nArtículo 20.-Los perritos no comerán solitos en la calle bajo ninguna circunstancia\nI.-No sé que estoy haciendo\n2.-No , yo tampoco\nArtículo 30.-Los perritos no comerán solitos en la calle bajo ninguna circunstancia\nI.-No sé que estoy haciendo\nII.-No,yo tampoco\nIII.-Yo menos"

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
	inputB := []byte(input)
	p := parser{
		lexer: lex(inputB),
		pos:   Pos(-1),
	}
	for p.peekItem(); !isItemEOF(p.item); p.peekItem() {
	}
	str := p.renderItem()
	if str != input {
		t.Error("expected '", input, "' got '", str, "'")
	}
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

func TestParseArticles(t *testing.T) {
	testContainersB := []byte(testContainers)
	p := Parse("", "", "", "", "", "", "", "", testContainersB)
	if len(p.Articles) != 3 {
		t.Error("It was expected 3 articles, found ", len(p.Articles))
	}
}

func TestParseFractions(t *testing.T) {
	testContainersB := []byte(testContainers)
	p := Parse("", "", "", "", "", "", "", "", testContainersB)
	if len(p.Articles[0].Fractions) != 1 {
		t.Error("It was expected 1 fraction not ", len(p.Articles[0].Fractions))
	}
	if len(p.Articles[1].Fractions) != 2 {
		t.Error("It was expected 2 fractions not ", len(p.Articles[1].Fractions))
	}
	if len(p.Articles[2].Fractions) != 3 {
		t.Error("It was expected 3 fractions not ", len(p.Articles[2].Fractions))
	}
}

func TestArtSubFraction(t *testing.T) {
	testContainersB := []byte(testContainers)
	p := Parse("", "", "", "", "", "", "", "", testContainersB)
	if len(p.Articles[0].Fractions[0].Sub) != 2 {
		t.Error("It was expected 2 subfractions not ", len(p.Articles[0].Fractions[0].Sub))
	}
}
