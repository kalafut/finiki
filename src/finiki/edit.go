package main

import (
	"net/http"
	"text/template"
)

// Edit is the edit endpoint of the Wiki
func (wiki *Wiki) Edit(w http.ResponseWriter, r *http.Request) {
	//path := ps.ByName("path")
	path := r.URL.Path
	page, err := wiki.store.GetPage(path)

	if err != nil {
		page = &Page{Content: "Nothin'"}
	}

	t, err := template.New("edit").Parse(loadTemplate("edit"))

	template.Must(t, err).Execute(w, map[string]string{
		"Path": path,
		"Text": page.Content,
	})
}
