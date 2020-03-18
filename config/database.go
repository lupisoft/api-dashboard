package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDBConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "lupisoft:Agus17533542.@/ddd_sample?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Error connection to database !!")
	}

	return db
}
