package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/mgo.v2/bson"
)

type movie struct { //exported Student
	ID    float64 `json:"id"`
	Title string  `json:"title"`
	Year  float64 `json:"year"`
	Typ   string  `json:"type"`
}

func main() {
	filebytes, err := ioutil.ReadFile("pages.json")
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	var M []movie
	json.Unmarshal(filebytes, &M)

	for _, m := range M {
		Mbytes, err := bson.Marshal(m)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println(string(Mbytes))
	}

}
