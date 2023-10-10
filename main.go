package main

import (
	"APIgateway/pcg/api"
	"APIgateway/pcg/types"
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

		var request types.Comment
		request.UniqueID = xid.New().String()
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			log.Printf("Timestamp: %s, Request ID: %s, IP: %s, HTTP Code: %d", time.Now().Format("2006-01-02 15:04:05"), request.UniqueID, c.ClientIP(), http.StatusBadRequest)
			return
		}
		fmt.Println(request)
		message, err := api.VerifyComment(request.CommentText, request.UniqueID)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			log.Printf("Timestamp: %s, Request ID: %s, IP: %s, HTTP Code: %d", time.Now().Format("2006-01-02 15:04:05"), request.UniqueID, c.ClientIP(), http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusOK, gin.H{"uniqueID": request.UniqueID, "message": message})
		log.Printf("Timestamp: %s, Request ID: %s, IP: %s, HTTP Code: %d", time.Now().Format("2006-01-02 15:04:05"), request.UniqueID, c.ClientIP(), http.StatusOK)

		err = api.AddComment(request.NewsID, request.ParentCommentID, request.CommentText, request.UniqueID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			log.Printf("Timestamp: %s, Request ID: %s, IP: %s, HTTP Code: %d", time.Now().Format("2006-01-02 15:04:05"), request.UniqueID, c.ClientIP(), http.StatusInternalServerError)
			return
		}
	})

	router.Run(":8080")
}
