package handlers

import (
	"net/http"
	"utkarsh/Fetch/api/db"
	"utkarsh/Fetch/api/models"
	"utkarsh/Fetch/api/utils"

	"github.com/gin-gonic/gin"
)

// Handler for end point to add reciepts to database and get back an generated id
func GetID(c *gin.Context) {
	// Get the reciept from the POST Request
	var newReciept models.RecieptModel
	if err := c.BindJSON(&newReciept); err != nil {
		// Return 400 if there was any Error with parsing the reciept.
		c.JSON(http.StatusBadRequest, gin.H{"Error while parsing reciept": err.Error()})
		return
	}

	var newId = ""
	for {
		// Get a new generated ID for the reciept
		newId = utils.GetNewID()
		// If ID is unique and does not exists in the database, break the for loop. Otherwise repeat.
		if db.ExistsInDatabase(newId) == false {
			break
		}
	}
	// Add the reciept with the new generated id to database
	db.AddToDatabase(newId, newReciept)

	// Return the ID in JSON format
	idModel := models.IDModel{
		ID: newId,
	}
	c.IndentedJSON(http.StatusCreated, idModel)

}
