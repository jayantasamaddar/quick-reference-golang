package lib

import (
	"log"
	"os"
)

type Persistence struct {
	LineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	var sep string
	if p.LineSeparator != "" {
		sep = p.LineSeparator
	} else {
		sep = "\n"
	}
	err := os.WriteFile(filename, []byte(j.String(sep)), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
