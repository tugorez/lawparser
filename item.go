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
	itemSubArticleLet //1.-

	itemContainer
	itemBook
	itemTitle
	itemChapter
	itemSection
	itemArticle
	itemFraction
)

var itemString = map[itemType]string{
	itemError:         "error",
	itemEOF:           "eof",
	itemWord:          "word",
	itemReserved:      "reserved",
	itemRoman:         "roman",
	itemOrdinal:       "ordinal",
	itemLatin:         "latin",
	itemNumber:        "number",
	itemNewLine:       "newline",
	itemSeparator:     "separator",
	itemSymbol:        "symbol",
	itemSubArticleNum: "subarticlenum",
	itemSubArticleLet: "subarticlelet",
	itemContainer:     "container",
	itemBook:          "book",
	itemTitle:         "title",
	itemChapter:       "chapter",
	itemSection:       "section",
	itemArticle:       "article",
	itemFraction:      "fraction",
}

func (self item) String() string {
	v, _ := itemString[self.typ]
	return fmt.Sprintf("(%s,%s)", v, self.val)
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
