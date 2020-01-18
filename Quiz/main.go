package main


import (
	"os"
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


func init() {
	flag.StringVar(&filename, "file", "problems.csv", "Supply an optional filename, defaults to problems.csv")
	flag.Parse()
}

var filename string

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

func RunQuiz(tasks []Task){
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
	fmt.Printf("You've scored %d out of %d Points!\n", result, maxPoints)
}


func main() {

	tasks, err := ParseCsv(filename)
	if err != nil {
		log.Fatal(err)
	}

	RunQuiz(tasks)

}
	