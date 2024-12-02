package main

import (
    "bufio"
    "fmt"
    "os"
	"strconv"
	"strings"
)

func main() {
    // Open the file
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

	// Arrays to store the left and right lists
	leftListArray := []int{}
	rightListArray := []int{}

    // Create a scanner to read the file line by line
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		// For each line, split it into two numbers, convert them to ints and store them in an array
		numbers := strings.Split(scanner.Text(), "   ")
		num1, err := strconv.Atoi(numbers[0])
		num2, err := strconv.Atoi(numbers[1])
		leftListArray = append(leftListArray, num1)
		rightListArray = append(rightListArray, num2)

		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return
		}
		
    }

	// Store the frequency of each number in the right list in a map
	frequencyMap := make(map[int]int)
	for _, num := range rightListArray {
		frequencyMap[num]++
	}

	// Loop through the left list, and for each number, multiply it by the frequency of the number in the right list
	total := 0
	for _, num := range leftListArray {
		if _, exists := frequencyMap[num]; exists {
			total += num * frequencyMap[num]
		}
	}

	fmt.Println(total)

    // Check for errors during scanning
    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    }
}