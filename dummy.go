package main

import (
	"encoding/json"
	"net/http"
)

type Dummy struct {
	ID      string `json:"id,omitempty"`
	Content string `json:"content,omitempty"`
}

var dummys []Dummy

func insertdata() {
	dummys = append(dummys, Dummy{ID: "1", Content: "GO Heroku! GO"})
}

func DummyJSON(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(dummys)
}
