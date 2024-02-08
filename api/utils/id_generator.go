package utils

import (
	"math/rand"
)

// Function to generate a random alphanumeric string of a given length
func generateRandomString(length int) string {
	// Define the character set i.e. Alphanumeric
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Create a byte slice with the desired length
	randomString := make([]byte, length)

	// Fill the byte slice with random characters from the character set
	for i := range randomString {
		randomString[i] = charSet[rand.Intn(len(charSet))]
	}

	// Convert the byte slice to a string and return it
	return string(randomString)
}

// Function to return new ID
func GetNewID() string {
	// Generate ID of length type 8-4-4-4-12 to match "7fb1377b-b223-49d9-a31a-5a02701dd310" format
	newID := generateRandomString(8) + "-" + generateRandomString(4) + "-" + generateRandomString(4) + "-" + generateRandomString(4) + "-" + generateRandomString(12)
	return string(newID)
}
