package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var config = readLocalCfg()

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	//storage := NewDumbStorage()
	storage := NewFlatFileStorage(config.DataLocation)
	w := NewWiki(storage)
	//f := NewFolder("sample")

	//printFolder(f, "")

	router := httprouter.New()
	router.GET("/show/*path", w.Show)
	router.GET("/edit/*path", w.Edit)
	router.POST("/edit/*path", w.Update)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func printFolder(f Folder, parent string) {
	for name, page := range f.pages {
		fmt.Printf("%s - %s\n", parent+name, (*page).String())
	}

	for name, folder := range f.folders {
		printFolder(folder, parent+name+"/")
	}
}
