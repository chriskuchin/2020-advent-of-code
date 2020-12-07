package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rowNum := 0
	groupSummary := map[string]bool{}
	sum := 0
	groupMember := 0
	for scanner.Scan() {
		row := scanner.Text()
		fmt.Println(groupMember, row)
		if row != "" {
			if groupMember == 0 {
				for _, q := range row {
					question := string(q)
					groupSummary[question] = true
				}
			} else {
				fmt.Println(groupSummary)
				for question := range groupSummary {
					fmt.Println(question, row, strings.Contains(row, question))
					if !strings.Contains(row, question) {
						delete(groupSummary, question)
					}
				}

			}
			groupMember++
		} else {
			fmt.Println(groupSummary, len(groupSummary), sum)
			groupMember = 0
			sum = sum + len(groupSummary)
			groupSummary = map[string]bool{}
		}
		rowNum++
	}
	sum = sum + len(groupSummary)

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
