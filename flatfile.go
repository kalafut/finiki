package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const pageInfoFilename = "pageinfo"

type FlatFileStorage struct {
	root string
}

type PageInfo struct {
	filename   string
	CurrentRev int
}

func NewFlatFileStorage(root string) *FlatFileStorage {
	return &FlatFileStorage{
		root: root,
	}
}

// GetPage returns the Page at path for a given rev. Use CurrentRev to request the latest version.
// An error is returned if the page or requested rev is not found.
func (s *FlatFileStorage) GetPage(path string, rev int) (*Page, error) {
	if rev == CurrentRev {
		pageInfo, err := s.NewPageInfo(path)
		if err != nil {
			return nil, err
		}
		rev = pageInfo.CurrentRev
	}

	revPath := filepath.Join(s.root, string(path), "revs", revToFile(rev))

	f, err := os.Open(revPath)

	if err != nil {
		return nil, ErrPageNotFound
	}

	page, err := DecodePage(f)
	if err == nil {
		return page, nil
	}

	return nil, ErrPageCorrupt
}

func (s *FlatFileStorage) PutPage(path string, page *Page) error {
	// Check whether a folder is already at this location
	folder := filepath.Join(s.root, string(path))
	pi := filepath.Join(folder, pageInfoFilename)

	f1, err := os.Open(folder)
	defer f1.Close()
	if err == nil {
		f2, err := os.Open(pi)
		defer f2.Close()
		if os.IsNotExist(err) {
			return ErrFolderExists
		}
	}

	pageInfo, err := s.NewPageInfo(path)
	if err != nil {
		return err
	}
	pageInfo.CurrentRev++
	pageInfo.Save()

	os.MkdirAll(filepath.Join(s.root, string(path), "revs"), 0755)

	fullpath := filepath.Join(s.root, string(path), "revs", revToFile(pageInfo.CurrentRev))

	f, err := os.Create(fullpath)
	defer f.Close()

	if err == nil {
		page.EncodeJSON(f)
	}

	return nil
}

// NewPageInfo reads or a PageInfo file
func (s *FlatFileStorage) NewPageInfo(path string) (*PageInfo, error) {
	var pInfo PageInfo

	pInfo.filename = filepath.Join(s.root, string(path), pageInfoFilename)
	pInfo.CurrentRev = -1

	f, err := os.Open(pInfo.filename)
	defer f.Close()

	if os.IsNotExist(err) {
		return &pInfo, pInfo.Save()
	} else if err != nil {
		return nil, errors.New("Can't open PageInfo file")
	}

	dec := json.NewDecoder(f)
	err = dec.Decode(&pInfo)

	if err != nil {
		return nil, ErrPageCorrupt
	} else {
		return &pInfo, nil
	}

	return &pInfo, nil
}

func (p *PageInfo) Save() error {
	os.MkdirAll(filepath.Join(filepath.Dir(p.filename)), 0755)

	f, err := os.Create(p.filename)
	defer f.Close()

	if err != nil {
		errors.New("Can't save PageInfo file")
	}

	enc := json.NewEncoder(f)
	enc.Encode(p)

	return nil
}

func (s *FlatFileStorage) DirList(path string) []string {
	list := []string{}

	filename := filepath.Join(s.root, string(path))
	f, err := os.Open(filename)
	defer f.Close()

	entries, err := f.Readdir(0)
	if err != nil {
		log.Fatal(err)
	}

	for _, fi := range entries {
		if !fi.IsDir() { // should always be false
			continue
		}
		pi := filepath.Join(s.root, string(path), fi.Name(), pageInfoFilename)

		f, err := os.Open(pi)
		defer f.Close()

		if os.IsNotExist(err) {
			list = append(list, fi.Name()+"/")
		} else {
			list = append(list, fi.Name())
		}
	}

	return list
}

func (s *FlatFileStorage) GetPageList(root string) []string {
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

func revToFile(rev int) string {
	return fmt.Sprintf("%08d", rev)
}
