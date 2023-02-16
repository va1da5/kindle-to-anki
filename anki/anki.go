package anki

import (
	"io"
	"kindle-flashcards/entity"
	"log"
	"strings"
	"text/template"
)

const tmpl string = `#separator:pipe
#html:true{{range $word := .}}
<h4>{{$word.Word}}</h4>{{range $usage := $word.Usage}}<p>{{$usage}}</p>{{end}}|{{range $definition := $word.Definitions}}<p><i>{{$definition.PartOfSpeech}}<i></p><ul>{{range $explanation := $definition.Definitions }}<li>{{$explanation.Definition | remove_newline}}</li>{{end}}</ul>{{end}}</br>{{range $book := $word.Books}}<p><strong>{{$book.Book}}</strong> - <i>{{$book.Authors}}</i></p>{{end}}{{end}}
`

func replace(input, from, to string) string {
	return strings.Replace(input, from, to, -1)
}

func remove_newline(input string) string {
	return strings.Replace(input, "\n", " ", -1)
}

var funcMap = template.FuncMap{
	"replace":        replace,
	"remove_newline": remove_newline,
}

func GenerateCards(words []entity.Word, wr io.Writer) {
	t := template.Must(template.New("tmpl").Funcs(funcMap).Parse(tmpl))

	err := t.Execute(wr, words)
	if err != nil {
		log.Fatalln(err)
	}
}
