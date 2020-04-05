package persistence

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"main/config"
	"main/domain"
	"reflect"
	"testing"
)

func TestGetAllCategories(t *testing.T) {
	cases := []struct {
		expectCategories []*domain.Category
		expectErr        error
		mock             func(m sqlmock.Sqlmock)
	}{
		{
			[]*domain.Category{{
				ID:   2,
				Name: "metric1",
			}, {
				ID:   2,
				Name: "metric1",
			},
			},
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery("SELECT \\* FROM `categorias`").WillReturnRows(sqlmock.NewRows([]string{"id", "nombre"}).AddRow("2", "metric1").AddRow("2", "metric1"))
			},
		},
		{
			nil,
			errors.New("record not found"),
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery("SELECT \\* FROM `categorias`").WillReturnRows(sqlmock.NewRows([]string{""})).WillReturnError(errors.New("record not found"))
			},
		},
	}
	for i, c := range cases {
		sqlDb, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		getConnection := func(dialect string, connLine string) (db *gorm.DB, err error) {
			dbGorm, err := gorm.Open("mysql", sqlDb)
			if err != nil {
				panic("failed to connect database")
			}
			dbGorm.DB().SetMaxIdleConns(25)
			dbGorm.DB().SetMaxOpenConns(50)
			dbGorm.SingularTable(true)
			return dbGorm, nil
		}
		dbClient := config.NewDBClientBuilderMock(getConnection)

		c.mock(mock)
		repo := NewCategoryRepository(dbClient)
		result, err := repo.GetAll()
		if !reflect.DeepEqual(err, c.expectErr) {
			t.Fatalf("#%d: want error %#v, got %#v", i, c.expectErr, err)
		}

		if err != nil {
			continue
		}

		if !reflect.DeepEqual(result, c.expectCategories) {
			t.Errorf("#%d: want %#v, got %#v", i, c.expectCategories, result)
		}
	}
}
