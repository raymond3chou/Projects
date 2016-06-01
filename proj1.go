package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("proj1.txt")
	if err != nil {
		log.Fatal(err) // if the file does not exist or can't be opened an error msg will be printed
	}
	defer file.Close() // close the file after everything has been executed, placed here so not forgotten

	scanner := bufio.NewScanner(file) // returns a new Scanner
	for scanner.Scan() {              // iterates through the text file until the end, whereby it returns false
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil { // if something happens, that is not related to EOF then the error is printed.
		log.Fatal(err)
	}
}
