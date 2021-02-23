package main

import "fmt"

func indexToTurn(i int) int {
	return i + 1
}

func calculateSpokenNumber(previousTurnSpoken int, currentTurnSpoken int) int {
	if previousTurnSpoken != 0 {
		return currentTurnSpoken - previousTurnSpoken
	} else {
		return 0
	}
}

func calculateLastSpokenNumberAtTurnWithStartingNumbers(turnToSpeak int, startingNumbers []int) (spokenNumber int) {
	memory := make(map[int]int)
	previousTurnSpoken := 0
	for i:=0; i<turnToSpeak; i++ {
		if i>=len(startingNumbers) { // main logic
			spokenNumber = calculateSpokenNumber(previousTurnSpoken, i)
		} else { // loading logic
			spokenNumber = startingNumbers[i]
		}
		previousTurnSpoken = memory[spokenNumber]
		memory[spokenNumber] = indexToTurn(i)
	}
	return
}

func main() {
	number:= calculateLastSpokenNumberAtTurnWithStartingNumbers(30000000,[]int{9,12,1,4,17,0,18})
	fmt.Printf("Number: %d\n", number)
}

