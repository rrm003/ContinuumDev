package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		fmt.Println("Error in GET")
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	f, err := os.Create("net1_out.html")
	fmt.Fprintf(f, string(body), nil)
}
