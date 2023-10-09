package api

import (
	"APIgateway/pcg/types"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// VerificationServiceURL представляет URL вашего сервиса верификации.
var VerificationServiceURL = "http://localhost:8081" // Замените на фактический URL

// VerifyComment отправляет запрос на верификацию комментария в сервис верификации.
func VerifyComment(commentText, uniqueID string) (string, error) {
	verificationRequestBody := []byte(fmt.Sprintf(`{"commentText":"%s", "uniqueID":"%s"}`, commentText, uniqueID))
	verificationResponse, err := http.Post(VerificationServiceURL+"/verify", "application/json", bytes.NewBuffer(verificationRequestBody))
	if err != nil {
		return "", err
	}

	verificationResponseData, err := ioutil.ReadAll(verificationResponse.Body)
	if err != nil {
		return "var1", err
	}

	if err := json.Unmarshal(verificationResponseData, &types.VerificationResult); err != nil {
		return "var2", err
	}

	if types.VerificationResult.Error != "" {
		return "var3", fmt.Errorf(types.VerificationResult.Error)
	}

	return types.VerificationResult.Message, nil
}
