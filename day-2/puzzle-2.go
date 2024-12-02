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

		// Check if the original list is safe
		if calculateSafety(numbersInt) {
			safeReports++
			continue
		}

		// Check if removing one number makes it safe
		for i := 0; i < len(numbersInt); i++ {
			// Create a new slice by removing the number at index i
			newNumbersInt := []int{}
			newNumbersInt = append(newNumbersInt, numbersInt[:i]...)
			newNumbersInt = append(newNumbersInt, numbersInt[i+1:]...)

			// Check if the new combination is safe
			if calculateSafety(newNumbersInt) {
				safeReports++
				break // Exit the loop if a safe combination is found
			}
		}
	}

	fmt.Println(safeReports)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
