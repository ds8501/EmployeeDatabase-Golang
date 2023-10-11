package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html")

}

func about(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html")

	//fmt.Fprintf(w, "this is about page and 2+3 is ")
}

func renderTemplate(w http.ResponseWriter, r string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+r, "./templates/base.layout.html")
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error found", err)
		return
	}
}
