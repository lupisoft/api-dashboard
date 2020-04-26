package application

import (
	"main/domain"
	"main/domain/repository"
)

type MetricInteractor struct {
	MetricRepository repository.MetricRepository
}

func (m MetricInteractor) Get(id int) (*domain.Metric, error) {
	return m.MetricRepository.Get(id)
}
