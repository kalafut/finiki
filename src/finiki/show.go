package main

import (
	"html/template"
	"net/http"
	"regexp"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/russross/blackfriday"
)

// Show is the show endpoint of the Wiki
func (wiki *Wiki) Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	queryValues := r.URL.Query()
	rev := queryValues.Get("rev")

	path := ps.ByName("path")

	var page *Page
	var err error

	if r, e := strconv.Atoi(rev); e == nil {
		page, err = wiki.store.GetPageRev(path, r)
	} else {
		page, err = wiki.store.GetPage(path)
	}

	if err != nil {
		page = &Page{Content: "Nothin'"}
	}

	parsedContent := preParse(page.Content)
	vars := map[string]interface{}{
		"Path": "/edit" + string(path),
		"Text": BytesAsHTML(ParsedMarkdown(parsedContent)),
	}

	t := template.Must(template.New("show").Parse(loadTemplate("show")))
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

var reLink = regexp.MustCompile(`\[\[(.*?)\]\]`)

func preParse(in string) string {
	return reLink.ReplaceAllString(in, "[$1]($1)")
}
