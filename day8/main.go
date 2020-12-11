package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type (
	command struct {
		cmd     string
		value   int
		visited bool
	}
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rowNum := 0
	cmds := []*command{}
	for scanner.Scan() {
		row := scanner.Text()
		cmd := strings.Split(row, " ")
		val, err := strconv.Atoi(cmd[1])
		if err != nil {
			fmt.Println("Failed parsing number", err)
		}
		cmds = append(cmds, &command{
			cmd:     cmd[0],
			value:   val,
			visited: false,
		})
		rowNum++
	}

	currentValue := 0
	currentLocation := 0
	var currentCmd *command
	for {
		if cmds[currentLocation].visited {
			fmt.Println("Found repeat", currentCmd, currentValue)
			break
		}
		currentCmd = cmds[currentLocation]
		currentCmd.visited = true

		switch currentCmd.cmd {
		case "jmp":
			fmt.Println("jmp ", currentCmd)
			currentLocation += currentCmd.value
		case "acc":
			fmt.Println("acc ", currentCmd)
			currentValue += currentCmd.value
			currentLocation++
		case "nop":
			fmt.Println("nop ", currentCmd)
			currentLocation++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
