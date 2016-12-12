package main

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const DEFAULT_EXT = ".md"

var (
	ErrPageNotFound = errors.New("Page Not Found")
	ErrRevNotFound  = errors.New("Rev Not Found")
	ErrPageCorrupt  = errors.New("Page Corrupt")
	ErrFolderExists = errors.New("Folder already exists")
)

var enableDefaultExt = true

type SimpleFileStorage struct {
	root string
}

func NewSimpleFileStorage(root string) SimpleFileStorage {
	return SimpleFileStorage{
		root: root,
	}
}

// GetPage returns the Page at path for a given rev. Use CurrentRev to request the latest version.
// An error is returned if the page or requested rev is not found.
func (s *SimpleFileStorage) GetPage(path string) (string, error) {
	fullPath := filepath.Join(s.root, path)

	page, err := ioutil.ReadFile(appendExt(fullPath))
	if err != nil {
		page, err = ioutil.ReadFile(fullPath)
		if err != nil {
			return "", ErrPageNotFound
		}
	}

	return string(page), nil
}

func (s *SimpleFileStorage) PutPage(path string, page string) error {
	os.MkdirAll(filepath.Join(s.root, filepath.Dir(path)), 0755)

	fullpath := filepath.Join(s.root, path)

	f, err := os.Create(appendExt(fullpath))
	defer f.Close()

	if err == nil {
		f.WriteString(page)
	} else {
		return ErrFolderExists
	}

	return nil
}

func (s *SimpleFileStorage) DeletePage(path string) error {
	err := os.Remove(filepath.Join(s.root, appendExt(path)))
	if err != nil {
		panic(err)
	}
	return err
}

func (s *SimpleFileStorage) DirList(path string) []string {
	root := filepath.Join(s.root, path)
	return fList(root, true)
}

func (s *SimpleFileStorage) GetPageList(path string) []string {
	pages := make([]string, 0)

	root := filepath.Join(s.root, path)
	for _, page := range fList(root, false) {
		pages = append(pages, strings.TrimSuffix(page, DEFAULT_EXT))
	}
	return pages
}

func appendExt(path string) string {
	if enableDefaultExt && !strings.HasSuffix(path, DEFAULT_EXT) {
		path = path + DEFAULT_EXT
	}
	return path
}

func fList(root string, dir bool) []string {
	list := []string{}

	entries, err := ioutil.ReadDir(root)

	if err == nil {
		for _, fi := range entries {
			if dir != fi.IsDir() || strings.HasPrefix(fi.Name(), ".") ||
				(fi.IsDir() && strings.HasPrefix(fi.Name(), "__")) {
				continue
			}
			list = append(list, fi.Name())
		}
	}

	return list
}
