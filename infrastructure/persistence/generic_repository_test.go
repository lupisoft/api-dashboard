package persistence

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"main/config"
	"testing"
)

func TestGenericRepository_ExecuteQuery(t *testing.T) {
	clients := map[string]config.DBClient{"datamaaar": config.DBClient{}}

	repository := NewGenericRepository(clients)

	_, err := repository.ExecuteQuery("datamar", "mockQuery", map[string]string{})

	assert.Error(t, err)
	assert.Equal(t, errors.New("scope not implemented yet"), err)
}
