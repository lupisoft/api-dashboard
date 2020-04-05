package persistence

import (
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
	configDB := config.DataBase{Dialect: "mysql", StringConnection: "root@/dashboard?charset=utf8&parseTime=True&loc=Local"}
	dbClient := config.NewDBClientBuilder(configDB)
	params := make(map[string]string)
	params[":id_metric"] = "1"

	result, _ := executeQuery(dbClient , "select * from metricas where id >= :id_metric", params)

	assert.Equal(t, 1, len(result))
}
