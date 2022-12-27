package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const rock = 1
const paper = 2
const scissors = 3
const lose = 0
const draw = 3
const win = 6

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	acc := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strat := scanner.Text()
		acc += score(strat)
	}
	fmt.Printf("Your final score is %d", acc)
}

func score(strat string) int {
	switch strat {
	case "A X":
		return rock + draw
	case "B X":
		return rock + lose
	case "C X":
		return rock + win

	case "A Y":
		return paper + win
	case "B Y":
		return paper + draw
	case "C Y":
		return paper + lose

	case "A Z":
		return scissors + lose
	case "B Z":
		return scissors + win
	case "C Z":
		return scissors + draw
	default:
		panic("unexpected input")
	}
}
