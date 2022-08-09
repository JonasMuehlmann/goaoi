package main

import (
	"log"
	"os"
	"text/template"
)

var variants = []string{"Slice", "Map", "String", "Iterator"}

func main() {
	templateSource := template.Must(template.ParseFiles("goaoi.go.tpl"))

	outFile, err := os.Create("goaoi.go")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	err = templateSource.Execute(outFile, variants)
	if err != nil {
		log.Fatal(err)
	}
}
