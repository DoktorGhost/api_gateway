package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// VerificationServiceURL представляет URL вашего сервиса верификации.
var VerificationServiceURL = "http://localhost:8081"

var VerificationResult struct {
	UniqueID string `json:"uniqueID"`
	Message  string `json:"message"`
	Error    string `json:"error"`
}

// VerifyComment отправляет запрос на верификацию комментария в сервис верификации.
func VerifyComment(commentText, uniqueID string) (string, error) {
	verificationRequestBody := []byte(fmt.Sprintf(`{"commentText":"%s", "uniqueID":"%s"}`, commentText, uniqueID))
	verificationResponse, err := http.Post(VerificationServiceURL+"/verify", "application/json", bytes.NewBuffer(verificationRequestBody))
	if err != nil {
		return "", err
	}

	verificationResponseData, err := ioutil.ReadAll(verificationResponse.Body)
	if err != nil {
		return "", err
	}

	if err := json.Unmarshal(verificationResponseData, &VerificationResult); err != nil {
		return "", err
	}

	if VerificationResult.Error != "" {
		return "", fmt.Errorf(VerificationResult.Error)
	}

	return VerificationResult.Message, nil
}
