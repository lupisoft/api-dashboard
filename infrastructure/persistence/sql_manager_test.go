package persistence

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"main/config"
	"testing"
)

func TestGetQuery(t *testing.T) {
	query := "SELECT COALESCE(SUM(comp_pedidos.preciounidad * comp_pedidos.cantidad),0) as valor " +
		"FROM comp_pedidos " +
		"JOIN pedido ON comp_pedidos.id_pedido = pedido.id_pedidos " +
		"WHERE comp_pedidos.id_vendedor LIKE :id_vendedor " +
		"AND pedido.id_campania = :id_campania"
	params := make(map[string]string)
	params[":id_vendedor"] = "1"
	params[":id_campania"] = "2"

	expectedQuery := "SELECT COALESCE(SUM(comp_pedidos.preciounidad * comp_pedidos.cantidad),0) as valor FROM comp_pedidos JOIN pedido ON comp_pedidos.id_pedido = pedido.id_pedidos WHERE comp_pedidos.id_vendedor LIKE '1' AND pedido.id_campania = '2'"

	result := getQuery(query, params)

	assert.Equal(t, expectedQuery, result)
}

func TestExecuteQuery(t *testing.T) {
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
		dbGorm.LogMode(true)
		dbGorm.SingularTable(true)
		return dbGorm, nil
	}
	dbClient := config.NewDBClientBuilderMock(getConnection)

	expectedResult := []map[string]interface {}{{"id":"1", "nombre":"metric1"},{"id":"2", "nombre":"metric2"}}
	mock.ExpectQuery("select  \\* from metricas where id \\>\\= '1'").
		WillReturnRows(sqlmock.NewRows([]string{"id", "nombre"}).
			AddRow("1", "metric1").
			AddRow("2", "metric2"))

	params := make(map[string]string)
	params[":id_metric"] = "1"

	result, err := executeQuery(dbClient, "select * from metricas where id >= :id_metric", params)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, expectedResult, result)
}
