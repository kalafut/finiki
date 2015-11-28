package main

import "testing"

func TestInterface(t *testing.T) {
	s := NewFlatFileStorage("/tmp")

	testGetPage(t, s)
}
