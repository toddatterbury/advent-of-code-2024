package main

import (
    "bufio"
    "fmt"
    "os"
	"strconv"
	"strings"
	"sort"
	"math"
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

	// Sort the left and right lists
	sort.Ints(leftListArray)
	sort.Ints(rightListArray)

	// Calculate the absolute differences and keep a total sum 
	sum := 0
	for i, leftValue := range leftListArray {
		diff := int(math.Abs(float64(leftValue - rightListArray[i])))
		sum += diff
	}

	fmt.Println(sum)

    // Check for errors during scanning
    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    }
}