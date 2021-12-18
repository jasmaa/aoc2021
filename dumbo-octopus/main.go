package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type SuccessResponseFlashes struct {
	Flashes int `json:"flashes"`
}
type SuccessResponseStep struct {
	Step int `json:"step"`
}
type ErrorResponse struct {
	Error string `json:"error"`
}

func main() {
	http.HandleFunc("/flashes-100", func(w http.ResponseWriter, r *http.Request) {
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
		grid, err := GridFromString(string(rawdata))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp := ErrorResponse{Error: err.Error()}
			json.NewEncoder(w).Encode(resp)
			return
		}
		flashes := FindTotalFlashes(grid, 100)
		resp := SuccessResponseFlashes{Flashes: flashes}
		json.NewEncoder(w).Encode(resp)
	})

	http.HandleFunc("/first-all-flash", func(w http.ResponseWriter, r *http.Request) {
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
		grid, err := GridFromString(string(rawdata))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp := ErrorResponse{Error: err.Error()}
			json.NewEncoder(w).Encode(resp)
			return
		}
		step := FindFirstAllFlash(grid)
		resp := SuccessResponseStep{Step: step}
		json.NewEncoder(w).Encode(resp)
	})

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
