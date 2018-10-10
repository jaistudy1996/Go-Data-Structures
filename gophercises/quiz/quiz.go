package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	quiz := readCSV("./problems.csv")
	correctAnswers := 0
	for _, ques := range quiz {
		consoleReader := bufio.NewReader(os.Stdin)
		fmt.Print(ques[0], " ")
		answer, _ := consoleReader.ReadString('\n')
		if strings.TrimRight(answer, "\n") == ques[1] {
			correctAnswers++
		}
	}
	fmt.Printf("Score: %d/%d", correctAnswers, len(quiz))
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
