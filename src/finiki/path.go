package main

import "strings"

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
