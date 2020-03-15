package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

// defines the error message handler
func errMsgHandler(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

// defines the structure of the `problem` type
type problem struct {
	question string
	answer   string
}

//
// define parser to
//   * read in the multi-dimensional slice of `question, answer` i.e. CSV file data
//   * and then return a 1 dimensional slice of `question, answer` i.e. a value of type `problem`
//
func parseRecords(records [][]string) []problem {

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
