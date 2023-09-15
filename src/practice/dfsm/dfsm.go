// Example of a DFSM

package main

import (
	"fmt"
	"slices"
)

func isAcceptingState(currentState int, acceptingStates []int) bool {
	return slices.Contains(acceptingStates, currentState)
}

func dfsm(inputString string) bool {
	inputSymbolSet := []rune{'a', 'b', 'c'}
	transitionTable := [][]int{
		//a  b  c
		{0, 0, 0},
		{2, 1, 3}, // 1
		{4, 2, 1}, // 2
		{3, 4, 2}, // 3
		{1, 3, 2}, // 4
	}
	acceptingStates := []int{3}
	currentState := 1

	for _, symbol := range inputString {
		columnIndex := slices.Index(inputSymbolSet, symbol)
		if columnIndex == -1 {
			fmt.Printf("Invalid symbol: %c\n", symbol)
			return false
		}
		fmt.Printf("Current state: %d, Symbol: %c\n", currentState, symbol)
		currentState = transitionTable[currentState][columnIndex]
		fmt.Printf("New state: %d\n\n", currentState)
	}
	fmt.Printf("Final state: %d\n", currentState)
	return isAcceptingState(currentState, acceptingStates)
}

func main() {
	fmt.Println("Hello Wcforld")
	inputString := "abcc"
	if dfsm(inputString) {
		fmt.Printf("\"%s\" is ACCEPTED!\n", inputString)
	} else {
		fmt.Printf("\"%s\" is REJECTED.\n", inputString)
	}
}
