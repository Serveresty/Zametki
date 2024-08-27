package models

type SpellError struct {
	Pos  int      `json:"pos"`
	Len  int      `json:"len"`
	Word string   `json:"word"`
	S    []string `json:"s"`
}
