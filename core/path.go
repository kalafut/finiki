package core

import (
	"regexp"
	"strings"
)

func PathSplit(path string) (dir, page string) {
	if !PathValid(path) {
		return
	}

	idx := strings.LastIndex(path, "/")

	dir = path[0 : idx+1]
	if idx < len(path) {
		page = path[idx+1:]
	}

	return
}

var pathRe = regexp.MustCompile(`^/([[:alnum:]]+/)*([[:alnum:]]+)?$`)

func PathValid(path string) bool {
	return pathRe.MatchString(path)
	/*
		valid := false

		switch {
		case len(path) == 0:
		case path[0] != '/':
			case
		default:
			valid = true
		}

		return valid
	*/
}

func pagelistProc(root string, pagelist []string) ([]string, []string) {
	dirs := []string{}
	pages := []string{}

	rootLen := len(strings.Split(root, "/")) - 1
	println(rootLen)

	for _, page := range pagelist {
		if !strings.HasPrefix(page, root) {
			continue
		}
		els := strings.Split(page, "/")[1:]
		//fmt.Printf("%v\n", els)

		if len(els) == rootLen {
			pages = append(pages, root+els[len(els)-1])
		} else {
			e := root + els[rootLen-1]
			if len(dirs) == 0 || dirs[len(dirs)-1] != e {
				dirs = append(dirs, e)
			}
		}
	}

	return dirs, pages
}
