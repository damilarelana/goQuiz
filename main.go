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
	csvFilename := flag.String("csv", "quizData.csv", "a csv file containing quiz's question and answer data, in the format 'question, answer'")
	flag.Parse() // require to initialize the specified flags with the Operating system

	openedFile, err := os.Open(*csvFilename)
	if err != nil {
		errMsgHandler(fmt.Sprintf("Failed to open CSV file: %s\n", *csvFilename))
	}
	r := csv.NewReader(openedFile)

}
