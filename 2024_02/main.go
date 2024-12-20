package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// The levels are either all increasing or all decreasing.
// Any two adjacent levels differ by at least one and at most three.

// 6 Reports (line)
// 5 Levels (column)

// 7 6 4 2 1: Safe because the levels are all decreasing by 1 or 2.
// 1 2 7 8 9: Unsafe because 2 7 is an increase of 5.
// 9 7 6 2 1: Unsafe because 6 2 is a decrease of 4.
// 1 3 2 4 5: Unsafe because 1 3 is increasing but 3 2 is decreasing.
// 8 6 4 4 1: Unsafe because 4 4 is neither an increase or a decrease.
// 1 3 6 7 9: Safe because the levels are all increasing by 1, 2, or 3.

// --- Part Two ---
// The engineers are surprised by the low number of safe reports until they realize they forgot to tell you about the Problem Dampener.
// The Problem Dampener is a reactor-mounted module that lets the reactor safety systems tolerate a single bad level in what would otherwise be a safe report. It's like the bad level never happened!
// Now, the same rules apply as before, except if removing a single level from an unsafe report would make it safe, the report instead counts as safe.
// More of the above example's reports are now safe:
//     7 6 4 2 1: Safe without removing any level.
//     1 2 7 8 9: Unsafe regardless of which level is removed.
//     9 7 6 2 1: Unsafe regardless of which level is removed.
//     1 3 2 4 5: Safe by removing the second level, 3.
//     8 6 4 4 1: Safe by removing the third level, 4.
//     1 3 6 7 9: Safe without removing any level.
// Thanks to the Problem Dampener, 4 reports are actually safe!

var extraLife bool = true
var easyTry bool = false
var safeRows int = 0

func main() {

	var input [][]int = ReadFile("input.txt")
	safeRows = countSafe(input)
	fmt.Print("Number Of Safe Rows: ", safeRows)
	fmt.Println("-----------------------------")
}

func ReadFile(fileName string) [][]int {
	readFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileLines [][]int

	for fileScanner.Scan() {
		line := fileScanner.Text()
		numStrs := strings.Fields(line) // Split the line by spaces

		var nums []int

		for _, numStr := range numStrs {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Error converting to int", numStr, err)
				continue
			}
			nums = append(nums, num)
		}
		fileLines = append(fileLines, nums)
	}
	readFile.Close()
	return fileLines
}

func countSafe(inputList [][]int) int {

	for _, block := range inputList {
		fmt.Println("block values", block)

		extraLife = true
		easyTry = false
		isBlockSafe := checkNumsBlock(block)

		fmt.Println("Block is", isBlockSafe)
		if isBlockSafe {
			safeRows += 1
		}
	}
	return safeRows
}

func checkNumsBlock(numsBlock []int) bool {
	var rowCheck bool = true

	for i := 0; i < len(numsBlock)-1; i++ {
		checkNumsDiff(i, numsBlock, numsBlock[i], numsBlock[i+1], rowCheck, extraLife, easyTry)

	}
	return rowCheck
}

func checkNumsDiff(i int, numsBlock []int, currNum, nextNum int, rowCheck, extraLife, easyTry bool) (bool, bool, bool) {
	var rowDecrease bool = false
	var rowIncrease bool = false
	var numsDiff int

	numsDiff = nextNum - currNum
	fmt.Println("diff is", numsDiff)

	fmt.Println("debug extralife", extraLife)
	fmt.Println("debug easyTry", easyTry)

	if easyTry {
		if i+2 >= len(numsBlock) {
			rowCheck = false
			// return rowCheck, extraLife, easyTry
		}
		easyTry = false
	}

	if numsDiff > 3 || numsDiff < -3 {
		if !extraLife {
			rowCheck = false
			// return rowCheck, extraLife, easyTry
		} else {
			extraLife = false
			easyTry = true
			fmt.Println("extraLife lost delta trop gros", extraLife)
			// reloop here
			fmt.Println("debug exta loop 1")
			checkNumsDiff(i, numsBlock, numsBlock[i], numsBlock[i+2], rowCheck, extraLife, easyTry)
			// return rowCheck, extraLife, easyTry
		}
	}

	if numsDiff > 0 && numsDiff <= 3 {
		if rowDecrease {
			if !extraLife {
				rowCheck = false
				// return rowCheck, extraLife, easyTry
			} else {
				extraLife = false
				easyTry = true
				fmt.Println("extraLife lost row decrease puis increase", extraLife)
				// reloop here
				fmt.Println("debug exta loop 2")
				checkNumsDiff(i, numsBlock, numsBlock[i], numsBlock[i+2], rowCheck, extraLife, easyTry)
				// return rowCheck, extraLife, easyTry
			}
		}
		rowIncrease = true
	}

	if numsDiff < 0 && numsDiff >= -3 {
		if rowIncrease {
			if !extraLife {
				rowCheck = false
				// return rowCheck, extraLife, easyTry
			} else {
				extraLife = false
				easyTry = true
				fmt.Println("extraLife lost row increase puis decrease", extraLife)
				// reloop here
				fmt.Println("debug exta loop 3")
				checkNumsDiff(i, numsBlock, numsBlock[i], numsBlock[i+2], rowCheck, extraLife, easyTry)
				// return rowCheck, extraLife, easyTry
			}
		}
		rowDecrease = true
	}

	if numsDiff == 0 {
		if !extraLife {
			rowCheck = false
			return rowCheck, extraLife, easyTry
		} else {
			extraLife = false
			easyTry = true
			fmt.Println("extraLife lost, diff is 0", extraLife)
			// reloop here
			fmt.Println("debug exta loop 4")
			checkNumsDiff(i, numsBlock, numsBlock[i], numsBlock[i+2], rowCheck, extraLife, easyTry)
			// return rowCheck, extraLife, easyTry
		}
	}
	return rowCheck, extraLife, easyTry
}
