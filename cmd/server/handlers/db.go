package handlers

import "gorm.io/gorm"

var db *gorm.DB

// SetDB sets the global database instance for the handlers package.
func SetDB(database *gorm.DB) {
	db = database
}
