package main

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home!")
}
