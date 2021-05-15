package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type location struct {
	X float64 `json:"xPoint"`
	Y float64 `json:"yPoint"`
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var locationToEncrypt location

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error: can't read body")
	}

	json.Unmarshal(reqBody, &locationToEncrypt)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(locationToEncrypt)
}
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", createEvent).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8080", router))
}
