package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8090", nil)
}

func hello(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("Nylas-Object-Count", "123")
	w.Header().Add("Nylas-Application-Id", "2jn0bneuyz4tlm9zj3j9vhn27")
	fmt.Fprint(w, "Hello world")
}
