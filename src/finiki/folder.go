package main

import (
	"log"
	"os"
	"path/filepath"
)

const currentPage = "current"

/*
A Folder can contain either Folders or PageFolders
*/
type Folder struct {
	folders map[string]Folder
	pages   map[string]*Page
}

var root Folder

var join = filepath.Join

func NewFolder(path string) Folder {
	var f *os.File
	var err error

	folder := Folder{folders: make(map[string]Folder), pages: make(map[string]*Page)}

	//base := filepath.Base(path)

	// See if we're in a page folder or not
	cf := join(path, currentPage)
	if _, err := os.Stat(cf); err != nil {
		//page, err := readPageFolder(path)
	} else {
		//f2 := Folder{}
		//		f2.Load(cf)
		//f.folders[
	}

	// Traverse

	if f, err = os.Open(path); err != nil {
		log.Fatal(err)
	}

	entries, err := f.Readdirnames(0)
	if err != nil {
		log.Fatal(err)
	}

	for _, name := range entries {
		fullPath := join(path, name)

		if isPageFolder(fullPath) {
			p, err := readPageFolder(fullPath)
			if err != nil {
				log.Fatal(err)
			}
			folder.pages[name] = p
		} else {
			fi, err := os.Stat(fullPath)
			if err == nil && fi.IsDir() {
				// recurse!
				folder.folders[name] = NewFolder(fullPath)
			}
		}
	}

	_ = entries
	return folder
}

func readPageFolder(path string) (*Page, error) {
	f, err := os.Open(join(path, currentPage))
	defer f.Close()

	if err != nil {
		return nil, err
	}

	return DecodePage(f)
}

func isPageFolder(path string) bool {
	_, err := os.Stat(join(path, currentPage))

	return err == nil
}
