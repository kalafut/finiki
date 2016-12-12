package main

import (
	"html/template"
	"path/filepath"
	"strings"
)

var templates = initTemplates()

var stdFuncMap = template.FuncMap{
	"basename": basename,
	"dec": func(i int) int {
		return i - 1
	},
}

func basename(path string) string {
	return strings.TrimSuffix(filepath.Base(path), DEFAULT_EXT)
}

func initTemplates() map[string]*template.Template {
	tmpl := make(map[string]*template.Template)
	tmplMap := map[string][]string{
		"show.html":   {"recent.html", "header.html", "show.html", "base.html"},
		"dir.html":    {"recent.html", "header.html", "dir.html", "base.html"},
		"delete.html": {"recent.html", "header.html", "delete.html", "base.html"},
	}

	for name, files := range tmplMap {
		var fns []string
		for _, file := range files {
			fns = append(fns, "templates/"+file)
		}
		tmpl[name] = template.Must(template.New("").Funcs(stdFuncMap).ParseFiles(fns...))
	}

	return tmpl
}
