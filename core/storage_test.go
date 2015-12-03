// Tests for validating an implementation of Storage.

package core

import (
	"testing"

	"gopkg.in/tylerb/is.v1"
)

func testStorage(t *testing.T, s Storage) {
	testGetPage(t, s)
	testPutPage(t, s)
}

func testGetPage(t *testing.T, s Storage) {
	var pg Page
	var err error
	is := is.New(t)

	// Test page not found
	page, err := s.GetPage("/", CurrentRev)
	is.Nil(page)
	is.Equal(err, ErrPageNotFound)

	// Test putting and getting a page at the root level
	pg = NewPage()
	pg.Content = "test"

	s.PutPage("/t1", &pg)
	pg2, err := s.GetPage("/t1", CurrentRev)
	is.NotErr(err)

	is.Equal(pg.Content, pg2.Content)

	// Test putting and getting a page at a deep level
	pg = NewPage()
	pg.Content = "test2"

	s.PutPage("/a/b/c/d/e/t2", &pg)
	pg2, err = s.GetPage("/a/b/c/d/t2", CurrentRev)
	is.Nil(page)
	is.Equal(err, ErrPageNotFound)

	pg2, err = s.GetPage("/a/b/c/d/e/t2", CurrentRev)
	is.NotErr(err)

	is.Equal(pg.Content, pg2.Content)
}

func testPutPage(t *testing.T, s Storage) {
	var pg Page
	var err error
	is := is.New(t)

	// Test putting a page atop an existing folder
	pg = NewPage()
	pg.Content = "test2"

	err = s.PutPage("/putpage/a/test", &pg)
	is.NotErr(err)

	err = s.PutPage("/putpage", &pg)
	is.Equal(err, ErrFolderExists)

	err = s.PutPage("/putpage/a", &pg)
	is.Equal(err, ErrFolderExists)

	err = s.PutPage("/putpage/a/test", &pg)
	is.NotErr(err)
}

func TestPagelistProc(t *testing.T) {
	is := is.New(t)

	list1 := []string{
		"/a",
		"/b",
		"/d1/a",
		"/d1/b",
		"/d1/s1/c",
		"/d1/s1/ss1/d",
		"/d1/s1/ss1/e",
		"/d2/f",
	}

	tests := []struct {
		root     string
		pages    []string
		expDirs  []string
		expPages []string
	}{
		{
			"/",
			list1,
			[]string{
				"/d1",
				"/d2",
			},
			[]string{
				"/a",
				"/b",
			},
		},
	}

	for _, test := range tests {
		dirs, pages := pagelistProc(test.root, test.pages)
		is.Equal(test.expDirs, dirs)
		is.Equal(test.expPages, pages)
	}
}
