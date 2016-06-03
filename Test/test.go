package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type movie struct { //exported Student
	ID    float64 `json:"id"`
	Title string  `json:"title"`
	URL   string  `json:"url"`
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
		Mbytes, err := json.Marshal(m)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println(string(Mbytes))
	}

}
