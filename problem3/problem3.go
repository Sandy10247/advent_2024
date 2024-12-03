package problem3

import (
	"advent_2024/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Problem2() (totalMul int) {
	defer utils.TimeFn("Problem 3")()

	// read input from file
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

	return
}
