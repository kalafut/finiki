package main

import (
	"log"
	"net/http"

	"finiki"
)

var config = finiki.ReadLocalCfg()
var siteConfig = finiki.ReadSiteCfg(config.DataLocation)

func main() {
	storage := finiki.NewFlatFileStorage(config.DataLocation)
	w := finiki.NewWiki(storage)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", w.Route)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
