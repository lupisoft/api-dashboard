package repository

type GenericRepository interface {
	ExecuteQuery(scope string, query string, params map[string]string) ([]map[string]interface{}, error)
}
