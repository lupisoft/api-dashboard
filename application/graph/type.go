package graph

type HistogramGraph struct {
	Title       string
	Subtitle    string
	Description string
	TitleX      string
	TitleY      string
	X           []interface{}
	Y           []interface{}
}

func NewHistogramGraph(title string, subtitle string, description string, titleX string, titleY string, data []map[string]interface{}) HistogramGraph {
	var x []interface{}
	var y []interface{}
	for _, value := range data {
		x = append(x, value["x"])
		y = append(y, value["y"])
	}

	return HistogramGraph{
		Title:       title,
		Subtitle:    subtitle,
		Description: description,
		TitleX:      titleX,
		TitleY:      titleY,
		X:           x,
		Y:           y,
	}
}
