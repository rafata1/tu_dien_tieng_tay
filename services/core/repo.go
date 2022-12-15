package core

import (
	"github.com/jmoiron/sqlx"
	"github.com/templateOfService/connectors/mysql"
	"github.com/templateOfService/models"
)

type Repo struct {
	conn *sqlx.DB
}

func NewRepo() *Repo {
	return &Repo{
		conn: mysql.GetMySQLInstance(),
	}
}

func (r *Repo) InsertWord(word models.Word) (int, error) {
	query := "INSERT INTO words (word, type, language) VALUES (:word, :type, :language)"
	res, err := r.conn.NamedExec(query, word)
	if err != nil {
		return 0, err
	}
	id, _ := res.LastInsertId()
	return int(id), nil
}

func (r *Repo) UpdateWord(word models.Word) error {
	query := "UPDATE words SET word = ?, type  = ?, language = ? WHERE id = ?"
	_, err := r.conn.Exec(query, word.Word, word.Type, word.Language, word.ID)
	return err
}

func (r *Repo) ModifyDefinitions(wordID int, definitions []models.Definition) error {
	query := "DELETE FROM definitions WHERE word_id = ?"
	_, err := r.conn.Exec(query, wordID)
	if err != nil {
		return err
	}

	query = "INSERT INTO definitions (word_id, meaning, examples) VALUES (:word_id, :meaning, :examples)"
	_, err = r.conn.NamedExec(query, definitions)
	return err
}

func (r *Repo) GetAllWords() ([]models.Word, error) {
	var res []models.Word
	query := "SELECT * FROM words"
	err := r.conn.Select(&res, query)
	return res, err
}

func (r *Repo) GetAllDefinitions() ([]models.Definition, error) {
	var res []models.Definition
	query := "SELECT * FROM definitions"
	err := r.conn.Select(&res, query)
	return res, err
}

func (r *Repo) UpdatePronounce(id int, file string) error {
	query := "UPDATE words SET pronunciation = ? WHERE id = ?"
	_, err := r.conn.Exec(query, file, id)
	return err
}
