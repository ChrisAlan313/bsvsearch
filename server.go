package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/ChrisAlan313/bible"
)

func StartServer() {

}

func (b bible.Bible) formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	bookQuery := r.FormValue("book")

	bf := bible.BetterFilter{}
	spec := bible.BookSpecification{bookQuery}
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
