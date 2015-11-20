package main

import "errors"

var (
	ErrPageNotFound = errors.New("Page Not Found")
	ErrPageCorrupt  = errors.New("Page Corrupt")
)

type Storage interface {
	GetPage(path string) (*Page, error)
	GetPageRev(path string, rev int) (*Page, error)
	//GetRevList(path string) []something...
	PutPage(path string, page *Page) error
}
