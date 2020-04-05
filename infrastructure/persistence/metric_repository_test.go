package persistence

import (
	"context"
	"errors"
	"main/config"
	"main/domain"
	"reflect"
	"testing"
)

func TestGetUser(t *testing.T) {
	configDB := config.DataBase{Dialect: "mysql", StringConnection: "root@/dashboard?charset=utf8&parseTime=True&loc=Local"}
	dbClient := config.NewDBClientBuilder(configDB)

	cases := []struct {
		input      int
		expectMetric *domain.Metric
		expectErr  error
	}{
		{
			1,
			&domain.Metric{
				ID:   2,
				Name: "metric1",
			},
			nil,
		},
		{
			0,
			nil,
			errors.New("record not found"),
		},
	}
	for i, c := range cases {
		repo := NewMetricRepository(dbClient)
		user, err := repo.Get(context.Background(), c.input)
		if err != c.expectErr {
			t.Fatalf("#%d: want error %#v, got %#v", i, c.expectErr, err)
		}
		if err != nil {
			continue
		}
		if !reflect.DeepEqual(user, c.expectMetric) {
			t.Errorf("#%d: want %#v, got %#v", i, c.expectMetric, user)
		}
	}
}