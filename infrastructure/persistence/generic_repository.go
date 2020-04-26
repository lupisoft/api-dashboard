package persistence

import (
	"errors"
	"main/config"
	"main/domain/repository"
)

type genericRepository struct {
	dbClients map[string]config.DBClient
}

func NewGenericRepository(clients map[string]config.DBClient) repository.GenericRepository {
	return &genericRepository{dbClients: clients}
}

func (g genericRepository) ExecuteQuery(scope string, query string, params map[string]string) ([]map[string]interface{}, error) {
	if g.dbClients[scope].StringConnection == "" {
		return nil, errors.New("scope not implemented yet")
	}

	return executeQuery(g.dbClients[scope], query, params)
}
