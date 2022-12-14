package models

type Word struct {
	ID            int    `db:"id"`
	Word          string `db:"word"`
	Type          string `db:"type"`
	Language      string `db:"language"`
	Pronunciation string `db:"pronunciation"`

	Definitions           []Definition
	NormalizedWord        string
	NormalizedDefinitions string
}
