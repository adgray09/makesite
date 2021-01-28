package main

import "io/ioutil"

func main() {

	bytesToWrite := []byte("hello\ngo\n")
	err := ioutil.WriteFile("new-file.txt", bytesToWrite, 0644)
	if err != nil {
		panic(err)
	}
	
}
