package htmlparser

import (
	"github.com/mkboudreau/reservations/htmlparser/internal"
)



func ParseHtml(html string) DocumentNode {

	engine := internal.NewHtmlParsingEngine("<html></html>")
	document := engine.Parse()
	return document

}


