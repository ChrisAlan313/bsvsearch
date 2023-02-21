package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ChrisAlan313/bible"
)

func main() {
	bsv := bible.New("Biblia Sacra Vulgata", "./vuldat.txt")

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/form", bsv.formHandler)

	fmt.Printf("Starting server on port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
