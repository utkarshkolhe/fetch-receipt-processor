package db

import (
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
	"utkarsh/Fetch/api/models"
)

// A Map used as a stand in for database
var datababse = map[string]models.RecieptModel{}

// Add a recipt to database using ID as key
func AddToDatabase(id string, reciept models.RecieptModel) {
	datababse[id] = reciept
}

// Check if a ID of a reciept exists in database.
func ExistsInDatabase(id string) bool {
	if _, ok := datababse[id]; ok {
		return true
	}
	return false
}

// Get a Reciept from database using its associated id. Return empty recipt if not present
func GetReciept(id string) models.RecieptModel {
	if value, ok := datababse[id]; ok {
		return value
	}
	return models.RecieptModel{}
}

// Calculate number of alphanumeric characters in a string
func getAlphaNumericChars(name string) int {
	count := 0

	// Iterate over each character in the string
	for _, char := range name {
		// Check if the character is alphanumeric
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

// Calculate and return points for a ID of reciept. Returns in string format.
func GetPoints(id string) string {
	if value, ok := datababse[id]; ok {

		var points = 0
		// Add number of alpanumeric characters as points
		points += getAlphaNumericChars(value.Retailer)

		// Convert total points to int
		f, err := strconv.ParseFloat(value.Total, 64)
		if err != nil {
			// Handle the error if the conversion fails
			return "Error converting total points to int:" + value.Total + ". " + err.Error()
		}
		// If total is Integer, add 50 points
		if f == math.Trunc(f) {
			points += 50
		}
		// If total is a multiple of 0.25, add 25 points
		remainder := math.Mod(f, 0.25)
		if math.Abs(remainder) < 1e-10 {
			points += 25
		}

		//Add 5 points for every two pairs in items
		points += (5 * (len(value.Items) / 2))

		// For every item trimmed short description add correspoinding poinst
		for _, rec := range value.Items {
			// Trim the item short description
			trimmed := strings.TrimSpace(rec.ShortDescription)
			// Add points if trimmed desc length is divisible by 3
			if len(trimmed)%3 == 0 {
				price, err := strconv.ParseFloat(rec.Price, 64)
				if err != nil {
					// Handle the error if the conversion fails
					return "Error converting price of " + trimmed + "  to int:" + rec.Price + ". " + err.Error()
				}
				points += int(math.Ceil(price * 0.2))
			}
		}

		// Parse the date string into a time.Time object
		date, err := time.Parse("2006-01-02", value.PurchaseDate)
		if err != nil {
			return "Error Parsing Purchase Date" + value.PurchaseDate + ". " + err.Error()
		}
		// Extract the day component from the parsed date
		day := date.Day()
		// Add 6 points if date is odd.
		if day%2 == 1 {
			points += 6
		}
		// Create time objects for 2pm and 4 pm
		startTimeStr := "14:00"
		endTimeStr := "16:00"
		startTime, err := time.Parse("15:04", startTimeStr)
		if err != nil {
			return "Error Parsing 2pm to time" + err.Error()
		}
		endTime, err := time.Parse("15:04", endTimeStr)
		if err != nil {
			return "Error Parsing 4pm to time" + err.Error()
		}

		// Parse the Purchase time into a time object
		checkTime, err := time.Parse("15:04", value.PurchaseTime)
		if err != nil {
			return "Error Parsing Purcahse time" + value.PurchaseTime + "." + err.Error()
		}
		// Add 10 poinst if purchase time between 2pm and 4pm
		if checkTime.After(startTime) && checkTime.Before(endTime) {
			points += 10
		}
		return strconv.Itoa(points)
	}
	return "ID not found in database"
}
