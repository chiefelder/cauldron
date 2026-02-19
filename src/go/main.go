package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Cauldron struct {
	Id     int
	Lid    string
	Color  string
	Symbol string
	Count  string
}

type Triplet struct {
	C1 Cauldron
	C2 Cauldron
	C3 Cauldron
}

type CauldronSolution struct {
	Result       []Triplet
	Permutations int
}

func main() {
	flag.Parse()
	files := flag.Args()
	var solution CauldronSolution
	for _, file := range files {
		grid, err := parseFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		} else {
			printGrid(grid)
			solution = solveGrid(grid)
		}
		for _, triplet := range solution.Result {
			fmt.Printf("Result: %d, %d, %d\n", triplet.C1.Id, triplet.C2.Id, triplet.C3.Id)
			fmt.Printf("Permutations: %d\n", solution.Permutations)
		}
	}
}

func parseFile(filename string) ([]Cauldron, error) {
	var err error
	var ix int
	grid := make([]Cauldron, 0, 12)
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cauldron, err := parseLine(line)
		if err != nil {
			return nil, err
		}
		ix += 1
		cauldron.Id = ix
		grid = append(grid, cauldron)
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return grid, nil
}

func parseLine(line string) (Cauldron, error) {
	token := strings.Split(line, ",")
	if len(token) != 4 {
		return Cauldron{}, errors.New("Invalid file format!")
	}
	cauldron := Cauldron{
		Lid:    token[0],
		Color:  token[1],
		Symbol: token[2],
		Count:  token[3],
	}
	return cauldron, nil
}

func solveGrid(grid []Cauldron) CauldronSolution {
	length := len(grid)
	solution := CauldronSolution{
		Result:       make([]Triplet, 0, 1),
		Permutations: 0,
	}
	for ix := 0; ix < length-2; ix += 1 {
		for jx := ix + 1; jx < length-1; jx += 1 {
			for kx := jx + 1; kx < length; kx += 1 {
				triplet := Triplet{
					C1: grid[ix],
					C2: grid[jx],
					C3: grid[kx],
				}
				solution.Permutations += 1
				if checkCompliance(triplet) {
					solution.Result = append(solution.Result, triplet)
				}
			}
		}
	}
	return solution
}

func checkCompliance(triplet Triplet) bool {
	if attrMatches(triplet.C1.Lid, triplet.C2.Lid, triplet.C3.Lid) &&
		attrMatches(triplet.C1.Color, triplet.C2.Color, triplet.C3.Color) &&
		attrMatches(triplet.C1.Symbol, triplet.C2.Symbol, triplet.C3.Symbol) &&
		attrMatches(triplet.C1.Count, triplet.C2.Count, triplet.C3.Count) {
		return true
	}
	return false
}

func attrMatches(s1, s2, s3 string) bool {
	if (s1 == s2 && s1 == s3) || (s1 != s2 && s1 != s3 && s2 != s3) {
		return true
	}
	return false
}

func printGrid(grid []Cauldron) {
	for _, cauldron := range grid {
		fmt.Printf("%d: Lid: %s, Color: %s, Symbol: %s, Count: %s\n",
			cauldron.Id, cauldron.Lid, cauldron.Color, cauldron.Symbol, cauldron.Count)
	}
}
