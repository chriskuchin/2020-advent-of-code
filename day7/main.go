package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

var (
	ruleParser = regexp.MustCompile(`^([\w ]+) bags contain ((\d|no?) ([\w ]+),? )+bags?.`)
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rowNum := 0
	for scanner.Scan() {
		row := scanner.Text()
		parseRule(row)
		fmt.Println(row)
		rowNum++
		if rowNum == 3 {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func parseRule(ruleString string) {
	for _, val := range ruleParser.FindAllStringSubmatch(ruleString, -1) {
		for _, piece := range val {
			fmt.Println(piece)
		}
	}
}
