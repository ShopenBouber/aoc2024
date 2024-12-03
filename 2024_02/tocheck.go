package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := ReadFile("input.txt")
	safeRows := countSafe(input)
	fmt.Println("Number Of Safe Rows:", safeRows)
}

// ReadFile reads the input from a file and parses it into a 2D slice of integers.
func ReadFile(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data [][]int

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		var row []int
		for _, numStr := range line {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Error converting to int:", numStr, err)
				continue
			}
			row = append(row, num)
		}
		data = append(data, row)
	}

	return data
}

// countSafe counts the number of safe rows, considering single-level removals.
func countSafe(input [][]int) int {
	safeCount := 0
	for _, block := range input {
		if isSafe(block) || isSafeWithOneRemoval(block) {
			safeCount++
		}
	}
	return safeCount
}

// isSafe checks if a block is inherently safe.
func isSafe(block []int) bool {
	if len(block) < 2 {
		return true // Trivial case
	}

	isIncreasing := block[1] > block[0]
	for i := 0; i < len(block)-1; i++ {
		diff := block[i+1] - block[i]
		if diff < -3 || diff > 3 || (diff > 0 && !isIncreasing) || (diff < 0 && isIncreasing) {
			return false
		}
	}
	return true
}

// isSafeWithOneRemoval checks if a block can be made safe by removing one level.
func isSafeWithOneRemoval(block []int) bool {
	for i := 0; i < len(block); i++ {
		// Create a copy of the block with one level removed.
		modified := append(block[:i], block[i+1:]...)
		if isSafe(modified) {
			return true
		}
	}
	return false
}
