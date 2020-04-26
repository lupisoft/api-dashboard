package application

import (
	"github.com/fatih/structs"
	"main/application/graph"
	"main/domain"
	"main/domain/repository"
	"strings"
)

type histogramInteractor struct {
	genericRepository repository.GenericRepository
}

func NewHistogramInteractor(repository repository.GenericRepository) TypeMetricInteractor {
	return &histogramInteractor{genericRepository: repository}
}

func (h histogramInteractor) Execute(metricRequest MetricRequest) (*MetricResponse, error) {
	result, err := h.genericRepository.ExecuteQuery(metricRequest.Scope, metricRequest.Query, metricRequest.Params)
	if err != nil {
		return nil, err
	}

	res, err := makeData(metricRequest.Params, metricRequest.Metric, result)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func makeData(paramRequest map[string]string, metric domain.Metric, resultQuery []map[string]interface{}) (*MetricResponse, error) {
	internalParamsMap := metric.GetMapInternalParam()
	titleX := internalParamsMap["titulox"]
	titleY := internalParamsMap["tituloy"]

	addMap(paramRequest, internalParamsMap)

	histogramGraph := graph.NewHistogramGraph(metric.Title, metric.Subtitle, metric.Description, titleX, titleY, resultQuery)
	text := textToHuman(metric.Template, paramRequest, resultQuery) //$this->textToHuman($this->template, $allparams, $data)
	result := &MetricResponse{Graph: structs.Map(histogramGraph), Data: resultQuery, TextToHuman: text}

	return result, nil
}

func textToHuman(template string, param map[string]string, data []map[string]interface{}) string {

	splittedText := strings.Split(template, "[repeat]")
	textReplace := strings.Replace(splittedText[1],"[/repeat]","",-1)
	var textRepeat string
	var textHuman string

	for i := 0; i < len(data); i ++ {
		text := strings.Replace(textReplace,":x",data[i]["x"].(string),-1)
		text = strings.Replace(text,":y",data[i]["y"].(string),-1)

		textRepeat = textRepeat + text
	}

	textHuman = splittedText[0] + textRepeat

	return textHuman
}

func addMap(a map[string]string, b map[string]string) {
	for k, v := range b {
		a[k] = v
	}
}
