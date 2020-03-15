package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"
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
	returnedValue := make([]problem, len(records))
	for i, record := range records { // iterate over the multi-dimensional slice
		returnedValue[i] = problem{
			question: record[0],                    // remember that each record is a 2 element slice [...] that represents a `question, answer` pair
			answer:   strings.TrimSpace(record[1]), // strings.TrimSpace() helps to remove spaces around answers from the CSV file
		}
	}
	return returnedValue
}

//
// defines the quiz handler that:
//   * extracts the question and answer from the CSV
//   * parses it into a problem struct format
// 	 * asks the user by iterating through the parsed data
//
func questionHandler(file *os.File) {
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		errMsgHandler("Failed to parse the provided CSV file")
	}
	problems := parseRecords(records)

	// iterate through the questions with the user
	correctAnsCount := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
		var userAnswer string                        // define variable to store users answer to question
		fmt.Scanf("%s\n", &userAnswer)               // read in the user's answer to question, while removing all useless spaces around string
		if reflect.DeepEqual(userAnswer, p.answer) { // compare user's answer to the actual answer
			correctAnsCount++
		}
	}
	fmt.Printf("You got %d out %d questions correct.\n", correctAnsCount, len(problems))
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
	questionHandler(openedFile)
}
