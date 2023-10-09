package main

import (
	"APIgateway/pcg/api"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func main() {

	router := gin.Default()

	router.POST("/create-comment", func(c *gin.Context) {
		uniqueID := xid.New().String()
		var request struct {
			CommentText string `json:"commentText"`
		}
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			log.Printf("Timestamp: %s, Request ID: %s, IP: %s, HTTP Code: %d", time.Now().Format("2006-01-02 15:04:05"), uniqueID, c.ClientIP(), http.StatusBadRequest)
			return
		}

		fmt.Println(request.CommentText)
		message, err := api.VerifyComment(request.CommentText, uniqueID)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(message)
			log.Printf("Timestamp: %s, Request ID: %s, IP: %s, HTTP Code: %d", time.Now().Format("2006-01-02 15:04:05"), uniqueID, c.ClientIP(), http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, gin.H{"uniqueID": uniqueID, "message": message})
		log.Printf("Timestamp: %s, Request ID: %s, IP: %s, HTTP Code: %d", time.Now().Format("2006-01-02 15:04:05"), uniqueID, c.ClientIP(), http.StatusOK)
	})

	router.Run(":8080") // Порт вашего API Gateway
}
