package main

import (
	"fmt"
	"os"
	"io"
	"log"
	"encoding/csv"
	"strings"
)



func main () {

	// open a file
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	// get length in bytes
	fileInfo, err := os.Lstat("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("size: ", fileInfo.Size())

	// allocate byte array
	b := make([]byte, fileInfo.Size())

	// read file into byte arrray
	r, err := f.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)

	//convert to string
	s := string(b)
	fmt.Println(s)

	lines := csv.NewReader(strings.NewReader(s))


	for {
		record, err := lines.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}

	args := os.Args[1:]
	fmt.Println(args)

}