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
	w.Header().Add("Nylas-Namespace-Public-Id", "d94bher0h4rksgeut1zmuzmn1")
	w.Header().Add("Nylas-Namespace-Id", "12597569835")
	w.Header().Add("Nylas-Http-Request", "GET /messages?received_after=1645066330&limit=100&offset=0 HTTP/1.1")
	w.Header().Add("Nylas-Application-Public-Id", "bekpzuz9qqer3jsh43utlrm6x")
	w.Header().Add("Nylas-Application-Id", "2jn0bneuyz4tlm9zj3j9vhn27")
	w.Header().Add("Nylas-Endpoint", "namespace_api.send_api.send")
	w.Header().Add("Nylas-Request-Uid", "AC462A55:D9E6_0A4D00A6:01BB_620DB985_9FA5B:3CA6")
	w.Header().Add("Nylas-Event", "request handled")
	w.Header().Add("Nylas-Provider-Name", "gmail")
	w.Header().Add("Nylas-Result-Count", "123")
	fmt.Fprint(w, "Hello world")
}
