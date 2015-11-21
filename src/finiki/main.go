package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var config = readLocalCfg()
var siteConfig = readSiteCfg()

func main() {
	storage := NewFlatFileStorage(config.DataLocation)
	w := NewWiki(storage)

	router := httprouter.New()
	router.GET("/show/*path", w.Show)
	router.GET("/edit/*path", w.Edit)
	router.POST("/edit/*path", w.Update)

	log.Fatal(http.ListenAndServe(":8080", router))
}
