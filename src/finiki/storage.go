package main

import (
	"os"
	"path"
	"time"
)

type Storage interface {
	GetPage(path string) Page
	//GetPageRev(path string, rev int) *Page
	//GetRevList(path string) []something...
	PutPage(path string, page Page) error
}

type DumbStorage struct {
	db map[string]Page
}

func NewDumbStorage() *DumbStorage {
	return &DumbStorage{
		db: make(map[string]Page),
	}
}

func (s *DumbStorage) GetPage(path string) Page {
	p, ok := s.db[path]
	if !ok {
		p = Page{Content: "This is some *test* **Markdown** for new page: `" + path + "`!"}
		s.db[path] = p
	}

	return p
}

func (s *DumbStorage) PutPage(path string, page Page) error {
	s.db[path] = page

	return nil
}

type FlatFileStorage struct {
	root string
}

func NewFlatFileStorage(root string) *FlatFileStorage {
	return &FlatFileStorage{
		root: root,
	}
}

func (s *FlatFileStorage) GetPage(reqpath string) Page {
	fullpath := path.Join(s.root, reqpath, currentPage)

	f, err := os.Open(fullpath)
	if err == nil {
		ptr, err := DecodePage(f)
		if err == nil {
			return *ptr
		}
	}

	return Page{Date: time.Now(), Content: "This is some *test* **Markdown** for new page: `" + reqpath + "`!"}
}

func (s *FlatFileStorage) PutPage(reqpath string, page Page) error {
	os.MkdirAll(path.Join(s.root, reqpath), 0755)

	fullpath := path.Join(s.root, reqpath, currentPage)

	f, err := os.Create(fullpath)
	defer f.Close()

	if err == nil {
		page.EncodeJSON(f)

	}

	return nil
}
