package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var conn *sqlx.DB

func Connect() error {
	var err error
	conn, err = sqlx.Connect("mysql", "root:1@(localhost:3306)/dictionary")
	return err
}

func GetMySQLInstance() *sqlx.DB {
	return conn
}
