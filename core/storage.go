package core

import (
	"errors"
	"strings"
)

const CurrentRev = -1

var (
	ErrPageNotFound = errors.New("Page Not Found")
	ErrRevNotFound  = errors.New("Rev Not Found")
	ErrPageCorrupt  = errors.New("Page Corrupt")
	ErrFolderExists = errors.New("Folder already exists")
)

/*
Storage is the interface that includes all methods a wiki storage
implementation must provide.

GetPage returns a Page for the requested Path, at the given rev. If rev
is CurrentRev the latest version is returned.

PutPage saves a Path at the requested Path. If the Path does not exist or
is currently a page, it is created/updated. If the Path is currently a folder
ErrExists is returned.


*/
type Storage interface {
	GetPage(path Path, rev int) (*Page, error)
	PutPage(path Path, page *Page) error
	DirList(path Path) PathList
	GetPageList(root string) []string
}

/*
Path is valid wiki path. Paths conform to the following spec:

	1. Begins with "/"
	2. Has 0 or more intermediate, slash-delimited folders
	3. Ends with a page or folder name

	/page
	/a/b/c/page
	/a/c/d/folder

*/
type Path string

func (p Path) IsDir() bool {
	return strings.HasSuffix(string(p), "/")
}

type PathList []Path

// Sort support
func (d PathList) Len() int {
	return len(d)
}

func (d PathList) Less(i, j int) bool {
	if d[i].IsDir() == d[j].IsDir() {
		return d[i] < d[j]
	} else {
		return d[i].IsDir()
	}
}

func (d PathList) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
