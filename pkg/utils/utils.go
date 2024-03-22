package utils

import "math/rand"

// Create random string with length
func RandomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	template := make([]rune, length)

	for i := range template {
		template[i] = letters[rand.Intn(len(letters))]
	}

	return string(template)
}
