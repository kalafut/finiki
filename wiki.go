package main

import "net/http"

// Wiki represents the entire Wiki, contains the db
type Wiki struct {
	store Storage
}

// NewWiki creates a new Wiki
func NewWiki(s Storage) *Wiki {
	// Setup the wiki.
	w := &Wiki{store: s}

	return w
}

// DB returns the database associated with the handler.
func (w *Wiki) Store() Storage {
	return w.store
}

func (wiki Wiki) Route(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	path := r.URL.Path
	action := queryValues.Get("action")

	switch {
	case r.PostFormValue("update") == "update":
		wiki.Update(w, r)
	case action == "edit":
		wiki.Edit(w, r)
	default:
		if len(wiki.store.GetPageList(path)) > 0 || len(wiki.store.DirList(path)) > 0 {
			wiki.Dir(w, r)
		} else {
			wiki.Show(w, r)
		}
	}
}
