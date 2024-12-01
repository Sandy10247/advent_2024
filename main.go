package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func problem1() {
	// read input from file
	data, err := os.ReadFile("problem1_input.txt")
	if err != nil {
		panic("File Opening issue idiot 🤪")
	}
	// seperate the data into arr1, arr2
	dataStr := string(data)
	arr1, arr2 := make([]string, 0), make([]string, 0)
	for _, item := range strings.Split(dataStr, "\n") {
		splittedData := strings.Split(item, "   ")
		arr1 = append(arr1, splittedData[0])
		arr2 = append(arr2, splittedData[1])
	}

	fmt.Printf("array :- %v\n", arr2)
	slices.Sort(arr1)
	slices.Sort(arr2)
	finalDist := 0
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
			finalDist += x - y
		} else {
			finalDist += y - x
		}
	}

	fmt.Println("Final Dist :- ", finalDist)

}

func main() {
	problem1()
}
