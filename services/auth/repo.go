package auth

import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
    "github.com/templateOfService/connectors/mysql"
    "github.com/templateOfService/models"
    "log"
)

type Repo struct {
    conn *sqlx.DB
}

func NewRepo() *Repo {
    return &Repo{
        conn: mysql.GetMySQLInstance(),
    }
}

func (r *Repo) InsertUser(user models.User) error {
    query := `INSERT INTO users(username, email, password, phone) VALUE(:username, :email, :password, :phone)`
    _, err := r.conn.NamedExec(query, user)
    if err != nil {
        log.Printf("Error inserting user: %s", err.Error())
    }
    return err
}
