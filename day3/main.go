package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rowNum := 0
	treesHit1 := 0
	treesHit3 := 0
	treesHit5 := 0
	treesHit7 := 0
	treesHitVert := 0
	for scanner.Scan() {
		row := scanner.Text()
		hit1 := checkSlopeCoordinate(row, rowNum, 1)
		if hit1 {
			treesHit1++
		}
		hitVert := checkSlopeCoordinate(row, rowNum/2, 1)
		if hitVert && rowNum%2 == 0 {
			treesHitVert++
		}
		hit3 := checkSlopeCoordinate(row, rowNum, 3)
		if hit3 {
			treesHit3++
		}
		hit5 := checkSlopeCoordinate(row, rowNum, 5)
		if hit5 {
			treesHit5++
		}
		hit7 := checkSlopeCoordinate(row, rowNum, 7)
		if hit7 {
			treesHit7++
		}

		rowNum++
	}
	fmt.Printf("Right 1 Down 1: %d\n", treesHit1)
	fmt.Printf("Right 3 Down 1: %d\n", treesHit3)
	fmt.Printf("Right 5 Down 1: %d\n", treesHit5)
	fmt.Printf("Right 7 Down 1: %d\n", treesHit7)
	fmt.Printf("Right 1 Down 2: %d\n", treesHitVert)
	fmt.Println(treesHit1 * treesHit3 * treesHit5 * treesHit7 * treesHitVert)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func checkSlopeCoordinate(row string, rowNum, right int) bool {
	coordinate := (rowNum * right) % len(row)
	return row[coordinate] == '#'
}

/*
*
****
*******
**********
*************
****************
*******************
**********************
*************************
****************************
*******************************


 */
