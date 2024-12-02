package problem2

import (
	"advent_2024/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isIncreasing(arr []int) bool {
	curr := arr[0]
	for _, v := range arr[1:] {
		if (v > curr) && utils.Diff(v, curr) < 4 && utils.Diff(v, curr) > 0 {
			curr = v
		} else {
			return false
		}
	}

	return true
}

func isDecreasing(arr []int) bool {
	curr := arr[0]
	for _, v := range arr[1:] {
		if v < curr && utils.Diff(v, curr) < 4 && utils.Diff(v, curr) > 0 {
			curr = v
		} else {
			return false
		}
	}

	return true
}

func Problem2() (pureSafe int, dampedSafe int) {
	defer utils.TimeFn("Problem 2")()

	// read input from file
	data, err := os.ReadFile(utils.PWD() + "/problem2/problem2_input.txt")
	if err != nil {
		panic(fmt.Errorf("File Opening issue idiot 🤪 %v", err))
	}
	// seperate the data into arr1, arr2
	dataStr := string(data)
	for _, item := range strings.Split(dataStr, "\n") {
		splittedData := strings.Split(item, " ")
		splittedDataInts := []int{}
		for _, itemStr := range splittedData {
			intVal, err := strconv.Atoi(itemStr)
			if err != nil {
				panic(fmt.Errorf("String to Int issue idiot 🤪 %v", err))
			}
			splittedDataInts = append(splittedDataInts, intVal)
		}

		if isIncreasing(splittedDataInts) {
			pureSafe += 1
		} else if isDecreasing(splittedDataInts) {
			pureSafe += 1
		} else {
			fmt.Printf("false :- %v\n", splittedDataInts)
		}

		// fmt.Printf("Before curr :- %v\n", splittedDataInts)
		// check damped safe record
		// if isIncreasing(splittedDataInts, true) {
		// 	dampedSafe += 1
		// } else if isDecreasing(splittedDataInts, true) {
		// 	dampedSafe += 1
		// }

	}
	return
}
