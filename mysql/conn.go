package mysql

import (
	"fmt"

	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB(conn_str string) *gorm.DB {
	db, err := gorm.Open(gmysql.Open(conn_str), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	return db
}
