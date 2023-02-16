package wiktionary

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const url string = "https://en.wiktionary.org/api/rest_v1/page/definition/%s"

type Definition struct {
	PartOfSpeech string `json:"partOfSpeech"`
	Language     string `json:"language"`
	Definitions  []struct {
		Definition string `json:"definition"`
	} `json:"definitions"`
}

type EnglishDefinition struct {
	English []Definition `json:"en"`
}

func GetDefinitions(word string) []Definition {
	res, err := http.Get(fmt.Sprintf(url, word))
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	var response EnglishDefinition

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Fatal(err)
	}

	return response.English
}
