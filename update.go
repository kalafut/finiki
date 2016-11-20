package main

import (
	"net/http"
)

// Edit is the edit endpoint of the Wiki
func (wiki *Wiki) Update(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	r.ParseForm()

	page, err := wiki.store.GetPage(path, CurrentRev)

	if err != nil {
		page = &Page{}
	}

	newText := r.PostFormValue("text")
	if newText != page.Content {
		page.SetContent(newText)

		wiki.store.PutPage(path, page)
	}

	http.Redirect(w, r, path, 302)
}