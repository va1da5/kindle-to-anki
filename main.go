package main

import (
	"bytes"
	"kindle-flashcards/anki"
	"kindle-flashcards/entity"
	"kindle-flashcards/kindle"
	"kindle-flashcards/wiktionary"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const database string = "vocab.db"

func main() {

	db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var words []entity.Word
	var kindleWords []kindle.Word
	var lookups []kindle.Lookup

	db.Find(&kindleWords)

	for _, word := range kindleWords {
		usage := []string{}
		books := []entity.Book{}

		definitions := wiktionary.GetDefinitions(word.Stem)

		if definitions == nil {
			continue
		}

		db.Preload("BookInfo").Find(&lookups, "word_key = ?", word.Id)
		for _, lookup := range lookups {
			usage = append(usage, lookup.Usage)
			books = append(books, entity.Book{
				Book:    lookup.BookInfo.Title,
				Authors: lookup.BookInfo.Authors,
			})
		}

		words = append(words, entity.Word{
			Word:        word.Stem,
			Usage:       usage,
			Books:       books,
			Definitions: definitions,
		})

	}

	buf := new(bytes.Buffer)

	anki.GenerateCards(words, buf)

	err = os.WriteFile("anki-flashcards.txt", buf.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
