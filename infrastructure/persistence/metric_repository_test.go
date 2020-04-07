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

func TestGetMetric(t *testing.T) {
	cases := []struct {
		input        int
		expectMetric *domain.Metric
		expectErr    error
		mock         func(m sqlmock.Sqlmock)
	}{
		{
			1,
			&domain.Metric{
				ID:   2,
				Name: "metric1",
			},
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery("SELECT \\* FROM `metricas`").
					WillReturnRows(sqlmock.NewRows([]string{"id", "nombre"}).
						AddRow("1", "metric1").AddRow("2", "metric1"))
			},
		},
		{
			0,
			nil,
			errors.New("record not found"),
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery("SELECT \\* FROM `metricas`").
					WillReturnRows(sqlmock.NewRows([]string{""})).
					WillReturnError(errors.New("record not found"))
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
		repo := NewMetricRepository(dbClient)
		metric, err := repo.Get(c.input)
		if !reflect.DeepEqual(err, c.expectErr) {
			t.Fatalf("#%d: want error %#v, got %#v", i, c.expectErr, err)
		}
		if err != nil {
			continue
		}
		if !reflect.DeepEqual(metric, c.expectMetric) {
			t.Errorf("#%d: want %#v, got %#v", i, c.expectMetric, metric)
		}
	}
}
