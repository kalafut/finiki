package main

import (
	"bufio"
	"errors"
	"io"
	"io/ioutil"
	"strings"
	"time"
)

const dateFmt = time.RFC822

type Folder struct {
	folders map[string]Folder
	pages   map[string]Page
}

type Page struct {
	Title   string
	Date    time.Time
	Tags    []string
	Content string
}

type PageEncoder interface {
	Encode(page *Page)
}

type PageDecoder interface {
	Decode() (*Page, error)
}

type PageStoreV1 struct {
	w io.Writer
	r io.Reader
}

func NewPageStoreV1(w io.Writer, r io.Reader) PageStoreV1 {
	return PageStoreV1{w: w, r: r}
}

func NewPage(input io.Reader) (*Page, error) {
	var p Page

	r := bufio.NewReader(input)
	for {
		line, err := r.ReadString('\n')
		if strings.HasPrefix(line, "---") {
			break
		}

		field, val, err := parseField(line)
		if err != nil {
			return nil, err
		}
		switch field {
		case "title":
			p.Title = val
		case "date":
			p.Date, err = time.Parse(dateFmt, val)
			if err != nil {
				return nil, errors.New("Invalid date")
			}
		}
	}

	b, _ := ioutil.ReadAll(r)
	p.Content = string(b)

	return &p, nil
}

func (p *Page) Encode(w io.Writer) {
	var buffer = bufio.NewWriter(w)

	buffer.WriteString("title: " + p.Title + "\n")
	buffer.WriteString("date: " + p.Date.Format(dateFmt) + "\n")
	buffer.WriteString("tags: tbd\n")
	buffer.WriteString("---\n")
	buffer.WriteString(p.Content)

	buffer.Flush()
}

func parseField(s string) (string, string, error) {
	idx := strings.Index(s, ":")

	if idx < 1 {
		return "", "", errors.New("Invalid header line")
	}

	field := strings.TrimSpace(s[0:idx])
	val := strings.TrimSpace(s[idx+1:])

	return field, val, nil
}
