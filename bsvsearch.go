package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	bsv := NewBible("Biblia Sacra Vulgata")
	content := loadLinesFromFile("./vuldat.txt")
	bsv, err := bsv.Load(content)
	if err != nil {
		log.Fatal(err)
	}

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", bsv.formHandler)

	fmt.Printf("Starting server on port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
