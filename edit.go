package main

import (
	"net/http"
	"text/template"
)

// Edit is the edit endpoint of the Wiki
func (wiki *Wiki) Edit(w http.ResponseWriter, r *http.Request) {
	//path := ps.ByName("path")
	path := r.URL.Path
	page, err := wiki.store.GetPage(path, CurrentRev)

	if err != nil {
		page = &Page{Content: "Nothin'"}
	}

	vars := map[string]string{
		"Path": path,
		"Text": page.Content,
	}

	tmpl := make(map[string]*template.Template)
	tmpl["edit.html"] = template.Must(template.ParseFiles("templates/edit.html", "templates/base.html"))
	tmpl["edit.html"].ExecuteTemplate(w, "base", vars)

}
