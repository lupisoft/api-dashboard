package repository

import (
	"main/domain"
)

// CategoryRepository represent repository of the category
// Expect implementation by the infrastructure layer
type CategoryRepository interface {
	GetAll() ([]*domain.Category, error)
}
