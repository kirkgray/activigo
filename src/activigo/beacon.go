package main

import (
	"fmt"
	"net/http"
)

type Activity struct {
    Headers []string
    Body  []byte
}

func handler(w http.ResponseWriter, r *http.Request) {
	a, _ := loadActivity( r )
}

func loadActivity(r *http.Request) (*Activity, error) {
	
	header = r.Header
	
	return &Activity{Headers: r.Header, Body: r.Body}, nil
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
