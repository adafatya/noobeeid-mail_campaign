package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const BASE_URL = "http://localhost:8001"

type MailRequest struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Message string   `json:"message"`
	Type    string   `json:"type"`
}

type SuccessResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type FailResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func sendEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]any{
			"error": "method not allowed",
		})
	}
	msg := MailRequest{}
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"success": false,
			"message": "Terdapat kesalahan pada badan request",
			"error":   err.Error(),
		})
		return
	}

	byteReq, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(byteReq)
	resp, err := http.Post(BASE_URL+"/send", "application/json", buf)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// set response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	if resp.StatusCode == http.StatusOK {
		body := SuccessResponse{}
		if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"success": false,
				"message": "Terdapat kesalahan pada badan response mail service",
				"error":   err.Error(),
			})
			return
		}
		json.NewEncoder(w).Encode(body)
		return
	}
	body := FailResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"success": false,
			"message": "Terdapat kesalahan pada badan response mail service",
			"error":   err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(body)
}
