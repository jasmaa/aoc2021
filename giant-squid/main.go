package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type SuccessResponse struct {
	Score int `json:"score"`
}
type ErrorResponse struct {
	Error string `json:"error"`
}

func main() {
	http.HandleFunc("/first", func(w http.ResponseWriter, r *http.Request) {
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
		game, err := GameFromString(string(rawdata))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp := ErrorResponse{Error: err.Error()}
			json.NewEncoder(w).Encode(resp)
			return
		}
		score, err := game.FindFirstWinningScore()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp := ErrorResponse{Error: err.Error()}
			json.NewEncoder(w).Encode(resp)
			return
		}
		resp := SuccessResponse{Score: score}
		json.NewEncoder(w).Encode(resp)
	})

	http.HandleFunc("/last", func(w http.ResponseWriter, r *http.Request) {
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
		game, err := GameFromString(string(rawdata))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp := ErrorResponse{Error: err.Error()}
			json.NewEncoder(w).Encode(resp)
			return
		}
		score, err := game.FindLastWinningScore()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp := ErrorResponse{Error: err.Error()}
			json.NewEncoder(w).Encode(resp)
			return
		}
		resp := SuccessResponse{Score: score}
		json.NewEncoder(w).Encode(resp)
	})

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
