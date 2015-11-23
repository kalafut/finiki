package main

import (
	"log"
	"net/http"
)

var config = readLocalCfg()
var siteConfig = readSiteCfg()

func main() {
	storage := NewFlatFileStorage(config.DataLocation)
	w := NewWiki(storage)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", w.Route)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (wiki *Wiki) Route(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	action := queryValues.Get("action")

	switch {
	case r.PostFormValue("update") == "update":
		wiki.Update(w, r)
	case action == "edit":
		wiki.Edit(w, r)
	default:
		wiki.Show(w, r)
	}
}
