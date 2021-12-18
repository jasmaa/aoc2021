package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type SuccessResponse struct {
	Paths int `json:"paths"`
}
type ErrorResponse struct {
	Error string `json:"error"`
}

func main() {
	http.HandleFunc("/paths-without-revisit", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			resp := ErrorResponse{Error: "method not allowed"}
			json.NewEncoder(w).Encode(resp)
			return
		}
		rawdata, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp := ErrorResponse{Error: "invalid input"}
			json.NewEncoder(w).Encode(resp)
			return
		}
		caves, err := CavesFromString(string(rawdata))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp := ErrorResponse{Error: err.Error()}
			json.NewEncoder(w).Encode(resp)
			return
		}
		paths := FindAllPathsWithoutRevisit(caves)
		resp := SuccessResponse{Paths: len(paths)}
		json.NewEncoder(w).Encode(resp)
	})

	http.HandleFunc("/paths-with-one-revisit", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			resp := ErrorResponse{Error: "method not allowed"}
			json.NewEncoder(w).Encode(resp)
			return
		}
		rawdata, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp := ErrorResponse{Error: "invalid input"}
			json.NewEncoder(w).Encode(resp)
			return
		}
		caves, err := CavesFromString(string(rawdata))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp := ErrorResponse{Error: err.Error()}
			json.NewEncoder(w).Encode(resp)
			return
		}
		paths := FindAllPathsWithOneRevisit(caves)
		resp := SuccessResponse{Paths: len(paths)}
		json.NewEncoder(w).Encode(resp)
	})

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
