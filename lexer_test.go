package lawparser

import (
	"testing"
)

var textest = "abcdefghijkl"
var spaces = " 	 	 	"
var newlines = "\n \n \n"
var words = "comida comer perro"
var digits = "12 34 56 1000 10"
var reserveds = "artículo capítulo sección Sección Capitulo"
var romans = "I II III V X XXX"
var ordinals = "Primero segundo tercero cuarto quinto quinta"
var latins = "Bis bis ter Ter"
var symbols = "@ * . - ¬ | ° ¨ ^ .-"
var containers = "\nartículo 1.-\nCapítulo 10\n\nArtículo primero.-\nArtículo décimo.-\nSección 30 bis\n"
var subarticlelet = "\na)\nb)\nc)"
var subarticlenum = "\n1)\nprimero)\n1 bis)"
var undefinedB = []byte{1, 2, 3, 4}

func TestPeek(t *testing.T) {
	textestB := []byte(textest)
	l := lex(textestB)
	if l.char != 'a' {
		t.Error("expected a got ", string(l.char))
	}
	backup := l.pos
	l.peek()
	l.peek()
	if l.char != 'c' {
		t.Error("expected c got ", string(l.char))
	}
	l.peek()
	l.peek()
	if l.char != 'e' {
		t.Error("expected e got ", string(l.char))
	}
	l.reset(backup)
	if l.char != 'a' {
		t.Error("expected a got ", string(l.char))
	}
}
func TestReset(t *testing.T) {
	textestB := []byte(textest)
	l := lex(textestB)
	if l.char != 'a' {
		t.Error("expected a got ", string(l.char))
	}
	backup := l.pos //a
	l.peek()
	l.peek()
	backup2 := l.pos //c
	l.peek()
	l.peek()
	l.peek()
	if l.char != 'f' {
		t.Error("expected f got ", string(l.char))
	}
	l.reset(backup2)
	if l.char != 'c' {
		t.Error("expected c got ", string(l.char))
	}
	l.reset(backup)
	if l.char != 'a' {
		t.Error("expected a got ", string(l.char))
	}
}
func TestConfirm(t *testing.T) {
	textestB := []byte(textest)
	l := lex(textestB)
	if l.char != 'a' {
		t.Error("expected a got ", string(l.char))
	}
	l.peek()
	l.peek()
	l.confirm()
	if l.char != 'c' {
		t.Error("expected c got ", string(l.char))
	}
	if l.pos != Pos(0) {
		t.Error("expected pos==0, got pos==", l.pos)
	}
}
func TestLexSpace(t *testing.T) {
	spacesB := []byte(spaces)
	l := lex(spacesB)
	for it := l.nextItem(); it.typ != itemEOF; it = l.nextItem() {
	}
}
func TestLexNewLine(t *testing.T) {
	newlinesB := []byte(newlines)
	l := lex(newlinesB)
	for item := l.nextItem(); item.typ != itemEOF; item = l.nextItem() {
		if item.typ != itemNewLine {
			t.Error("Expected itemNewLine got ", item)
		}
	}
}

func TestLexWord(t *testing.T) {
	wordsB := []byte(words)
	l := lex(wordsB)
	for item := l.nextItem(); item.typ != itemEOF; item = l.nextItem() {
		if item.typ != itemWord {
			t.Error("Expected itemWord got ", item)
		}
	}
}

func TestLexDigits(t *testing.T) {
	digitsB := []byte(digits)
	l := lex(digitsB)
	for item := l.nextItem(); item.typ != itemEOF; item = l.nextItem() {
		if item.typ != itemNumber {
			t.Error("Expected itemNumber got ", item)
		}
	}
}

func TestLexSymbols(t *testing.T) {
	symbolsB := []byte(symbols)
	l := lex(symbolsB)
	for item := l.nextItem(); item.typ != itemEOF; item = l.nextItem() {
		if item.typ != itemSymbol {
			t.Error("Expected itemSymbol got ", item)
		}
	}
}

func TestLexReserveds(t *testing.T) {
	reservedsB := []byte(reserveds)
	l := lex(reservedsB)
	for item := l.nextItem(); item.typ != itemEOF; item = l.nextItem() {
		if item.typ != itemReserved {
			t.Error("Expected itemReserved got ", item)
		}
	}
}
func TestLexOrdinals(t *testing.T) {
	ordinalsB := []byte(ordinals)
	l := lex(ordinalsB)
	for item := l.nextItem(); item.typ != itemEOF; item = l.nextItem() {
		if item.typ != itemOrdinal {
			t.Error("Expected itemOrdinal got ", item)
		}
	}
}
func TestLexRomans(t *testing.T) {
	romansB := []byte(romans)
	l := lex(romansB)
	for item := l.nextItem(); item.typ != itemEOF; item = l.nextItem() {
		if item.typ != itemRoman {
			t.Error("Expected itemRoman got ", item)
		}
	}
}
func TestLexLatins(t *testing.T) {
	latinsB := []byte(latins)
	l := lex(latinsB)
	for item := l.nextItem(); item.typ != itemEOF; item = l.nextItem() {
		if item.typ != itemLatin {
			t.Error("Expected itemLatin got ", item)
		}
	}
}
func TestLexContainers(t *testing.T) {
	containersB := []byte(containers)
	l := lex(containersB)
	for item := l.nextItem(); item.typ != itemEOF; item = l.nextItem() {
		if item.typ < itemContainer {
			t.Error("Expected itemContainer got ", item)
		}
	}
}

func TestLexSubArticleLet(t *testing.T) {
	subarticleletB := []byte(subarticlelet)
	l := lex(subarticleletB)
	for item := l.nextItem(); item.typ != itemEOF; item = l.nextItem() {
		if item.typ != itemSubArticleLet {
			t.Error("Expected itemSubArticleLet got ", item)
		}
	}
}
func TestLexSubArticleNum(t *testing.T) {
	subarticlenumB := []byte(subarticlenum)
	l := lex(subarticlenumB)
	for item := l.nextItem(); item.typ != itemEOF; item = l.nextItem() {
		if item.typ != itemSubArticleNum {
			t.Error("Expected itemSubArticleNum got ", item)
		}
	}
}

func TestLexUndefined(t *testing.T) {
	l := lex(undefinedB)
	for item := l.nextItem(); item.typ != itemEOF; item = l.nextItem() {
		t.Error("Expecting no items here, got ", item)
	}
}
