package main

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// Edit is the edit endpoint of the Wiki
func (wiki *Wiki) Edit(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	path := ps.ByName("path")
	page, err := wiki.store.GetPage(path)

	if err != nil {
		page = &Page{Content: "Nothin'"}
	}

	t, err := template.New("edit").Parse(loadTemplate("edit"))

	template.Must(t, err).Execute(w, map[string]string{
		"Path": "/edit" + path,
		"Text": page.Content,
	})
}
