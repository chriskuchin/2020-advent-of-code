package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rowNum := 0
	maxID := 0
	seats := []int{}
	for scanner.Scan() {
		row := scanner.Text()

		maxRow := 127
		minRow := 0
		for i := 0; i < 7; i++ {
			midPoint := ((maxRow - minRow) / 2) + minRow
			upperMidpoint := midPoint + 1

			if row[i] == 'F' {
				maxRow = midPoint
			} else if row[i] == 'B' {
				minRow = upperMidpoint
			}
		}

		maxCol := 7
		minCol := 0
		for i := 7; i < 10; i++ {
			midPoint := ((maxCol - minCol) / 2) + minCol
			upperMidpoint := midPoint + 1
			if row[i] == 'L' {
				maxCol = midPoint
			} else if row[i] == 'R' {
				minCol = upperMidpoint
			}
		}

		boardingID := maxRow*8 + maxCol
		seats = append(seats, boardingID)

		if boardingID > maxID {
			maxID = boardingID
		}
		rowNum++
	}

	fmt.Println(maxID)
	sort.Sort(sort.IntSlice(seats))

	expectedSeat := 0
	for _, assignedSeat := range seats {
		fmt.Println(expectedSeat, assignedSeat)
		if expectedSeat == 0 {
			expectedSeat = assignedSeat + 1
		} else if assignedSeat != expectedSeat {
			fmt.Println(assignedSeat, expectedSeat)
			break
		} else {
			expectedSeat++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
