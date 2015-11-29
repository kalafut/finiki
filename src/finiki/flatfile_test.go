package finiki

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestInterface(t *testing.T) {
	dir, _ := ioutil.TempDir("", "")
	s := NewFlatFileStorage(dir)

	testStorage(t, s)
	os.RemoveAll(dir)
}
