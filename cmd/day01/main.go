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
	fmt.Printf("The elf with the most food has %d calories", calories[0])
}
