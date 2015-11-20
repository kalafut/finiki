package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Edit is the edit endpoint of the Wiki
func (wiki *Wiki) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	path := ps.ByName("path")
	r.ParseForm()

	page, err := wiki.store.GetPage(path)

	if err != nil {
		page = &Page{}
	}

	newText := r.PostFormValue("text")
	page.SetContent(newText)

	wiki.store.PutPage(path, page)

	http.Redirect(w, r, "/show"+path, 302)
}
