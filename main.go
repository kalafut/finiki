package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = "8080"

var config = ReadLocalCfg()
var siteConfig = ReadSiteCfg(config.DataLocation)

func main() {
	storage := NewSimpleFileStorage(config.DataLocation)
	w := NewWiki(storage)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", w.Route)

	fmt.Println("Starting finiki server on port " + PORT + "...")
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
