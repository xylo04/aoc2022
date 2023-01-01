package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var moveRegex, _ = regexp.Compile("move (\\d+) from (\\d+) to (\\d+)")

type ContainerState struct {
	stacks []Stack[rune]
}

func (c ContainerState) String() string {
	tops := ""
	for _, s := range c.stacks {
		tops += string(s.Peek())
	}
	return fmt.Sprintf("ContainerState with %d stacks, tops %s", len(c.stacks), tops)
}

type Move struct {
	num  uint
	from uint
	to   uint
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	state := readContainerState(scanner)
	moves := readMoves(scanner)

	runAllMoves(state, moves)
	fmt.Println(state)
}

func runAllMoves(state ContainerState, moves []Move) {
	for _, m := range moves {
		runOneMove(state, m)
	}
}

func runOneMove(state ContainerState, m Move) {
	for i := 0; uint(i) < m.num; i++ {
		container, _ := state.stacks[m.from].Pop()
		state.stacks[m.to].Push(container)
	}
}

func readMoves(scanner *bufio.Scanner) []Move {
	var moves []Move
	for scanner.Scan() {
		line := scanner.Text()
		matches := moveRegex.FindStringSubmatch(line)
		m := Move{parseInt(matches[1]), parseInt(matches[2]) - 1, parseInt(matches[3]) - 1}
		moves = append(moves, m)
	}
	return moves
}

// readContainerState reads containers from the scanner
func readContainerState(scanner *bufio.Scanner) ContainerState {
	// Read lines into a stack so we can replay them backwards
	var lines Stack[string]
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		lines.Push(scanner.Text())
	}

	// First line only indicates how many stacks there are
	firstLine, _ := lines.Pop()
	numStacks := (len(firstLine) + 1) / 4
	println(numStacks)

	state := ContainerState{stacks: make([]Stack[rune], numStacks)}
	for line, more := lines.Pop(); more; line, more = lines.Pop() {
		for i := 0; i < numStacks; i++ {
			ch := rune(line[i*4+1])
			if ch != ' ' {
				state.stacks[i].Push(ch)
			}
		}
	}
	return state
}

func parseInt(s string) uint {
	i, _ := strconv.ParseInt(s, 10, 32)
	return uint(i)
}
