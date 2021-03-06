package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DBClient struct {
	StringConnection string
	Dialect          string
	GetConnection    func(dialect string, connLine string) (*gorm.DB, error)
}

func NewDBClientBuilder(dbConfig DataBase) DBClient {
	return DBClient{
		StringConnection: dbConfig.StringConnection,
		Dialect:          dbConfig.Dialect,
		GetConnection: func(dialect string, connLine string) (db *gorm.DB, err error) {
			dbGorm, err := gorm.Open(dialect, connLine)
			if err != nil {
				return nil, err
			}
			dbGorm.DB().SetMaxIdleConns(20)
			dbGorm.DB().SetMaxOpenConns(40)
			dbGorm.SingularTable(true)
			dbGorm.LogMode(true)
			return dbGorm, nil
		},
	}
}

func NewDBClientBuilderMock(getConnection func(dialect string, connLine string) (*gorm.DB, error)) DBClient {
	return DBClient{
		StringConnection: "mock",
		Dialect:          "mock",
		GetConnection:    getConnection,
	}
}
