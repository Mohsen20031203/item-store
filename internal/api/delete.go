package internal

import "github.com/gin-gonic/gin"

func (s *Server) DeleteItem(c *gin.Context) {
	key := c.Param("key")

	err := s.DB.Delete(key)
	if err != nil {
		c.JSON(404, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(200, gin.H{"message": "Item deleted successfully"})
}
