package internal

import "github.com/gin-gonic/gin"

func (s *Server) List(c *gin.Context) {

	keys, err := s.DB.List()
	if err != nil {
		c.JSON(500, gin.H{"error": "Cannot list items"})
		return
	}

	c.JSON(200, gin.H{"keys": keys, "message": "Items listed successfully"})
}
