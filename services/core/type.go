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
	Word        string       `json:"word"`
	Type        string       `json:"type"`
	Definitions []Definition `json:"definitions"`
}

type Definition struct {
	Meaning  string   `json:"meaning"`
	Examples []string `json:"example"`
}
