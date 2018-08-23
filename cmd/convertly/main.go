package main

import (
	"fmt"
	"net/http"

	"github.com/willmadison/convertly"
)

func main() {

	http.Handle("/", convertly.RootHandler(convertly.NewFixerExchanger("43d8eaede49bad82ef3e2c4dbcfafbbe")))

	fmt.Println("Convertly serving requests on localhost:8888")
	http.ListenAndServe(":8888", nil)
}
