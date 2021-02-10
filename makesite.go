package main

import (
	"bytes"
	"context"
	"flag"
	"html/template"
	"io/ioutil"
	"path/filepath"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

type words struct {
	Content string
}

func main() {

	var fileName, directory string // both are nil
	flag.StringVar(&fileName, "file", "", "text file to be rendered as html")
	flag.StringVar(&directory, "dir", "", "directoy to look for text files in")
	flag.Parse()

	if fileName != "" {
		_, err := ioutil.ReadFile(fileName)
		if err != nil {
			panic(err)
		} else {
			writeFile(fileName)
		}
	}

	// prints all matches
	if directory != "" {
		allTextFilesInDirectory, _ := filepath.Glob(directory + "/*.txt")
		for _, file := range allTextFilesInDirectory {
			writeFile(file)
		}
	}
}

func writeFile(fileName string) {
	content, err := ioutil.ReadFile(fileName)
	fileContent := words{string(content)}

	var translated string

	translated = translateText(language.Afrikaans.String(), fileContent.Content)
	fileContent = words{string(translated)}

	var b bytes.Buffer

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	err = t.Execute(&b, fileContent)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(fileName[0:len(fileName)-4]+".html", b.Bytes(), 0644)
	if err != nil {
		panic(err)
	}

}

func translateText(targetLanguage, text string) string {
	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		panic(err)
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		panic(err)
	}

	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		panic(err)
	}

	if len(resp) == 0 {
		panic(err)
	}
	return resp[0].Text

}
