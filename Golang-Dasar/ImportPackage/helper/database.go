package helper

import "fmt"

var connect string

func init() {
	fmt.Println("Connection tetap terpanggil (Blank Identifier")
	connect = "MySQL"
}

func GetDataBase() string {
	return connect
}
