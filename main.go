package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Read the data.dat file downloaded from the link
	file, err := os.Open("data.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create scanner to read from file
	scanner := bufio.NewScanner(file)

	// Variable to hold how many of 0s or 1s in each line
	var score = 0
	// Loop through each line
	for scanner.Scan() {
		// Get string in the line
		var line = scanner.Text()
		// Count 0 and 1 in this line
		var count_zero = strings.Count(line, "0")
		var count_one = strings.Count(line, "1")

		// check if 0's count is multiple of 3, then increment score
		if count_zero%3 == 0 {
			score++
			continue
		}
		// check if 1's count is multiple of 2, then increment score
		if count_one%2 == 0 {
			score++
			continue
		}
		// use `continue` to go to next loop iteration
	}

	fmt.Println("total score: ", score)

	// catch any error
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
