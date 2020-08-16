package main

import (
	"fmt"
	"time"
	"os"
	"io"
	"bufio"
	"log"
	"encoding/csv"
	"strings"
	"strconv"
)

func readFile (name string) *csv.Reader {
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}

	fileInfo, err := os.Lstat(name)  // get length in bytes
	if err != nil {
		log.Fatal(err)
	}
	
	b := make([]byte, fileInfo.Size()) // allocate byte array

	f.Read(b) // read file into byte arrray
	if err != nil {
		log.Fatal(err)
	}

	
	s := string(b) //convert to string

	return csv.NewReader(strings.NewReader(s))
}

func ask (question string) {
	fmt.Printf("%v=\n", question)
}



func main () {

	var count int
	filename := "problems.csv"

	args := os.Args[1:]
	if args[0] == "-f" {
		filename = args[1]
	}

	lines := readFile(filename)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Press any key to start")
	reader.ReadString('\n')
	start := time.Now()
	for {
		t := time.Now()
		elapsed := t.Sub(start)
		// sec := time.ParseDuration(elapsed)
		// fmt.Printf("%v", elapsed.Seconds())
		if elapsed.Seconds() > 30 {
			fmt.Println("Time's up!")
			break
		}
		record, err := lines.Read()
		if err == io.EOF {
			break
		}
		// fmt.Printf("%q/n", strings.Split(fmt.Printf("%q", record), " "))
		question := record[0]
		ask(question)
		clianswer, _ := reader.ReadString('\n')
		answer := strings.TrimSuffix(clianswer, "\n")
		intUserAnswer, _ := strconv.Atoi(answer)
		intAnswer, _ := strconv.Atoi(record[1])
		fmt.Printf("%v, %v\n", answer, intAnswer)
		if intUserAnswer == intAnswer {
			count++
		}
		// answer := record[1]
	}

	fmt.Printf("\nYou got %v right", count)

}