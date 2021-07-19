package main

import "github.com/liuyong-go/dbreverse"

func main() {
	var tables = []string{"users", "roles", "apis"}
	dbreverse.Create("mysql", "root:123@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local", "./", "test", tables...)
}
