package internal

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// ListItem godoc
// @Summary      List all item keys
// @Description  Retrieves a list of all item keys stored in the database
// @Tags         database
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "Returns list of keys and a success message"
// @Failure      500  {object}  map[string]string       "Cannot list items error"
// @Router       /items [get]
func (s *Server) ListItem(c *gin.Context) {

	// List all items in the database
	kvs, err := s.DB.List()
	if err != nil {
		c.JSON(500, gin.H{"error": "Cannot list items"})
		return
	}

	// Create a map to hold the results
	result := make(map[string]interface{})
	for key, value := range kvs {
		//unmarshal value to interface{}

		// Check if value is empty
		var inf interface{}
		if err := json.Unmarshal(value, &inf); err != nil {
			c.JSON(500, gin.H{"error": "Cannot unmarshal item value"})
			return
		}
		result[key] = inf
	}
	// Return the result as JSON
	c.JSON(200, result)
}
