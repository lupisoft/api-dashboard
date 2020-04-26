package application

import "main/domain"

type TypeMetricInteractor interface {
	Execute(metricRequest MetricRequest) (*MetricResponse, error)
}

type MetricResponse struct {
	Graph       map[string]interface{}
	Data        []map[string]interface{}
	TextToHuman string
}

type MetricRequest struct {
	Scope  string
	Query  string
	Params map[string]string
	Metric domain.Metric
}
