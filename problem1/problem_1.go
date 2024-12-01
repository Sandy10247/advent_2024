package problem1

import (
	"advent_2024/utils"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Problem1() (distance int, occurencesTotal int) {
	defer utils.TimeFn("Problem1")()
	// read input from file
	data, err := os.ReadFile(utils.PWD() + "/problem1/problem1_input.txt")
	if err != nil {
		panic(fmt.Errorf("File Opening issue idiot 🤪 %v", err))
	}
	// seperate the data into arr1, arr2
	dataStr := string(data)
	arr1, arr2 := make([]string, 0), make([]string, 0)
	for _, item := range strings.Split(dataStr, "\n") {
		splittedData := strings.Split(item, "   ")
		arr1 = append(arr1, splittedData[0])
		arr2 = append(arr2, splittedData[1])
	}

	slices.Sort(arr1)
	slices.Sort(arr2)
	for i := 0; i < len(arr1); i += 1 {

		x, err := strconv.Atoi(arr1[i])
		if err != nil {
			panic("String to Int Op issue idiot 🤪")

		}
		y, err := strconv.Atoi(arr2[i])
		if err != nil {
			panic("String to Int Op issue idiot 🤪")

		}
		if x > y {
			distance += x - y
		} else {
			distance += y - x
		}
	}

	occrencesMap := map[string]int{}
	// prepare occurances map for arr2
	for _, item := range arr2 {
		_, doesExist := occrencesMap[item]
		if doesExist {
			occrencesMap[item] += 1
		} else {
			occrencesMap[item] = 1
		}
	}

	// calculate occurance total count
	for _, item := range arr1 {
		val, _ := occrencesMap[item]
		itemInt, err := strconv.Atoi(item)
		if err != nil {
			panic(fmt.Errorf("Idiot couldnt convert string to int "))
		}
		occurencesTotal += (itemInt * val)
	}

	return
}
