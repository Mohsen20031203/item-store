package internal

import (
	"io"

	"github.com/gin-gonic/gin"
)

func (s *Server) Update(c *gin.Context) {

	key := c.Param("key")

	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot read body"})
		return
	}

	err = s.DB.Put(key, bodyBytes)
	if err != nil {
		c.JSON(500, gin.H{"error": "Cannot put item"})
		return
	}
	c.JSON(200, gin.H{"id": key, "message": "Item updated successfully"})
}
