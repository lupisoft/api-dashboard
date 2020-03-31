package domain

type Metric struct {
	ID             int         `gorm:"primary_key;column:id"`
	TypeMetric     TypeMetric  `gorm:"foreignkey:TypeMetricID"`
	TypeMetricID   int         `gorm:"column:id_tipo_metrica"`
	Subcategory    Subcategory `gorm:"foreignkey:SubcategoryID"`
	SubcategoryID  int         `gorm:"column:id_subcategoria"`
	Name           string      `gorm:"column:nombre"`
	Title          string      `gorm:"column:titulo"`
	Subtitle       string      `gorm:"column:subtitulo"`
	Description    string      `gorm:"column:descripcion"`
	Icon           string      `gorm:"column:icon"`
	Url            string      `gorm:"column:url"`
	Scope          string      `gorm:"column:ambito"`
	Query          string      `gorm:"column:query"`
	Params         string      `gorm:"column:parametros"`
	InternalParams string      `gorm:"column:parametros_internos"`
	Template       string      `gorm:"column:template"`
	Status         string      `gorm:"column:activo"`
}

type TypeMetric struct {
	ID   int    `gorm:"primary_key;column:id"`
	Name string `gorm:"column:nombre"`
}

type Subcategory struct {
	ID         int      `gorm:"primary_key;column:id"`
	Category   Category `gorm:"foreignkey:CategoryID"`
	CategoryID int      `gorm:"column:id_categoria"`
	Name       string   `gorm:"column:nombre"`
	Status     int      `gorm:"column:activo"`
}

type Category struct {
	ID     int    `gorm:"primary_key;column:id"`
	Name   string `gorm:"column:nombre"`
	Icon   string `gorm:"column:icon"`
	Status int    `gorm:"column:activo"`
}

func (Metric) TableName() string {
	return "metricas"
}

func (TypeMetric) TableName() string {
	return "tipos_metricas"
}

func (Subcategory) TableName() string {
	return "subcategorias"
}

func (Category) TableName() string {
	return "categorias"
}
