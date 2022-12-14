package models

type Definition struct {
	WordID   int    `db:"word_id"`
	Meaning  string `db:"meaning"`
	Examples string `db:"examples"`
}
