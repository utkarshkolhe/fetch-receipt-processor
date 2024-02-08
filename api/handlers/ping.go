package handlers

import (
	"net/http"
	"utkarsh/Fetch/api/models"

	"github.com/gin-gonic/gin"
)

// A test End Point to check if services are online.
func GetPing(c *gin.Context) {
	idModel := models.IDModel{
		ID: "Online",
	}
	c.IndentedJSON(http.StatusCreated, idModel)
}
