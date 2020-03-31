package repository

import (
	"context"
	"main/domain"
)

// MetricRepository represent repository of the metric
// Expect implementation by the infrastructure layer
type MetricRepository interface {
	Get(ctx context.Context, id int) (*domain.Metric, error)
}