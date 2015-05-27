package lawparser

import (
	"fmt"
)

type item struct {
	typ itemType
	val string
}

type itemType int

const (
	itemError itemType = iota //error occurred
	itemEOF
	itemWord     //any word
	itemReserved // any of these words[book|title|chapter|section|article|fraction]

	itemRoman //roman numbers
	itemOrdinal
	itemLatin
	itemNumber

	itemNewLine
	itemSeparator

	itemSymbol

	itemSubArticleNum //a)
	itemsubArticleLet //1.-

	itemContainer
	itemBook
	itemTitle
	itemChapter
	itemSection
	itemArticle
	itemFraction
)

func (self item) String() string {
	switch {
	case self.typ == itemError:
		return self.val
	case self.typ == itemEOF:
		return "EOF"
	case self.typ > itemWord:
		return fmt.Sprintf("<%s>", self.val)
	default:
		return fmt.Sprintf("%s", self.val)
	}
}

//These are "all" posible ways to write  the diferent token containers
//This is a kind of mapping from diferent ways to write it to just one
var Reserved = map[string]string{
	"libro":  containerBook,
	"libros": containerBook,

	"título":  containerTitle,
	"titulo":  containerTitle,
	"títulos": containerTitle,
	"titulos": containerTitle,

	"capítulo":  containerChapter,
	"capitulo":  containerChapter,
	"capitulos": containerChapter,

	"sección":   containerSection,
	"seccion":   containerSection,
	"secciones": containerSection,

	"artículo":  containerArticle,
	"articulo":  containerArticle,
	"artículos": containerArticle,
	"articulos": containerArticle,

	"fracción":   containerFraction,
	"fraccion":   containerFraction,
	"fracciones": containerFraction,
}

const (
	containerBook     = "libro"
	containerTitle    = "título"
	containerChapter  = "capítulo"
	containerSection  = "sección"
	containerArticle  = "artículo"
	containerFraction = "fracción"
)

//We map the reserved item  to its itemType as container
var containerType = map[string]itemType{
	containerBook:     itemBook,
	containerTitle:    itemTitle,
	containerChapter:  itemChapter,
	containerSection:  itemSection,
	containerArticle:  itemArticle,
	containerFraction: itemFraction,
}
