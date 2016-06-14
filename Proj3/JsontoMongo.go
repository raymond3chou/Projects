package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"log"
	"os"

	"gopkg.in/mgo.v2"
)

type movie struct { //exported Student
	ID    float64 `json:"id"`
	Title string  `json:"title"`
	Year  float64 `json:"year"`
	Type  string  `json:"type"`
}

func fileRead(filename string) []movie {
	filebytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	var Mfile []movie
	json.Unmarshal(filebytes, &Mfile)
	return Mfile
}

func mongoinsert(mfile []movie) bool {
	session, err := mgo.Dial("localhost") //attempts a connection and creates a session
	if err != nil {
		panic(err)
	}

	defer session.Close()
	connection := session.DB("test").C("movies") //name of DB and Collection
	// need to seperate JSON into a map/struct somehow?
	for _, info := range mfile {
		err := connection.Insert(info)
		if err != nil {
			log.Fatal(err)
			return false
		}
	}
	return true
}

func main() {

	Mfile := fileRead("pages.json")

	fmt.Println(mongoinsert(Mfile))
}
