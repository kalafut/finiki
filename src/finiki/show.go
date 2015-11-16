package main

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/russross/blackfriday"
)

// Show is the show endpoint of the Wiki
func (wiki *Wiki) Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	path := ps.ByName("path")
	page := wiki.store.GetPage(path)

	vars := map[string]interface{}{
		"Path": "/edit" + string(path),
		"Text": BytesAsHTML(ParsedMarkdown(page.Content)),
	}

	t := template.Must(template.New("show").Parse(showTpl))
	t.Execute(w, vars)
}

// RedirectToShow redirects to the show endpoint using a HTTP 302
/*
func (w *Wiki) RedirectToShow(c web.C, rw http.ResponseWriter, r *http.Request) {
	http.Redirect(rw, r, "/"+c.URLParams["name"], 302)
}
*/

// BytesAsHTML returns the template bytes as HTML
func BytesAsHTML(b []byte) template.HTML {
	return template.HTML(string(b))
}

// ParsedMarkdown returns provided bytes parsed as Markdown
func ParsedMarkdown(b string) []byte {
	return blackfriday.MarkdownCommon([]byte(b))
}
