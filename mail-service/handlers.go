package main

import (
	"encoding/json"
	"net/http"
)

type MailRequest struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Message string   `json:"message"`
	Type    string   `json:"type"`
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

	if err := sendMailGoMail(msg.From, msg.To, msg.Subject, msg.Message, msg.Type); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{
			"success": false,
			"message": "Terdapat kesalahan saat pengiriman email",
			"error":   err.Error(),
		})
		return
	}

	// set response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": "Berhasil kirim email!",
	})
}
