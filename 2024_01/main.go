package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var list1 []int
	var list2 []int
	var diffList []int
	var totalDistance int

	content, error := os.ReadFile("input.txt")

	if error != nil {

		log.Fatal(error)
	}

	input := string(content)

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		parts := strings.Fields(line)

		var1, err1 := strconv.Atoi(parts[0])
		var2, err2 := strconv.Atoi(parts[1])

		if err1 != nil || err2 != nil {
			fmt.Println("Error converting to int", err1, err2)
			continue
		}

		list1 = append(list1, var1)
		list2 = append(list2, var2)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	for i := range list1 {
		varDiff := diff(list1[i], list2[i])
		diffList = append(diffList, varDiff)
	}

	totalDistance = sumArray(diffList)

	fmt.Println("Total Distance:", totalDistance)

	fmt.Println("----------end of part 1--------")

	countList2 := countValueOccur(list2)

	var totalScore int
	var arrScore []int

	for _, v := range list1 {
		tmpScore := countList2[v] * v
		arrScore = append(arrScore, tmpScore)
	}

	totalScore = sumArray(arrScore)
	fmt.Println("total Score ", totalScore)
	fmt.Println("----------end of part 2--------")
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func sumArray(numbers []int) int {
	result := 0
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}
	return result
}

func countValueOccur(arr []int) map[int]int {
	dict := make(map[int]int)
	for _, num := range arr {
		dict[num] = dict[num] + 1
	}
	return dict
}
