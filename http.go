package main

import (
	"fmt"
	"net/http"
)

func HandleHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, HTTP!")
}
