package internal

import "github.com/gin-gonic/gin"

func (s *Server) Get(c *gin.Context) {
	key := c.Param("key")
	item, err := s.DB.Get(key)
	if err != nil {
		c.JSON(404, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(200, gin.H{"id": key, "item": item})
}
