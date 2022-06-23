/*
	A "hello world" plugin in Go,
	which reads a request header and sets a response header.
*/

package main

import (
	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
)

func main() {
	server.StartServer(New, Version, Priority)
}

var Version = "0.2"
var Priority = 1

type Config struct {
	Message string
}

func New() interface{} {
	return &Config{}
}

func (conf Config) Access(kong *pdk.PDK) {
	message := conf.Message
	if message == "" {
		message = "hello"
	}
	kong.Log.Notice("Message: " + message)

	// Set target
	kong.Service.SetTarget("mockbin.com", 80)

	// Set headers
	for key, value := range getRequestHeaders() {
		err := kong.ServiceRequest.SetHeader(key, value)
		if err != nil {
			kong.Log.Err("Error when setting request header: " + err.Error())
		}
	}

	// Set query params
	kong.ServiceRequest.SetQuery(map[string][]string{"hello-query": {"world"}})

	// Set header again
	kong.ServiceRequest.SetHeader("x-lowercase-zhi", "qu")
}

func getRequestHeaders() map[string]string {
	return map[string]string{
		"X-Nylas-Name":             "Zhi Qu",
		"X-Nylas-Email":            "hello@qu.com",
		"X-Nylas-Email1":           "hello@qu.com",
		"X-Nylas-Email11":          "hello@qu.com",
		"X-Nylas-Email111":         "hello@qu.com",
		"X-Nylas-Email1111":        "hello@qu.com",
		"X-Nylas-Email11111":       "hello@qu.com",
		"X-Nylas-Email111111":      "hello@qu.com",
		"X-Nylas-Email1111111":     "hello@qu.com",
		"X-Nylas-Email11111111":    "hello@qu.com",
		"X-Nylas-Email111111111":   "hello@qu.com",
		"X-Nylas-Email1111111111":  "hello@qu.com",
		"X-Nylas-Email11111111111": "hello@qu.com",
	}
}
