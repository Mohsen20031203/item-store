package internal

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AnyJSON map[string]interface{}

// CreateItem godoc
// @Summary      Create a new item
// @Description  Creates a new item with a generated UUID and stores the request body in the database
// @Tags         database
// @Accept       json
// @Produce      json
// @Param        data  body      AnyJSON  true  "Any JSON object"
// @Success      200  {object}  map[string]string  "Returns the ID and success message"
// @Failure      400  {object}  map[string]string  "Cannot read body error"
// @Failure      500  {object}  map[string]string  "Cannot put item error"
// @Router       /items [post]
func (s *Server) CreateItem(c *gin.Context) {

	// Generate a new UUID for the item
	newUUID := uuid.New()

	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot read body"})
		return
	}

	if len(bodyBytes) == 0 {
		c.JSON(400, gin.H{"error": "Body cannot be empty"})
		return
	}
	var in interface{}
	if err := json.Unmarshal(bodyBytes, &in); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}

	err = s.DB.Put(newUUID.String(), bodyBytes)
	if err != nil {
		c.JSON(500, gin.H{"error": "Cannot put item"})
		return
	}
	c.JSON(200, gin.H{"id": newUUID.String(), "message": "Item created successfully"})
}
