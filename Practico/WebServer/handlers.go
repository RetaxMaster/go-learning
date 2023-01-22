package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home!")
}

func PostRequest(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var metadata MetaData

	err := decoder.Decode(&metadata)

	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}

	fmt.Fprintf(w, "Payload %v\n", metadata)

}

func UserPostRequest(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var user User

	err := decoder.Decode(&user)

	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}

	response, err := user.ToJSON()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}

/*

curl -X POST -H "Content-Type: application/json" \
     -d '{"name": "linuxize", "email": "linuxize@example.com"}' \
     localhost:3000/create

*/
