package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"
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

// create the user answerchannel [required to make the user response and timer to be unblocking with each other]
//   * required by getAnswer()
//   * required by questionHandler()
var answerChannel = make(chan string)

// define flags
var csvFilename *string = flag.String("csv", "quizData.csv", "a csv file containing question/answer data, in a 'question, answer' format per record line")
var maxTimeLimit *int = flag.Int("limit", 30, "the maximum allowed duration of time to answer each quiz question in seconds")

// define parser to
//   * read in the multi-dimensional slice of `question, answer` i.e. CSV file data
//   * and then return a 1 dimensional slice of `question, answer` i.e. a value of type `problem`
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

// defines the quiz handler that:
//   * extracts the question and answer from the CSV
//   * parses it into a problem struct format
// 	 * asks the user by iterating through the parsed data
func questionHandler(file *os.File) {
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		errMsgHandler("Failed to parse the provided CSV file")
	}
	problems := parseRecords(records)

	correctAnsCount := 0         // initialize counter for number of questions answer correctly
	for i, p := range problems { // iterate through the questions with the user
		timer := time.NewTimer(time.Duration(*maxTimeLimit) * time.Second) // initialize (and re-initialize) a timer for each question
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
		go getAnswer() // get userAnswer

		select { // when user is within time limit, then there is no message tine the channel timer.C
		case <-timer.C: //checks if timer expired before any question was answered
			fmt.Println(questionCompletionMsg(problems, correctAnsCount))
			return
		case userAnswer := <-answerChannel: // when there is a user answer in the answerChannel
			if reflect.DeepEqual(userAnswer, p.answer) { // compare user's answer to the actual answer
				correctAnsCount++
			}
		}
	}
	fmt.Println(questionCompletionMsg(problems, correctAnsCount))
}

// defines the function to be made a goroutine to ensure that timer and user answer are not blocking
func getAnswer() {
	var userAnswer string          // define variable to store users answer to question
	fmt.Scanf("%s\n", &userAnswer) // read in the user's answer to question, while removing all useless spaces around string
	answerChannel <- userAnswer
}

// defines the function to help print quiz completion messaging
func questionCompletionMsg(problems []problem, correctAnsCount int) string {
	// calculate quiz completion accuracy
	numQuestions := len(problems) // calculate number of questions in the quiz
	correctAnsPercentage := 100.0 * float64(correctAnsCount) / float64(numQuestions)

	// return the quiz completion result string
	return fmt.Sprintf("\n You got %d out %d questions correct i.e. %.4g%% accuracy.\n", correctAnsCount, numQuestions, correctAnsPercentage)
}

func main() {
	// parse flags
	flag.Parse() // required to initialize the specified flags with the Operating system

	// open CSV file and read the records (i.e. all the lines) from it
	openedFile, err := os.Open(*csvFilename)
	if err != nil {
		errMsgHandler(fmt.Sprintf("Failed to open CSV file: %s\n", *csvFilename))
	}
	questionHandler(openedFile)
}
