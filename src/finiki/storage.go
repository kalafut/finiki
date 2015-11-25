package main

import (
	"errors"
	"strings"
)

var (
	ErrPageNotFound = errors.New("Page Not Found")
	ErrPageCorrupt  = errors.New("Page Corrupt")
)

type DirEntry string
type DirEntries []DirEntry

func (d DirEntries) Len() int {
	return len(d)
}

func (d DirEntries) Less(i, j int) bool {
	if d[i].IsFolder() == d[j].IsFolder() {
		return d[i] < d[j]
	} else {
		return d[i].IsFolder()
	}
}

func (d DirEntries) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (s DirEntry) IsFolder() bool {
	return strings.HasSuffix(string(s), "/")
}

type Storage interface {
	GetPage(path string) (*Page, error)
	GetPageRev(path string, rev int) (*Page, error)
	//GetRevList(path string) []something...
	PutPage(path string, page *Page) error
	DirList(path string) []string
}
