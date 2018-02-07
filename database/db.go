package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DBCon *sql.DB
var err error

func Init() {
	DBCon, err = sql.Open("mysql", "root:@/ideaparLog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	return
}

