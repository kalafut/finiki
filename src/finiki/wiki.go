package main

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
