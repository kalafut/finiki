package main

import (
	"log"
	"net/http"

	"github.com/kalafut/finiki/core"
	"github.com/kalafut/finiki/flatfile"
	"github.com/kalafut/finiki/wiki"
)

var config = core.ReadLocalCfg()
var siteConfig = core.ReadSiteCfg(config.DataLocation)

func main() {
	storage := flatfile.NewFlatFileStorage(config.DataLocation)
	w := wiki.NewWiki(storage)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", w.Route)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
