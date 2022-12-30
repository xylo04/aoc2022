package main

import (
	"bufio"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/juliangruber/go-intersect/v2"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Group", "Intersection", "Priority"})

	acc := 1
	pri := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		e1 := scanner.Text()
		scanner.Scan()
		e2 := scanner.Text()
		scanner.Scan()
		e3 := scanner.Text()

		i1 := intersect.SimpleGeneric([]rune(e1), []rune(e2))
		i2 := intersect.SimpleGeneric([]rune(e2), []rune(e3))
		i := intersect.SimpleGeneric(i1, i2)[0]
		p := priority(int(i))
		t.AppendRows([]table.Row{{acc, string(i), p}})
		pri += p
		acc++
	}
	t.AppendFooter(table.Row{"", "", pri})
	t.Render()
}

func priority(ch int) int {
	if ch >= 'a' && ch <= 'z' {
		return ch - int('a') + 1
	}
	if ch >= 'A' && ch <= 'Z' {
		return ch - int('A') + 27
	}
	panic("ch out of range")
}
