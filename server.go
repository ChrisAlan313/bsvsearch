package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func StartServer() {

}

func (b Bible) formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	bookQuery := r.FormValue("book")

	bf := BetterFilter{}
	spec := BookSpecification{bookQuery}
	verses := bf.Filter(b.Verses, spec)

	tmpl := template.Must(template.ParseFiles("templates/queryResponse.gohtml"))
	data := struct {
		Verses []Verse
	}{verses}
	if err := tmpl.Execute(w, data); err != nil {
		fmt.Fprintf(w, "Execute() err: %v", err)
		return
	}
}
