package main

import (
	"fmt"
	"log"
	"net/http"
)

var config = ReadLocalCfg()
var siteConfig = ReadSiteCfg(config.DataLocation)

func main() {
	storage := NewFlatFileStorage(config.DataLocation)
	w := NewWiki(storage)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", w.Route)

	fmt.Println("Starting finiki server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
