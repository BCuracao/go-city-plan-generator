package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Open the OSM file
	osmFile, err := os.Open("map.osm")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File successfully opened")
	// Defer the closing of our OSM file so that we can parse it later on
	defer osmFile.Close()

	// Read the OSM file as a byte array
	byteValues, _ := ioutil.ReadAll(osmFile)

}
