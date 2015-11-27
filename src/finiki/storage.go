package main

import "errors"

var (
	ErrPageNotFound = errors.New("Page Not Found")
	ErrPageCorrupt  = errors.New("Page Corrupt")
)

type Storage interface {
	GetPage(path Path) (*Page, error)
	GetPageRev(path Path, rev int) (*Page, error)
	//GetRevList(path string) []something...
	PutPage(path Path, page *Page) error
	DirList(path Path) PathList
}
