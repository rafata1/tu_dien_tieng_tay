package core

import (
	"github.com/jmoiron/sqlx"
	"github.com/templateOfService/connectors/mysql"
)

type Repo struct {
	conn *sqlx.DB
}

func NewRepo() *Repo {
	return &Repo{
		conn: mysql.GetMySQLInstance(),
	}
}
