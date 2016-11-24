package main

import "errors"

var (
	ErrPageNotFound = errors.New("Page Not Found")
	ErrRevNotFound  = errors.New("Rev Not Found")
	ErrPageCorrupt  = errors.New("Page Corrupt")
	ErrFolderExists = errors.New("Folder already exists")
)
