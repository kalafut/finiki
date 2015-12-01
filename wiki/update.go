package wiki

import (
	"net/http"

	"github.com/kalafut/finiki/core"
)

// Edit is the edit endpoint of the Wiki
func (wiki *Wiki) Update(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	r.ParseForm()

	page, err := wiki.store.GetPage(core.Path(path), core.CurrentRev)

	if err != nil {
		page = &core.Page{}
	}

	newText := r.PostFormValue("text")
	page.SetContent(newText)

	wiki.store.PutPage(core.Path(path), page)

	http.Redirect(w, r, path, 302)
}
