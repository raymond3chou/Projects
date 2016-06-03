package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/mgo.v2"
)

type Movie struct { //exported Student
	ID    float64 `json:"id"`
	Title string  `json:"title"`
	Year  float64 `json:"year"`
	Type  string  `json:"type"`
}

func editFile() []Movie { //opens file and the reads file. Returns as Movie Array
	filebytes, err := ioutil.ReadFile("pages.json")
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	var Mfile []Movie
	json.Unmarshal(filebytes, &Mfile)
	return (Mfile)
}

func mongoconnect(Mfile []Movie) {
	session, err := mgo.Dial("localhost") //attempts a connection and creates a session
	if err != nil {
		panic(err)
	}
	defer session.Close()
	connection := session.DB("test").C("movies") //name of DB and Collection

	for i := range Mfile {
		err = connection.Insert(Mfile[i])
		if err != nil {
			log.Fatal(err)
		}
	}
}
func main() {

	Mfile := editFile()
	mongoconnect(Mfile)
}
