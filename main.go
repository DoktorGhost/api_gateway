package main

import (
	"APIgateway/pcg/api"
	"APIgateway/pcg/types"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func main() {

	router := gin.Default()

	router.POST("/create-comment", func(c *gin.Context) {

		var request types.Request
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

	router.DELETE("/del-comment/:id", func(c *gin.Context) {
		var request types.Request
		commentID := c.Param("id")

		commentIDInt, err := strconv.Atoi(commentID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
			return
		}

		request.UniqueID = xid.New().String()
		request.ID = commentIDInt

		err = api.DeleteComment(request.ID, request.UniqueID)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			log.Printf("Timestamp: %s, Request ID: %s, IP: %s, HTTP Code: %d", time.Now().Format("2006-01-02 15:04:05"), request.UniqueID, c.ClientIP(), http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, gin.H{"commentID": commentID, "message": "Comment deleted successfully"})
		log.Printf("Timestamp: %s, Request ID: %s, IP: %s, HTTP Code: %d", time.Now().Format("2006-01-02 15:04:05"), request.UniqueID, c.ClientIP(), http.StatusOK)
	})

	router.GET("/get-comment/:id", func(c *gin.Context) {
		commentID := c.Param("id")

		commentIDInt, err := strconv.Atoi(commentID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
			return
		}

		uniqueID := xid.New().String()
		comment, err := api.GetComment(commentIDInt, uniqueID)
		fmt.Println(comment)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			log.Printf("Timestamp: %s, Request ID: %s, IP: %s, HTTP Code: %d", time.Now().Format("2006-01-02 15:04:05"), uniqueID, c.ClientIP(), http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, comment)
		log.Printf("Timestamp: %s, Request ID: %s, IP: %s, HTTP Code: %d", time.Now().Format("2006-01-02 15:04:05"), uniqueID, c.ClientIP(), http.StatusOK)
	})

	router.GET("/get-comments/:newsId", func(c *gin.Context) {
		commentID := c.Param("newsId")

		commentIDInt, err := strconv.Atoi(commentID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
			return
		}

		uniqueID := xid.New().String()
		comment, err := api.GetCommentsByNewsID(commentIDInt, uniqueID)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			log.Printf("Timestamp: %s, Request ID: %s, IP: %s, HTTP Code: %d", time.Now().Format("2006-01-02 15:04:05"), uniqueID, c.ClientIP(), http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, comment)
		log.Printf("Timestamp: %s, Request ID: %s, IP: %s, HTTP Code: %d", time.Now().Format("2006-01-02 15:04:05"), uniqueID, c.ClientIP(), http.StatusOK)
	})

	router.Run(":8080")
}
