package internal

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// GetItem godoc
// @Summary      Get item by key
// @Description  Retrieves an item from the database using the given key
// @Tags         database
// @Produce      json
// @Param        key   path      string  true  "Item Key"
// @Success      200   {object}  map[string]interface{}  "Returns the item with its ID"
// @Failure      404   {object}  map[string]string       "Item not found error"
// @Router       /items/{key} [get]
func (s *Server) GetItem(c *gin.Context) {
	key := c.Param("key")

	if key == "" {
		c.JSON(400, gin.H{"error": "Key cannot be empty"})
		return
	}
	item, err := s.DB.Get(key)
	if err != nil {
		c.JSON(404, gin.H{"error": "Item not found"})
		return
	}
	var inf interface{}
	if err := json.Unmarshal(item, &inf); err != nil {
		c.JSON(500, gin.H{"error": "Cannot unmarshal item value"})
		return
	}
	c.JSON(200, inf)
}
