package main

import (
	"github.com/ancientlore/unhtml"
	"log"
	"os"
)

func main() {
	err := unhtml.HtmlToText(os.Stdin, os.Stdout)
	if err != nil {
		log.Print(err)
	}
}
