package dbreverse

import "github.com/liuyong-go/dbreverse/mysql"

func Create(driver string, connStr string, path string, schema string, tables ...string) {
	switch driver {
	case "mysql":
		mysql.NewGenerator(connStr, path, schema).Genertate(tables...)
	default:
		mysql.NewGenerator(connStr, path, schema).Genertate(tables...)
	}

}
