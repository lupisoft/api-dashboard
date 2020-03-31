package persistence

import (
	"context"
	"main/config"
	"main/domain"
	"main/domain/repository"
)

type metricRepository struct {
	dbClient config.DBClient
}

func NewMetricRepository(client config.DBClient) repository.MetricRepository {
	return &metricRepository{dbClient: client}
}

func (m metricRepository) Get(ctx context.Context, id int) (*domain.Metric, error) {
	var metric domain.Metric
	db, err := m.dbClient.GetConnection(m.dbClient.Dialect, m.dbClient.StringConnection)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err := db.Preload("TypeMetric").
		Preload("Subcategory").
		Preload("Subcategory.Category").
		Where("id = ?", id).
		First(&metric).Error; err != nil {
		return nil, err
	}

	return &metric, nil
}
