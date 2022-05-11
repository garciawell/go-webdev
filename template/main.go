package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("index.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("index.html")

	if err != nil {
		log.Println("error creating file", err)
	}

	defer nf.Close()

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}

}
