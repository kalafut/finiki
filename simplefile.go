package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

//const pageInfoFilename = "pageinfo"

type SimpleFileStorage struct {
	root string
}

//type PageInfo struct {
//	filename   string
//	CurrentRev int
//}

func NewSimpleFileStorage(root string) *SimpleFileStorage {
	return &SimpleFileStorage{
		root: root,
	}
}

// GetPage returns the Page at path for a given rev. Use CurrentRev to request the latest version.
// An error is returned if the page or requested rev is not found.
func (s *SimpleFileStorage) GetPage(path string, rev int) (*Page, error) {
	revPath := filepath.Join(s.root, path)

	page, err := ioutil.ReadFile(revPath)
	if err == nil {
		p := Page{Content: string(page)}
		return &p, nil
	}

	return nil, ErrPageNotFound
}

func (s *SimpleFileStorage) PutPage(path string, page *Page) error {
	os.MkdirAll(filepath.Join(s.root, filepath.Dir(path)), 0755)

	fullpath := filepath.Join(s.root, path)

	f, err := os.Create(fullpath)
	defer f.Close()

	if err == nil {
		f.WriteString(page.Content)
	}

	return nil
}

func (s *SimpleFileStorage) DirList(path string) []string {
	list := []string{}

	root := filepath.Join(s.root, path)
	entries, err := ioutil.ReadDir(root)

	if err != nil {
		log.Fatal(err)
	}

	for _, fi := range entries {
		if !fi.IsDir() {
			continue
		}

		list = append(list, fi.Name()+"/")
	}

	return list
}

func (s *SimpleFileStorage) GetPageList(root string) []string {
	fmt.Println(root)
	pages := []string{}

	filepath.Walk(s.root, func(path string, info os.FileInfo, err error) error {
		if filepath.Base(path) == pageInfoFilename {
			pages = append(pages, filepath.ToSlash(filepath.Dir(path))[len(s.root):])
			return filepath.SkipDir
		}
		return nil
	})

	return pages
}
