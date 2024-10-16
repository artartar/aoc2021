package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// flag to allow any input file to be utilized
	dailyInput := flag.String("input", "day1_input.txt", "Puzzle input provided for the daily puzzle")
	userFile := *dailyInput

	file, err := os.Open(userFile)

	if err != nil {
		fmt.Printf("Error opening file %s\n", userFile)
	}

	// After all flag declaration is complete, parsing required
	flag.Parse()
	defer file.Close()

	puzzleData := bufio.NewScanner(file)
	var compareMe []int
	// var count int

	for puzzleData.Scan() {
		
		line := puzzleData.Text()
		num, lineErr := strconv.Atoi(line)
		
		// Error handling for conversion of line to int
		if lineErr != nil {
			fmt.Printf("Could not convert %s to an int", line)
		}

		compareMe = append(compareMe, num)
		
	}

	// Check for errors during scanning
	if scanErr := puzzleData.Err(); scanErr != nil {
		fmt.Println("Error scanning file:", err)	
	}

	fmt.Printf("%d", len(compareMe))
	
}
