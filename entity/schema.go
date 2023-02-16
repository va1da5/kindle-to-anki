package entity

import "kindle-flashcards/wiktionary"

type Book struct {
	Book    string `json:"book"`
	Authors string `json:"authors"`
}

type Word struct {
	Word        string                  `json:"word"`
	Usage       []string                `json:"usage"`
	Books       []Book                  `json:"books"`
	Definitions []wiktionary.Definition `json:"definitions"`
}
