package main

import (
	"io/ioutil"
	"os"
	"testing"

	"gopkg.in/tylerb/is.v1"
)

func TestInterface(t *testing.T) {
	dir, _ := ioutil.TempDir("", "")
	s := NewFlatFileStorage(dir)
	testStorage(t, s)
	os.RemoveAll(dir)

	dir, _ = ioutil.TempDir("", "")
	s2 := NewSimpleFileStorage(dir)
	testStorage(t, s2)
	os.RemoveAll(dir)

}

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
