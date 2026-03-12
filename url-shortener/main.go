package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Request struct {
	URL string `json:"url"`
}

var store = map[string]string{}

func shorten(w http.ResponseWriter, r *http.Request) {

	var req Request
	json.NewDecoder(r.Body).Decode(&req)

	code := "abc123"

	store[code] = req.URL

	resp := map[string]string{
		"short": code,
	}

	json.NewEncoder(w).Encode(resp)
}

func redirect(w http.ResponseWriter, r *http.Request) {

	code := strings.TrimPrefix(r.URL.Path, "/")

	url := store[code]

	if url == "" {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}

func health(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("ok"))
}

func main() {

	http.HandleFunc("/shorten", shorten)
	http.HandleFunc("/health", health)
	http.HandleFunc("/", redirect)

	fmt.Println("Server running on :8080")

	http.ListenAndServe("0.0.0.0:8080", nil)
}