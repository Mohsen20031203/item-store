package internal

import (
	"io"

	"github.com/gin-gonic/gin"
)

// UpdateItem godoc
// @Summary      Update an item
// @Description  Updates the value of an item by its key
// @Tags         items
// @Accept       octet-stream
// @Produce      json
// @Param        key   path      string  true  "Item Key"
// @Param        data  body      string  true  "Raw binary data"
// @Success      200   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /{key} [put]
func (s *Server) UpdateItem(c *gin.Context) {

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
