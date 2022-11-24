package mysql

import (
    "github.com/jmoiron/sqlx"
    "os"
)

var conn *sqlx.DB

func Connect() error {
    var err error
    conn, err = sqlx.Connect("mysql", os.Getenv("MYSQL"))
    return err
}

func GetMySQLInstance() *sqlx.DB {
    return conn
}
