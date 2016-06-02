package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func readFile(file io.Reader) {
	scanner := bufio.NewScanner(file) // returns a new Scanner
	for scanner.Scan() {              // iterates through the text file until the end, whereby it returns false
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil { // if something happens, that is not related to EOF then the error is printed.
		log.Fatal(err)
	}
}

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err) // if the file does not exist or can't be opened an error msg will be printed
	}
	return file
}

func main() {

	filename := "updatingDocuments.js"

	file := openFile(filename) // Opens specified file

	defer file.Close() // close the file after everything has been executed, placed here so not forgotten

	readFile(file) //reads specified file
}
