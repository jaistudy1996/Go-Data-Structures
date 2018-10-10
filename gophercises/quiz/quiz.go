package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	quiz := readCSV("./problems.csv")
	correctAnswers := 0
	answeredQues := 0
	timer := time.AfterFunc(time.Second*30, func() {
		fmt.Printf("\nScore: %d/%d. Total questions answered: %d",
			correctAnswers, len(quiz), answeredQues)
		os.Exit(0) // Exit when we are out of time.
	})
	defer timer.Stop()
	for _, ques := range quiz {
		consoleReader := bufio.NewReader(os.Stdin)
		fmt.Print(ques[0], " ")
		answer, _ := consoleReader.ReadString('\n')
		answeredQues++
		if strings.TrimRight(answer, "\n") == ques[1] {
			correctAnswers++
		}
	}
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func readCSV(fileName string) [][]string {
	csvFile, err := os.Open(fileName)
	checkError(err)
	csvReader := csv.NewReader(csvFile)
	var records [][]string

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		records = append(records, record)
	}
	return records
}
