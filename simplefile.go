package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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
	} else {
		return ErrFolderExists
	}

	return nil
}

func (s *SimpleFileStorage) DirList(path string) []string {
	root := filepath.Join(s.root, path)
	return fList(root, true)
}

func (s *SimpleFileStorage) GetPageList(path string) []string {
	root := filepath.Join(s.root, path)
	return fList(root, false)
}

func fList(root string, dir bool) []string {
	list := []string{}

	entries, err := ioutil.ReadDir(root)

	if err == nil {
		for _, fi := range entries {
			if (dir != fi.IsDir()) || strings.HasPrefix(fi.Name(), ".") {
				continue
			}

			list = append(list, fi.Name())
		}
	}

	return list

}
