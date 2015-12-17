package wiki

import (
	"html/template"
	"net/http"
	"regexp"
	"strconv"

	"github.com/kalafut/finiki/core"
)

// Show is the show endpoint of the Wiki
func (wiki *Wiki) Show(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	rev := queryValues.Get("rev")

	//path := ps.ByName("path")
	path := r.URL.Path

	var page *core.Page
	var err error
	rev_rqst := core.CurrentRev

	if r, e := strconv.Atoi(rev); e == nil {
		rev_rqst = r
	}
	page, err = wiki.store.GetPage(path, rev_rqst)

	if err != nil {
		page = &core.Page{Content: "Nothin'"}
	}

	parsedContent := preParse(page.Content)
	vars := map[string]interface{}{
		"Path":  path + "?action=edit",
		"Text":  BytesAsHTML(ParsedMarkdown(parsedContent)),
		"Title": path[1:],
	}

	tmpl := make(map[string]*template.Template)
	tmpl["show.html"] = template.Must(template.ParseFiles("templates/header.html", "templates/show.html", "templates/base.html"))
	tmpl["show.html"].ExecuteTemplate(w, "base", vars)
}

// RedirectToShow redirects to the show endpoint using a HTTP 302
/*
func (w *Wiki) RedirectToShow(c web.C, rw http.ResponseWriter, r *http.Request) {
	http.Redirect(rw, r, "/"+c.URLParams["name"], 302)
}
*/

var reLink = regexp.MustCompile(`\[\[(.*?)\]\]`)

func preParse(in string) string {
	return reLink.ReplaceAllString(in, "[$1]($1)")
}
