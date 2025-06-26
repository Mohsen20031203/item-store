package internal

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
)

// UpdateItem godoc
// @Summary      Update an item
// @Description  Updates the value of an item by its key
// @Tags         database
// @Accept       octet-stream
// @Produce      json
// @Param        key   path      string  true  "Item Key"
// @Param        data  body      AnyJSON  true  "Any JSON object"
// @Success      200   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /items/{key} [put]
func (s *Server) UpdateItem(c *gin.Context) {

	key := c.Param("key")

	_, err := s.DB.Get(key)
	if err != nil {
		c.JSON(404, gin.H{"error": "Item not found"})
		return
	}

	if key == "" {
		c.JSON(400, gin.H{"error": "Key cannot be empty"})
		return
	}

	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot read body"})
		return
	}
	var inf interface{}
	if err := json.Unmarshal(bodyBytes, &inf); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}

	err = s.DB.Put(key, bodyBytes)
	if err != nil {
		c.JSON(500, gin.H{"error": "Cannot put item"})
		return
	}
	c.JSON(200, gin.H{"id": key, "message": "Item updated successfully"})
}
