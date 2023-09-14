// Example of a DFSM

package main

import (
	"fmt"
	"slices"
)

func dfsm(inputString string) bool {
	inputSymbolSet := []rune{'a', 'b', 'c'}
	transitionTable := [][]int{
		{0, 0, 0},
		{2, 1, 3}, // 1
		{4, 2, 1}, // 2
		{3, 4, 2}, // 3
		{1, 3, 2}, // 4
	}
	acceptingStates := []int{3}
	currentState := 1

	for _, symbol := range inputString {
		currentState = transitionTable[currentState][slices.Index(inputSymbolSet, symbol)]
	}

	return isAcceptingState()
}

func main() {
	fmt.Println("Hello Wcforld")
	inputString := "abcc"
	if dfsm(inputString) {
		fmt.Println("\"%s\" is ACCEPTED!", inputString)
	} else {
		fmt.Println("\"%s\" is REJECTED.", inputString)
	}
}
