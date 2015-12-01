package core

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"log"
	"strings"
	"time"
)

const dateFmt = "2006-01-02 15:04 MST"

const (
	Markdown = iota
)

type Page struct {
	Title        string
	LastModified time.Time
	Tags         []string
	Encoding     int
	Content      string
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

func NewPage() Page {
	return Page{LastModified: time.Now()}
}

func DecodePage(input io.Reader) (*Page, error) {
	var p Page

	dec := json.NewDecoder(input)
	dec.Decode(&p)

	/*
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
	*/

	return &p, nil
}

func (p *Page) Encode(w io.Writer) {
	var buffer = bufio.NewWriter(w)

	buffer.WriteString("title: " + p.Title + "\n")
	buffer.WriteString("date: " + p.LastModified.Format(dateFmt) + "\n")
	buffer.WriteString("tags: tbd\n")
	buffer.WriteString("---\n")
	buffer.WriteString(p.Content)

	buffer.Flush()
}

func (p *Page) EncodeJSON(w io.Writer) {
	p.LastModified = p.LastModified.Round(time.Second)
	b, err := json.MarshalIndent(p, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(b)
}

func (p *Page) SetContent(s string) {
	p.Content = s
	p.LastModified = roundedNow()
}

func (p *Page) String() string {
	return p.Title
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

func roundedNow() time.Time {
	return time.Now().Round(time.Second)
}
