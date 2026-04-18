package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	question string
	answer string
}

func main() {
	filename := flag.String("csv", "problems.csv", "CSV file containing quiz")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV, ", err)
		return
	}
	
	problems := parseLines(records)
	score := 0

	for i, p := range problems {
		fmt.Printf("problem #%d: %s = \n", i + 1, p.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.answer {
			score++
		}
	}

	fmt.Println("your total score is: ", score, " and the no. of incorrect answers were: ", len(records) - score)
	fmt.Printf("%s", records)
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	fmt.Printf("the type of ret is: %T\n", ret)

	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer: line[1],
		}
	}
	return ret
}