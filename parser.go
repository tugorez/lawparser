package lawparser

import (
	"encoding/json"
	"strings"
)

type parser struct {
	items []item
	item  item
	lexer *lexer
	legalDocument
	pos Pos
}

func (self parser) Json() []byte {
	d, _ := json.Marshal(self)
	return d
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
	str := self.renderItem()
	self.items = self.items[self.pos:]
	self.pos = 0
	self.item = self.items[self.pos]
	return str
}

func (self *parser) resetItem(p Pos) {
	self.pos = p
	self.item = self.items[self.pos]
}

func (self *parser) renderItem() string {
	str := ""
	for i := 0; i < int(self.pos); i++ {
		str += self.items[i].val + " "
	}
	return strings.TrimSpace(str)
}

func Parse(name, country, state, town, order, category, topic, url string, input []byte) (p *parser, err error) {
	l := legalDocument{
		Name:          name,
		Country:       country,
		State:         state,
		Town:          town,
		Order:         order,
		LegalCategory: category,
		Topic:         topic,
		Url:           url,
	}
	p = &parser{
		lexer:         lex(input),
		pos:           Pos(-1),
		legalDocument: l,
	}
	p.Header = parseHeader(p)
	p.Articles = parseContainers(p)
	p.Footer = parseFooter(p)
	return p
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
	} else if isItemContainer(p.item) {
		order := p.item
		p.peekItem()
		p.confirmItem()
		title := parseTitleContainer(p)
		child := p.item
		for isSameOrder(child, p.item) {
			articulos := parseContainer(p)
			arts = append(arts, articulos...)
		}
		setParents(order, title, arts)
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
	for ; isTitle(p.item); p.peekItem() {
	}
	title = p.confirmItem()
	return
}

func parseArtHeaders(p *parser) (headers []subarticle) {
	for isArtHeader(p.item) {
		headers = append(headers, parseArtHeader(p))
	}
	return
}
func parseArtHeader(p *parser) (header subarticle) {
	for ; isArtHeader(p.item) && !isDot(p.item); p.peekItem() {
	}
	if isDot(p.item) {
		p.peekItem()
	}
	header.Body = p.confirmItem()
	return
}
func parseArtFractions(p *parser) (fractions []fraction) {
	for isItemSubArticle(p.item) {
		fractions = append(fractions, parseArtFraction(p))
	}
	return
}
func parseArtFraction(p *parser) (f fraction) {
	f.Num = p.item.val
	fracType := p.item
	p.peekItem()
	p.confirmItem()
	f.Body = parseArtFractionBody(p)
	f.Sub = parseArtFractionSubs(p, fracType)
	return
}

func parseArtFractionBody(p *parser) string {
	for ; isArtFractionBody(p.item); p.peekItem() {
	}
	return p.confirmItem()
}
func parseArtFractionSubs(p *parser, fracType item) (f []subfraction) {
	for isFracSub(p.item, fracType) {
		f = append(f, parseArtFractionSub(p))
	}
	return
}
func parseArtFractionSub(p *parser) (f subfraction) {
	f.Num = p.item.val
	p.peekItem()
	p.confirmItem()
	for ; isArtFractionBody(p.item); p.peekItem() {
	}
	f.Body = p.confirmItem()
	return
}
func parseArtFooters(p *parser) (footers []subarticle) {
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
	return !isItemEOF(t)
}
func isItemContainer(t item) bool {
	return t.typ >= itemContainer
}
func isSameOrder(t1, t2 item) bool {
	return t1.typ == t2.typ && (t1.typ >= itemContainer)
}

func isArticle(t item) bool {
	return t.typ == itemArticle
}

func isItemSubArticle(t item) bool {
	return t.typ == itemSubArticleLet || t.typ == itemSubArticleNum
}

func isArtHeader(t item) bool {
	return !isItemEOF(t) && !isItemContainer(t) &&
		!isItemSubArticle(t)
}
func isArtFractionBody(t item) bool {
	return !isItemEOF(t) && !isItemContainer(t) &&
		!isItemSubArticle(t)
}
func isTitle(t item) bool {
	return !isItemEOF(t) && !isItemContainer(t)
}
func isDot(t item) bool {
	return t.val == "."
}
func isFracSub(t, fracType item) bool {
	return !isItemEOF(t) && !isItemContainer(t) && !(t.typ == fracType.typ)

}

func setParents(order item, title string, arts articles) {
	o := Reserved[order.val]
	v := order.val
	p := parent{Order: o, Num: v, Title: title}
	for i := 0; i < len(arts); i++ {
		arts[i].Parents = append(arts[i].Parents, p)
	}
}
