package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type SuccessResponse struct {
	Increases int `json:"increases"`
}
type ErrorResponse struct {
	Error string `json:"error"`
}

func main() {
	http.HandleFunc("/one-window", func(w http.ResponseWriter, r *http.Request) {
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
		depths, err := DepthsFromString(string(rawdata))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp := ErrorResponse{Error: "invalid input"}
			json.NewEncoder(w).Encode(resp)
			return
		}
		increases := CountDepthIncreasesOneWindow(depths)
		resp := SuccessResponse{Increases: increases}
		json.NewEncoder(w).Encode(resp)
	})

	http.HandleFunc("/three-window", func(w http.ResponseWriter, r *http.Request) {
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
		depths, err := DepthsFromString(string(rawdata))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp := ErrorResponse{Error: "invalid input"}
			json.NewEncoder(w).Encode(resp)
			return
		}
		increases := CountDepthIncreasesThreeWindow(depths)
		resp := SuccessResponse{Increases: increases}
		json.NewEncoder(w).Encode(resp)
	})

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
