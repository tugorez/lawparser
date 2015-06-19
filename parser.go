package lawparser

import (
	"strings"
)

type parser struct {
	items []item
	item  item
	lexer *lexer
	legalDocument
	pos Pos
}

func Parse(input []byte) {
	p := &parser{
		lexer: lex(input),
		pos:   Pos(-1),
	}
	p.Header = parseHeader(p)
	p.Articles = parseContainers(p)
	p.Footer = parseFooter(p)

}

func (self *parser) peekItem() {
	if int(self.pos)+1 >= len(self.items) {
		it := self.lexer.nextItem()
		if isItemEOF(it) || isItemError(it) {
			self.items = append(self.items, item{itemEOF, ""})
		} else {
			self.items = append(self.items, it)
		}
	}
	self.pos++
	self.item = self.items[self.pos]
}

func (self *parser) confirmItem() string {
	str := ""
	for i := 0; i < int(self.pos); i++ {
		str += self.items[i].val + " "
	}
	self.items = self.items[self.pos:]
	self.pos = 0
	self.item = self.items[self.pos]
	return strings.TrimSpace(str)
}

func (self *parser) resetItem(p Pos) {
	self.pos = p
	self.item = self.items[self.pos]
}

/*parsers*/
func parseHeader(p *parser) string {
	for p.peekItem(); isItemHeader(p.item); p.peekItem() {
	}
	return p.confirmItem()
}

func parseContainers(p *parser) (arts articles) {
	order := p.item
	for isSameOrder(order, p.item) {
		arts = append(arts, parseContainer(p)...)
	}
	return
}
func parseContainer(p *parser) (arts articles) {
	if isArticle(p.item) {
		art := parseArticle(p)
		arts = append(arts, art)
	} else {
		order := p.item
		title := parseTitleContainer(p)
		articulos := parseContainer(p)
		setParents(order, title, articulos)
		arts = append(arts, articulos...)
	}
	return
}
func parseArticle(p *parser) (art article) {
	art.Id = p.item.val
	p.peekItem()
	p.confirmItem()
	art.Headers = parseArtHeaders(p)
	art.Fractions = parseArtFractions(p)
	art.Footers = parseArtFooters(p)
	return
}
func parseTitleContainer(p *parser) (title string) {
	return
}

func parseArtHeaders(p *parser) (headers []subarticle) {
	return
}
func parseArtFractions(p *parser) (headers []fraction) {
	return
}
func parseArtFooters(p *parser) (headers []subarticle) {
	return
}
func parseFooter(p *parser) string {
	for ; isItemFooter(p.item); p.peekItem() {
	}
	return p.confirmItem()
}

/*validators*/
func isItemEOF(t item) bool {
	return t.typ == itemEOF
}
func isItemError(t item) bool {
	return t.typ == itemError
}
func isItemHeader(t item) bool {
	return !isItemEOF(t) && !isItemContainer(t)
}

func isItemFooter(t item) bool {
	return false
}
func isItemContainer(t item) bool {
	return t.typ >= itemContainer
}
func isSameOrder(t1, t2 item) bool {
	return false
}

func isArticle(t1 item) bool {
	return false
}

func setParents(order item, title string, arts articles) {

}
