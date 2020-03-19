### goQuiz

A simple Golang implementation of a Quiz that:

* asks a user a simple questions
* validates their answer
* computes their % correct answer
* times the user in between questions
* sources both the question and answers from a CSV file

***

The code leverages channels/go-routines alongside the following:

* `flags` package
* `csv` package
* `os` package
* `time` package
* `reflect` package
* `strings` package

***

### To build and run

Make sure you are within your local Golang work environment, then run the following:

```bash
    $ git clone git@github.com:damilarelana/goQuiz.git
    $ cd goQuiz
    $ go build . && ./goQuiz
```

### To run with different Quiz data

The original quiz data is stored as `quizData.csv`. Its data content/filename can be changed. If the quiz data content/filename is changed, then proceed to use the `csv` flag to re-run using the following:

```bash
    $ ./goQuiz -csv="<newFileName.csv>"
```
where `<newFileName.csv>` refers to the full-path for the new quiz data content/filename.

### To run the quiz with different time duration per question

The `limit` flag can be use to adjust the countdown timer (in `seconds`) for each question i.e. the quiz would end if the user does not supply an answer within a time limit. Currently, the default value is `30` i.e. `30 seconds`. To chnage the limit use:

```bash
    $ ./goQuiz -limit="45"
```
where `45` means the user now `45 seconds` to answer each question.