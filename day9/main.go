package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rowNum := 0
	prev := []int{}
	for scanner.Scan() {
		row := scanner.Text()
		val, err := strconv.Atoi(row)
		if err != nil {
			fmt.Println("Failed converting number: ", err)
			continue
		}

		if len(prev) == 25 {
			if containsSum(prev, val) {
				fmt.Println("found sum", rowNum, val)
			} else {
				fmt.Println("didn't find sum", rowNum, val)
				break
			}
		}
		prev = append(prev, val)
		rowNum++
		if rowNum > 25 {
			prev[0] = 0
			prev = prev[1:]
		}
		fmt.Println(rowNum, val, len(prev))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func containsSum(list []int, val int) bool {
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list); j++ {
			if j == i {
				continue
			} else if list[i]+list[j] == val {
				fmt.Println(list[i], list[j], val)
				return true
			}
		}
	}
	return false
}
