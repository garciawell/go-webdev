package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.html", 452)

	if err != nil {
		log.Fatal(err)
	}
}
