package main

import (
	"net/http"
	"regexp"
	"strings"
)

const RECENT_CNT = 8

// Wiki represents the entire Wiki, contains the db
type Wiki struct {
	store SimpleFileStorage
}

// NewWiki creates a new Wiki
func NewWiki(s SimpleFileStorage) *Wiki {
	// Setup the wiki.
	w := &Wiki{store: s}

	return w
}

// DB returns the database associated with the handler.
func (w *Wiki) Store() SimpleFileStorage {
	return w.store
}

func (wiki Wiki) Route(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	path := r.URL.Path
	action := queryValues.Get("action")

	isDir := len(wiki.store.GetPageList(path)) > 0 || len(wiki.store.DirList(path)) > 0

	if isDir && !strings.HasSuffix(path, "/") {
		http.Redirect(w, r, path+"/", http.StatusSeeOther)
		return
	}

	switch {
	case r.PostFormValue("update") == "update":
		wiki.Update(w, r)
	case action == "edit":
		wiki.Edit(w, r)
	case isDir:
		wiki.Dir(w, r)
	default:
		wiki.Show(w, r)
	}
}

// Show is the show endpoint of the Wiki
func (wiki *Wiki) Show(w http.ResponseWriter, r *http.Request) {
	var page string
	var err error

	path := r.URL.Path

	page, err = wiki.store.GetPage(path)

	if err != nil {
		page = "Nothin'"
	} else {
		saveRecent(path, wiki.store)
	}

	parsedContent := preParse(page)
	vars := map[string]interface{}{
		"Path":        path + "?action=edit",
		"Text":        BytesAsHTML(ParsedMarkdown(parsedContent)),
		"Title":       path[1:],
		"RecentPaths": loadRecent(wiki.store, true),
	}

	templates["show.html"].ExecuteTemplate(w, "base", vars)
}

// RedirectToShow redirects to the show endpoint using a HTTP 302
/*
func (w *Wiki) RedirectToShow(c web.C, rw http.ResponseWriter, r *http.Request) {
	http.Redirect(rw, r, "/"+c.URLParams["name"], 302)
}
*/

var reLink = regexp.MustCompile(`\[\[(.*?)\]\]`)

func preParse(in string) string {
	return reLink.ReplaceAllString(in, "[$1]($1)")
}

func saveRecent(path string, s SimpleFileStorage) {
	page, _ := s.GetPage("__system/recent")

	lines := strings.Split(page, "\n")
	lines = append([]string{path}, lines...)
	dedupe := make([]string, 0)
	uniqLines := make(map[string]bool)

	for _, line := range lines {
		if _, ok := uniqLines[line]; !ok {
			dedupe = append(dedupe, line)
			uniqLines[line] = true
		}
	}

	page = strings.Join(dedupe, "\n")
	s.PutPage("__system/recent", page)
}

func loadRecent(s SimpleFileStorage, skipFirst bool) []string {
	var text string
	var start int

	recents, err := s.GetPage("__system/recent")
	if err == nil {
		text = recents
	}

	list := strings.Split(strings.TrimSpace(text), "\n")

	// skip the first entry since it will be the page we're on
	if skipFirst {
		start = 1
	}
	return list[Min(start, len(list)):Min(len(list), RECENT_CNT+1)]
}

func (wiki *Wiki) Update(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	r.ParseForm()

	page, _ := wiki.store.GetPage(path)
	newText := r.PostFormValue("text")
	if newText != page {
		wiki.store.PutPage(path, newText)
	}

	http.Redirect(w, r, path, 302)
}

func (wiki *Wiki) Dir(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	dirs := wiki.store.DirList(path)
	pages := wiki.store.GetPageList(path)

	// TODO put this back when Path type is sorted
	//sort.Sort(paths)

	vars := map[string]interface{}{
		"Path":        path + "?action=edit",
		"Dirs":        dirs,
		"Pages":       pages,
		"RecentPaths": loadRecent(wiki.store, false),
	}

	templates["dir.html"].ExecuteTemplate(w, "base", vars)
}
