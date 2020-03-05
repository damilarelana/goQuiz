package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func errMsgHandler(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {

	// define flags
	csvFilename := flag.String("csv", "quizData.csv", "a csv file containing question/answer data, in a 'question, answer' format per record line")

	// parse flags
	flag.Parse() // required to initialize the specified flags with the Operating system

	// open CSV file and read the records (i.e. all the lines) from it
	openedFile, err := os.Open(*csvFilename)
	if err != nil {
		errMsgHandler(fmt.Sprintf("Failed to open CSV file: %s\n", *csvFilename))
	}
	r := csv.NewReader(openedFile)
	records, err := r.ReadAll()
	if err != nil {
		errMsgHandler("Failed to parse the provided CSV file")
	}
	fmt.Println(records)
}
