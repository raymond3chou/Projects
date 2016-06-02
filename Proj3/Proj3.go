package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Student struct { //exported Student
	Name  string
	Phone string
}

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

func mongoconnect() {
	session, err := mgo.Dial("localhost") //attempts a connection and creates a session
	if err != nil {
		panic(err)
	}
	defer session.Close()
	connection := session.DB("test").C("students") //name of DB and Collection
	err = connection.Insert(&Student{"Ale", "+55 53 8116 9639"})
	if err != nil {
		log.Fatal(err)
	}
	result := Student{}
	err = connection.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Phone:", result.Phone)
}
func main() {

	filename := "updatingDocuments.js"

	file := openFile(filename) // Opens specified file

	defer file.Close() // close the file after everything has been executed, placed here so not forgotten
	mongoconnect()     //establish a connection to MongoDB on localhost
	readFile(file)     //reads specified file
}
