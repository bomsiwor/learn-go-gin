package utils

import (
	"math/rand"

	"gorm.io/gorm"
)

// Create random string with length
func RandomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	template := make([]rune, length)

	for i := range template {
		template[i] = letters[rand.Intn(len(letters))]
	}

	return string(template)
}

// Paginate retrieve query using offset & limit
func Paginate(page, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// If page lower or equal than  0, force reset to 1
		if page <= 0 {
			page = 1
		}

		// If limit lower or equal than 0, force reset to 10
		pageSize := limit
		switch {
		case (pageSize > 100):
			pageSize = 100

		case pageSize <= 0:
			pageSize = 10
		}

		// Calculate offset
		offset := (page - 1) * pageSize

		return db.Offset(offset).Limit(pageSize)
	}
}
