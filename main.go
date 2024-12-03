package main

import (
	"advent_2024/problem1"
	"advent_2024/problem2"
	"advent_2024/problem3"

	"fmt"
)

func main() {
	distance, occurencesTotal := problem1.Problem1()
	fmt.Printf("problem 1\nDistance :- %v, Occurence Count :- %v\n", distance, occurencesTotal)

	pureSafe, dampedSafe := problem2.Problem2()
	fmt.Printf("total safe records :- \npure %v, damped :- %v\n", pureSafe, dampedSafe)

	totalMul := problem3.Problem2()
	fmt.Printf("Problem 3 :- %v\n", totalMul)
}
