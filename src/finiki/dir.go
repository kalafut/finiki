package main

import (
	"html/template"
	"net/http"
	"sort"
)

func (wiki *Wiki) Dir(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	paths := wiki.store.DirList(Path(path))

	// TODO put this back when Path type is sorted
	sort.Sort(paths)

	vars := map[string]interface{}{
		"Path":  path + "?action=edit",
		"Paths": paths,
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
