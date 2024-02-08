package handlers

import (
	"net/http"
	"strconv"
	"utkarsh/Fetch/api/db"
	"utkarsh/Fetch/api/models"

	"github.com/gin-gonic/gin"
)

// Handler for point request end point. Get points for a reciept associated with an id
func GetPoints(c *gin.Context) {
	// Geet id from get request
	id := c.Param("id")

	// Get points for the associated reciept and convert it to int if possible.
	points := db.GetPoints(id)
	numericPoints, err := strconv.Atoi(points)
	if err == nil {
		// Return points in JSON format
		pointsModel := models.PointsModel{
			Points: numericPoints,
		}
		c.IndentedJSON(http.StatusCreated, pointsModel)
	} else {
		// Return a 400 Bad Request response if any error was caused
		c.JSON(http.StatusBadRequest, gin.H{"error": points})
	}

}
