package main

import (
	"net/http"
	"strings"
)

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
	case strings.HasSuffix(path, "/"):
		wiki.Dir(w, r)
	default:
		wiki.Show(w, r)
	}
}
