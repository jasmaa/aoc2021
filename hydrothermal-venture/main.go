package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type SuccessResponse struct {
	Overlaps int `json:"overlaps"`
}
type ErrorResponse struct {
	Error string `json:"error"`
}

func main() {
	http.HandleFunc("/overlaps-limited", func(w http.ResponseWriter, r *http.Request) {
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
		lines, err := LinesFromString(string(rawdata))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp := ErrorResponse{Error: "invalid input"}
			json.NewEncoder(w).Encode(resp)
			return
		}
		overlaps := FindOverlapsLimited(lines)
		resp := SuccessResponse{Overlaps: overlaps}
		json.NewEncoder(w).Encode(resp)
	})

	http.HandleFunc("/overlaps-all", func(w http.ResponseWriter, r *http.Request) {
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
		lines, err := LinesFromString(string(rawdata))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp := ErrorResponse{Error: "invalid input"}
			json.NewEncoder(w).Encode(resp)
			return
		}
		overlaps := FindOverlapsAll(lines)
		resp := SuccessResponse{Overlaps: overlaps}
		json.NewEncoder(w).Encode(resp)
	})

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
