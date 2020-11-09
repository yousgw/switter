package main

import (
	"fmt"
	"net/http"
	"test-json"
)

func IndexHandler(w http.ResponseWriter,
	r *http.Request) {
	test_json.Output_data(w)
	fmt.Fprint(w, "hello world")
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":3000", nil)
}