package main

import (
	"bytes"
	"flag"
	"html/template"
	"io/ioutil"
	"path/filepath"
)

type words struct {
	Content string
}

func main() {

	// bytesToWrite := []byte("hello\ngo\n")
	// err := ioutil.WriteFile("new-file.txt", bytesToWrite, 0644)
	// if err != nil {
	// 	panic(err)
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

	// fptr := flag.String("fpath", "test1.txt", "file path to read from")
	// flag.Parse()
	// fmt.Println("value pf fpath is", *fptr)

	// files, err := ioutil.ReadDir(".")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// find all files inside directory that end with .txt

	// prints all matches
	if directory != "" {
		allTextFilesInDirectory, _ := filepath.Glob(directory + "/*.txt")
		for _, file := range allTextFilesInDirectory {
			writeFile(file)
		}
	}

	// var dir string

	// flagUsage1 := "finds all text files"
	// flag.StringVar(&dir, "dir", matches, flagUsage1)
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

	err = ioutil.WriteFile(fileName[0:len(fileName)-4]+".html", b.Bytes(), 0644)
	if err != nil {
		panic(err)
	}

}
