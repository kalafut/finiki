package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func (wiki *Wiki) Dir(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	queryValues := r.URL.Query()
	rev := queryValues.Get("rev")

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
		"Path": path + "?action=edit",
		"Text": BytesAsHTML(ParsedMarkdown(parsedContent)),
	}

	tmpl := make(map[string]*template.Template)
	tmpl["dir.html"] = template.Must(template.ParseFiles("templates/dir.html", "templates/base.html"))
	tmpl["dir.html"].ExecuteTemplate(w, "base", vars)
}

//// RedirectToShow redirects to the show endpoint using a HTTP 302
///*
//func (w *Wiki) RedirectToShow(c web.C, rw http.ResponseWriter, r *http.Request) {
//	http.Redirect(rw, r, "/"+c.URLParams["name"], 302)
//}
//*/
//
//// BytesAsHTML returns the template bytes as HTML
//func BytesAsHTML(b []byte) template.HTML {
//	return template.HTML(string(b))
//}
//
//// ParsedMarkdown returns provided bytes parsed as Markdown
//func ParsedMarkdown(b string) []byte {
//	return blackfriday.MarkdownCommon([]byte(b))
//}
//
//var reLink = regexp.MustCompile(`\[\[(.*?)\]\]`)
//
//func preParse(in string) string {
//	return reLink.ReplaceAllString(in, "[$1]($1)")
//}
