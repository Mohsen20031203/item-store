package internal

import "github.com/gin-gonic/gin"

// DeleteItem godoc
// @Summary      Delete item by key
// @Description  Deletes an item from the database using the given key
// @Tags         database
// @Produce      json
// @Param        key   path      string  true  "Item Key"
// @Success      200   {object}  map[string]string  "Item deleted successfully"
// @Failure      404   {object}  map[string]string  "Item not found error"
// @Router       /items/{key} [delete]
func (s *Server) DeleteItem(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.JSON(400, gin.H{"error": "Key cannot be empty"})
		return
	}

	err := s.DB.Delete(key)
	if err != nil {
		c.JSON(404, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(200, gin.H{"message": "Item deleted successfully"})
}
