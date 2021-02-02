package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	// "flag"
	// "fmt"
)

func main() {

	// bytesToWrite := []byte("hello\ngo\n")
	// err := ioutil.WriteFile("new-file.txt", bytesToWrite, 0644)
	// if err != nil {
	// 	panic(err)
	// }
	content, err := ioutil.ReadFile("first-post.txt")

	type words struct {
		Content string
	}

	var b bytes.Buffer

	fileContent := words{string(content)}

	

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	err = t.Execute(&b, fileContent)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("latest-post.html", b.Bytes(), 777)
	if err != nil {
		panic(err)
	}
}
