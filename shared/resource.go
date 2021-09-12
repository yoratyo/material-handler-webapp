package shared

import "gorm.io/gorm"

type (
	Resource struct {
		DB *gorm.DB
	}
)

// NewResource .
func NewResource(database *gorm.DB) Resource {
	return Resource{
		DB: database,
	}
}
