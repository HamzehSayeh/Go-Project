package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type InfobipOTPRequest struct {
	To       string `json:"to"`
	Message  string `json:"message"`
	Language string `json:"language"`
}

func SendOTP(phoneNumber string) error {
	apiKey := os.Getenv("INFOBIP_API_KEY")
	baseUrl := os.Getenv("INFOBIP_BASE_URL")

	url := fmt.Sprintf("%s/sms/2/text/single", baseUrl)
	otpMessage := InfobipOTPRequest{
		To:       phoneNumber,
		Message:  "Your OTP code is 123456", // Customize or generate OTP
		Language: "en",
	}

	jsonData, _ := json.Marshal(otpMessage)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", fmt.Sprintf("App %s", apiKey))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send OTP, status code: %d", resp.StatusCode)
	}

	return nil
}
