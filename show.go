package main

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const RECENT_CNT = 8

// Show is the show endpoint of the Wiki
func (wiki *Wiki) Show(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	rev := queryValues.Get("rev")

	//path := ps.ByName("path")
	path := r.URL.Path

	var page *Page
	var err error
	rev_rqst := CurrentRev

	if r, e := strconv.Atoi(rev); e == nil {
		rev_rqst = r
	}
	page, err = wiki.store.GetPage(path, rev_rqst)

	if err != nil {
		page = &Page{Content: "Nothin'"}
	} else {
		saveRecent(path, wiki.store)
	}

	parsedContent := preParse(page.Content)
	vars := map[string]interface{}{
		"Path":        path + "?action=edit",
		"Text":        BytesAsHTML(ParsedMarkdown(parsedContent)),
		"Title":       path[1:],
		"RecentPaths": loadRecent(wiki.store),
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

func saveRecent(path string, s Storage) {
	page, err := s.GetPage("__system/recent", 0)
	if err == ErrPageNotFound {
		page = &Page{}
	}
	lines := strings.Split(page.Content, "\n")
	lines = append([]string{path}, lines...)
	dedupe := make([]string, 0)
	uniqLines := make(map[string]bool)

	for _, line := range lines {
		if _, ok := uniqLines[line]; !ok {
			dedupe = append(dedupe, line)
			uniqLines[line] = true
		}
	}

	page.Content = strings.Join(dedupe, "\n")
	s.PutPage("__system/recent", page)
}

func loadRecent(s Storage) []string {
	var text string

	recents, err := s.GetPage("__system/recent", 0)
	if err == nil {
		text = recents.Content
	}

	list := strings.Split(strings.TrimSpace(text), "\n")

	// skip the first entry since it will be the page we're on
	return list[Min(1, len(list)):Min(len(list), RECENT_CNT+1)]
}
