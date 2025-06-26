package internal

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Server) PutItem(c *gin.Context) {

	newUUID := uuid.New()

	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot read body"})
		return
	}

	err = s.DB.Put(newUUID.String(), bodyBytes)
	if err != nil {
		c.JSON(500, gin.H{"error": "Cannot put item"})
		return
	}
	c.JSON(200, gin.H{"id": newUUID.String(), "message": "Item created successfully"})
}
