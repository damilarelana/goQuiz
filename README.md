### goQuiz

A simple Golang implementation of a Quiz that:

* asks a user a simple questions
* validates their answer
* computes their % correct answer
* times the user in between questions
* sources both the question and answers from a CSV file

***

The code leverages channels/go-routines alongside the following:

* flags package
* csv package
* os package
* time package