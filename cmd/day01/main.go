package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	calories := make([]int, 0)
	scanner := bufio.NewScanner(file)
	elf := 0
	calories = append(calories, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			elf++
			calories = append(calories, 0)
		}
		cal, _ := strconv.ParseUint(text, 10, 32)
		calories[elf] += int(cal)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(calories)))
	fmt.Printf("The elf with the most food has %d calories\n", calories[0])
	fmt.Printf("The top three elves are carrying a total of %d calories\n", sum(calories[0:3]))
}

func sum(inputs []int) int {
	var acc = 0
	for _, v := range inputs {
		acc += v
	}
	return acc
}
