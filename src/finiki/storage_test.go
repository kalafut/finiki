package finiki

import (
	"testing"

	"gopkg.in/tylerb/is.v1"
)

func testGetPage(t *testing.T, s Storage) {
	is := is.New(t)

	page, err := s.GetPage("/")

	is.Nil(page)
	is.Err(err)
}
