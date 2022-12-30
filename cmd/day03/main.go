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
	t.AppendHeader(table.Row{"#", "Compartment 1", "Compartment 2", "Intersection", "Priority"})

	acc := 1
	pri := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pack := scanner.Text()
		c1 := pack[:len(pack)/2]
		c2 := pack[len(pack)/2:]
		i := intersect.SimpleGeneric([]rune(c1), []rune(c2))[0]
		p := priority(int(i))
		t.AppendRows([]table.Row{{acc, c1, c2, string(i), p}})
		pri += p
		acc++
	}
	t.AppendFooter(table.Row{"", "", "", "", pri})
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
