package repository

import (
	"main/domain"
)

// MetricRepository represent repository of the metric
// Expect implementation by the infrastructure layer
type MetricRepository interface {
	Get(id int) (*domain.Metric, error)
}