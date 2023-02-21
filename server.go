package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/ChrisAlan313/bible"
)

type bsvServer struct {
	bible.Bible
}

func (s bsvServer) formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	b := s.Bible

	bookQuery := r.FormValue("book")

	bf := bible.BetterFilter{}
	spec := bible.BookSpecification{Book: bookQuery}
	verses := bf.Filter(b.Verses, spec)

	tmpl := template.Must(template.ParseFiles("templates/queryResponse.gohtml"))
	data := struct {
		Verses []bible.Verse
	}{verses}
	if err := tmpl.Execute(w, data); err != nil {
		fmt.Fprintf(w, "Execute() err: %v", err)
		return
	}
}
