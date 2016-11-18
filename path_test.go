package main

import (
	"testing"

	"gopkg.in/tylerb/is.v1"
)

func TestPathSplit(t *testing.T) {
	is := is.New(t)

	tests := []struct {
		path string
		dir  string
		page string
	}{
		{"/", "/", ""},
		{"/dog", "/", "dog"},
		{"/dog/", "/dog/", ""},
		{"/sub/cat", "/sub/", "cat"},
		{"/sub/cat/bird/", "/sub/cat/bird/", ""},
		{"/sub/cat/bird/bee", "/sub/cat/bird/", "bee"},
		{"", "", ""},
		{"bad", "", ""},
		{" /bad", "", ""},
		{"/bad//page", "", ""},
		{"/ok95", "/", "ok95"},
		{"/blah badok95", "", ""},
	}

	for _, test := range tests {
		dir, page := PathSplit(test.path)
		is.Equal(test.dir, dir)
		is.Equal(test.page, page)
	}
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
		dirs, pages := PagelistProc(test.root, test.pages)
		is.Equal(test.expDirs, dirs)
		is.Equal(test.expPages, pages)
	}
}
