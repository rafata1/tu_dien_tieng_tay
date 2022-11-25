package core

type BaseRes struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var SuccessMessage = "Success"
var ErrInternal = "Internal error"

type SearchRes struct {
	Words []WordRes `json:"words"`
}

type WordRes struct {
	ID                 int          `json:"id"`
	Word               string       `json:"word"`
	Type               string       `json:"type"`
	PronunciationSound string       `json:"pronunciation_sound"`
	Definitions        []Definition `json:"definitions"`
}

type Definition struct {
	Meaning  string   `json:"meaning"`
	Examples []string `json:"example"`
}

type AddWord struct {
	Word        string       `json:"word"`
	Type        string       `json:"type"`
	Language    string       `json:"language"`
	Definitions []Definition `json:"definitions"`
}

type AddWordRes struct {
	ID int `json:"id"`
}
