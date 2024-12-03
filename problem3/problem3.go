package problem3

import (
	"advent_2024/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Problem2() (totalMul int, totalWithCorruped int) {
	defer utils.TimeFn("Problem 3")()

	// read input from file
	// data, err := os.ReadFile(utils.PWD() + "/problem3/problem3_input_2_test.txt")
	data, err := os.ReadFile(utils.PWD() + "/problem3/problem3_input.txt")

	if err != nil {
		panic(fmt.Errorf("File Opening issue idiot 🤪 %v", err))
	}
	// seperate the data into arr1, arr2
	dataStr := string(data)
	pattern, compileErr := regexp.Compile(`([m+][u+][l+])\(([0-9]+,[0-9]+)\)`)
	if compileErr != nil {
		panic(fmt.Errorf("idiot Failed to create Compile 🤪 %v", err))

	}

	for _, item := range strings.Split(dataStr, "\n") {
		// fmt.Printf("item :- %v\n", item)
		matchedStrings := pattern.FindAllString(item, -1)
		for _, mulItem := range matchedStrings {
			cleanedString := strings.Split(strings.ReplaceAll(strings.ReplaceAll(mulItem, "mul(", ""), ")", ""), ",")
			x, err := strconv.Atoi(cleanedString[0])
			if err != nil {
				panic(fmt.Errorf("idiot Failed to convert to int 🤪 %v", err))
			}
			y, err := strconv.Atoi(cleanedString[1])
			if err != nil {
				panic(fmt.Errorf("idiot Failed to convert to int 🤪 %v", err))
			}
			totalMul += x * y

		}
	}

	// ref :-https://www.reddit.com/r/adventofcode/comments/1h5obsr/2024_day_3_regular_expressions_go_brrr/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button
	redditRegex, err := regexp.Compile("mul\\(\\d+,\\d+\\)|do\\(\\)|don't\\(\\)")
	if err != nil {
		fmt.Printf("reddit regex failed")
	}

	firstAns := 0
	for _, item := range strings.Split(dataStr, "\n") {
		// fmt.Printf("item :- %v\n", item)
		flag := true
		res := redditRegex.FindAllString(item, -1)
		for _, itemEach := range res {
			// fmt.Printf("itemEach :- %v\n", itemEach)
			if itemEach == "do()" {
				flag = true
			} else if itemEach == "don't()" {
				flag = false
			} else {
				splitted := strings.Split(itemEach[4:len(itemEach)-1], ",")
				x, err := strconv.Atoi(splitted[0])
				if err != nil {
					panic("x convertion failed")
				}
				y, err := strconv.Atoi(splitted[1])
				if err != nil {
					panic(fmt.Errorf("y conversion failed %v", err))
				}
				if flag {
					totalWithCorruped += x * y
				} else {
					firstAns += x * y
				}
			}

		}

	}

	return
}
