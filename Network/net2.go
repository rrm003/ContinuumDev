package main

import (
	"fmt"
	"net/http"
)

func homepage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>homepage\n</h1>")
}

func sec1(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>Section1\n</h1>")
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "<h2>%v: %v</h2>\n", name, h)
		}
	}
}

func sec1_1(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>Section1.1\n</h1>")
}

func sec1_2(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>Section1.2\n</h1>")
}

func sec2(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>Section2\n</h1>")
}

func main() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/sec1", sec1)
	http.HandleFunc("/sec1/sec1.1", sec1_1)
	http.HandleFunc("/sec1/sec1.2", sec1_2)
	http.HandleFunc("/sec2", sec2)
	http.ListenAndServe(":8090", nil)
}
