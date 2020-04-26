package application

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TesMakeDate(t *testing.T) {
/*
	data := []map[string]interface{}{{"id": "1", "nombre": "metric1"}, {"id": "2", "nombre": "metric2"}}
	expected := MetricResponse{
		Graph:       nil,
		Data:        nil,
		TextToHuman: "",
	}

	result, _ := makeData(data)

	assert.Equal(t, expected, result)*/
}

func TestTextToHuman(t *testing.T) {
	expected := "Te muestro el historial de unidades cargadas en los últimos 12 meses:\\n\\nCampaña Enero 2019: 1391 unidades\\nCampaña Febrero 2019: 1067 unidades\\n"
	param := map[string]string{"id_campania": "1179", "id_vendedor": "PAZSAN48","titulox":"campaña","tituloy":"cantidad"}
	template := "Te muestro el historial de unidades cargadas en los últimos 12 meses:\n\n[repeat]Campaña :x: :y unidades\n[/repeat]"
	data := []map[string]interface{}{{"x": "Enero", "y": "10000"}, {"x": "Febrero", "y": "20000"}}
	result := textToHuman(template,param, data)

	assert.Equal(t, expected, result)
}