package persistence

import (
	"main/config"
	"main/domain"
	"main/domain/repository"
)

type categoryRepository struct {
	dbClient config.DBClient
}

func NewCategoryRepository(client config.DBClient) repository.CategoryRepository {
	return &categoryRepository{dbClient: client}
}

func (c categoryRepository) GetAll() ([]*domain.Category, error) {
	categories := make([]*domain.Category, 0)
	db, err := c.dbClient.GetConnection(c.dbClient.Dialect, c.dbClient.StringConnection)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err := db.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}
