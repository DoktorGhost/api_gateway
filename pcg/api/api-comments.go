package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var CommentServiceURL = "http://localhost:8082" // Замените на фактический URL вашего сервиса комментариев

// CommentRequest представляет структуру для отправки комментария в сервис комментариев.
type CommentRequest struct {
	NewsID          int    `json:"news_id"`
	ParentCommentID int    `json:"parent_id"`
	CommentText     string `json:"commentText"`
	UniqueID        string `json:"uniqueID"`
}

// AddComment отправляет комментарий в сервис комментариев.
func AddComment(newsId, parentCommentId int, commentText, uniqueID string) error {
	commentRequest := CommentRequest{
		NewsID:          newsId,
		ParentCommentID: parentCommentId,
		CommentText:     commentText,
		UniqueID:        uniqueID,
	}

	requestBody, err := json.Marshal(commentRequest)
	if err != nil {
		return err
	}

	response, err := http.Post(CommentServiceURL+"/add-comment", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to add comment. Status code: %d", response.StatusCode)
	}

	return nil
}
