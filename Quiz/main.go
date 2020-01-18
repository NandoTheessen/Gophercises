package main


import (
	"os"
	"time"
	"flag"
	"strings"
	"log"
	"bufio"
	"fmt"
	"encoding/csv"
	"io"
)

type Task struct {
	term string
	result string
}

var filename string
var timer int


func init() {
	flag.StringVar(&filename, "file", "problems.csv", "csv file in the format of 'question,answer', defaults to problems.csv")
	flag.IntVar(&timer, "time", 30, "the time limit for the quiz")
	flag.Parse()
}


// ParseCsv reads the content of a csv and parses it's contents into a predefined struct "Task"
// as a paremeter it takes the filename string, pointing to a file on the filesystem.
func ParseCsv(filename string) (tasks []Task, err error){
	var f *os.File
	if filename != "" {
		f, err = os.Open(filename)
		if err != nil {
			return nil, fmt.Errorf("File %s does not exists\n", filename)
		}
	} else {
		f, err = os.Open("./problems.csv")
	}

	defer f.Close()

	scanner := csv.NewReader(bufio.NewReader(f))

	for {
		val, err := scanner.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
		}
		tasks = append(tasks, Task{term: val[0], result: val[1]})
	}

	return tasks, nil
}

// RunQuiz iterates over the parsed Tasks, prompts the user with a term and then compares the user input
// with the desired solution
func RunQuiz(tasks []Task) string{
	stdinReader := bufio.NewReader(os.Stdin)
	var result int
	maxPoints := len(tasks)

	for _, task := range tasks {
		fmt.Printf("%s = ", task.term)
		text, _ := stdinReader.ReadString('\n')
		text = strings.TrimRight(text, "\n")
		if task.result == text {
			result++
		}
	}
	return fmt.Sprintf("You've scored %d out of %d Points!\n", result, maxPoints)
}


func main() {

	// Parse the CSV file & combine the output into an array of Tasks
	tasks, err := ParseCsv(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Create two channels, one for the result of the quizz, one to send interrupt signal
	// once the time is up.
	done := make(chan bool)
	res := make(chan string)

	// Timeout goroutine sending interrupt signal after <timeout>
	go func() {
		duration := time.Duration(timer)
		time.Sleep(duration * time.Second)
		done <- true
	}()
	// Quiz goroutine running the quiz in the meantime
	go func() {
		res <- RunQuiz(tasks)
	}()

	select {
	case <- done:
		fmt.Print("\nYou run out of time!\n")
		return
	case r := <- res:
			fmt.Println(r)
	}
}
	