package handlers

import (
	"encoding/json"
	"github.com/HamzehSayeh/go-project/services"
	"net/http"
)

type OTPRequest struct {
	PhoneNumber string `json:"phone_number"`
}

func SendOTP(w http.ResponseWriter, r *http.Request) {
	var otpReq OTPRequest
	if err := json.NewDecoder(r.Body).Decode(&otpReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Call the Infobip service to send the OTP
	err := services.SendOTP(otpReq.PhoneNumber)
	if err != nil {
		http.Error(w, "Failed to send OTP", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OTP sent successfully"))
}
