package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// A function to check if a list of numbers is safe
func calculateSafety(numbers []int) bool {
	isIncreasing := true
	isGraduallyIncreasing := true
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] >= numbers[i+1] {
			isIncreasing = false
			break
		}
		if int(math.Abs(float64(numbers[i]-numbers[i+1]))) > 3 || numbers[i] == numbers[i+1] {
			isGraduallyIncreasing = false
			break
		}
	}

	isDecreasing := true
	isGraduallyDecreasing := true
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] <= numbers[i+1] {
			isDecreasing = false
			break
		}
		if int(math.Abs(float64(numbers[i]-numbers[i+1]))) > 3 || numbers[i] == numbers[i+1] {
			isGraduallyDecreasing = false
			break
		}
	}

	return (isIncreasing && isGraduallyIncreasing) || (isDecreasing && isGraduallyDecreasing)
}

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Keep track of the number of safe reports
	safeReports := 0

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// For each line, split it into a list of numbers as integers
		numbers := strings.Split(scanner.Text(), " ")
		numbersInt := []int{}
		for _, number := range numbers {
			num, err := strconv.Atoi(number)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return
			}
			numbersInt = append(numbersInt, num)
		}

		// Check if the list of numbers is safe
		isSafe := calculateSafety(numbersInt)
		if isSafe {
			safeReports++
		}

	}

	fmt.Println(safeReports)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
