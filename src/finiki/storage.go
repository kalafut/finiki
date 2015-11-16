package main

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
