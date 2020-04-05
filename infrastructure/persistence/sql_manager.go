package persistence

import (
	"main/config"
	"strings"
)

func executeQuery(client config.DBClient, query string, params map[string]string) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	db, err := client.GetConnection(client.Dialect, client.StringConnection)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	preparedQuery := getQuery(query, params)
	rows, err := db.Raw(preparedQuery).Rows()
	if err != nil {
		return nil, err
	}

	cols, _ := rows.Columns()

	for rows.Next() {
		row := make(map[string]interface{})
		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		rows.Scan(columnPointers...)

		for i, colName := range cols {
			row[colName] = columns[i]
		}

		result = append(result, row)
	}

	return result, nil
}

func getQuery(query string, params map[string]string) string {
	for key, value := range params {
		if strings.Contains(query, key) {
			query = strings.Replace(query, key, "'"+value+"'", -1)
		}
	}

	return query
}
