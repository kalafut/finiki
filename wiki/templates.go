package wiki

import (
	"fmt"
	"io/ioutil"
	"log"
)

func loadTemplate(name string) string {
	b, err := ioutil.ReadFile(fmt.Sprintf("templates/%s.html", name))
	if err != nil {
		log.Fatal(err)
	}

	return string(b)
}
