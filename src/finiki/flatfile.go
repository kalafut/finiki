package finiki

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

// GetPageRev returns the Page at path for a given rev. An error
// is returned if the page or requested rev is not found.
func (s *FlatFileStorage) GetPageRev(path Path, rev int) (*Page, error) {
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

// GetPage returns the most recent rev of the Page at path
func (s *FlatFileStorage) GetPage(path Path) (*Page, error) {
	pageInfo, err := s.NewPageInfo(path)
	if err != nil {
		return nil, err
	}

	return s.GetPageRev(path, pageInfo.CurrentRev)
}

func (s *FlatFileStorage) PutPage(path Path, page *Page) error {
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
func (s *FlatFileStorage) NewPageInfo(path Path) (*PageInfo, error) {
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

func (s *FlatFileStorage) DirList(path Path) PathList {
	list := PathList{}

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
			list = append(list, Path(fi.Name()+"/"))
		} else {
			list = append(list, Path(fi.Name()))
		}
	}

	return list
}

func revToFile(rev int) string {
	return fmt.Sprintf("%08d", rev)
}
