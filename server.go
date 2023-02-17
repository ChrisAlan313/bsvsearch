package main

import (
	"fmt"
	"net/http"
)

func StartServer() {

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func (b Bible) formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	bookQuery := r.FormValue("book")

	bf := BetterFilter{}
	spec := BookSpecification{bookQuery}
	verses := bf.Filter(b.verses, spec)

	content := withHtmlBoilerplate(verses)

	fmt.Fprint(w, content)
}

func withHtmlBoilerplate(verses []Verse) string {
	startResponse := "<html><head><title>BSV Search Results</title></head><body><h1>BSV Search Results</h1><p>"
	vSeparator := "<br/>"
	endResponse := "</p></body><html>"

	var response string
	response += startResponse
	for _, v := range verses {
		response += v.content
		response += vSeparator
	}
	response += endResponse

	return response
}
