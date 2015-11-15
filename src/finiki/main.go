package main

import "fmt"

func main() {
	f := NewFolder("sample")

	printFolder(f, "")
}

func printFolder(f Folder, parent string) {
	for name, page := range f.pages {
		fmt.Printf("%s - %s\n", parent+name, (*page).String())
	}

	for name, folder := range f.folders {
		printFolder(folder, parent+name+"/")
	}
}
