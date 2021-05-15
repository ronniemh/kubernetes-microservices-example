package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
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

	locationString, _ := json.Marshal(true)

	text := []byte(locationString)
	key := []byte("S]5WC3+3W4LLT%w_*)X6TG[SR:eTdzvF")

	ciphertext, err := encrypt(text, key)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(ciphertext)
}

func encrypt(plaintext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", createEvent).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8080", router))
}
