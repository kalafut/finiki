package main

import (
	"testing"

	"gopkg.in/tylerb/is.v1"
)

/*
func TestPageString(t *testing.T) {
	is := is.New(t)

	page := Page{
		Title:   "some title",
		Date:    time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		Content: "This is some content",
	}

	var b bytes.Buffer
	page.Encode(&b)

	is.Equal(`title: some title
date: 2009-11-10 23:00 UTC
tags: tbd
---
This is some content`, b.String())

}

func TestNewPage(t *testing.T) {
	is := is.New(t)

	p1 := Page{
		Title:   "some title",
		Date:    time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		Content: "This is some content",
	}

	var b bytes.Buffer
	p1.Encode(&b)

	p2, err := NewPage(&b)
	is.NotErr(err)
	is.Equal(p1, *p2)
}
*/

func TestParseString(t *testing.T) {
	is := is.New(t)

	fld, val, err := parseField("title: This is super!")
	is.NotErr(err)
	is.Equal(fld, "title")
	is.Equal(val, "This is super!")

	_, _, err = parseField("This is not super!")
	is.Err(err)

	_, _, err = parseField(":This is not super!")
	is.Err(err)
}
