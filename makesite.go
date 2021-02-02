package main

import (
	"bytes"
	"flag"
	"html/template"
	"io/ioutil"
	// "fmt"
)

type words struct {
	Content string
}

func main() {

	// bytesToWrite := []byte("hello\ngo\n")
	// err := ioutil.WriteFile("new-file.txt", bytesToWrite, 0644)
	// if err != nil {
	// 	panic(err)
	var fileName string
	flagUsage := "text File to be rendered as html"
	flag.StringVar(&fileName, "file", "first-post.txt", flagUsage)
	writeFile(fileName)
	// }

}

func writeFile(fileName string) {
	content, err := ioutil.ReadFile(fileName)
	fileContent := words{string(content)}
	var b bytes.Buffer

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	err = t.Execute(&b, fileContent)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(fileName[0:len(fileName)-4]+".html", b.Bytes(), 777)
	if err != nil {
		panic(err)
	}

}
