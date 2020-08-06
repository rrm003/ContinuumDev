package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

func form(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Println(err)
	}

	data := struct {
		Submissions url.Values
	}{req.Form}
	tpl.ExecuteTemplate(w, "net3.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("net3.gohtml"))
}

func main() {
	http.HandleFunc("/", form)
	http.ListenAndServe(":8090", nil)
}
