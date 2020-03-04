package main

import "flag"

func main() {
	csvFilename := flag.String("csv", "quizData.csv", "a csv file containing the quiz's question and answer data, in the format 'question, answer'")
	flag.Parse() // require to initialize the specified flags with the Operating system
}
