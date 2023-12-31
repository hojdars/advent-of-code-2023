package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseLine(line string) (cardNumber int, winning map[int]struct{}, picked []int) {
	split := strings.Split(line, " ")
	winning = make(map[int]struct{}, 0)
	picked = make([]int, 0)

	leftSide := true
	for _, v := range split[1:] {
		if v == "" {
			continue
		}

		if v == "|" {
			leftSide = false
			continue
		}

		if v[len(v)-1] == ':' {
			number, err := strconv.Atoi(v[:len(v)-1])
			if err != nil {
				panic("cannot convert card number")
			}
			cardNumber = number
			continue
		}

		number, err := strconv.Atoi(v)
		if err != nil {
			panic("cannot convert lottery number")
		}
		if leftSide {
			winning[number] = struct{}{}
		} else {
			picked = append(picked, number)
		}
	}
	return
}

func GetNumberOfWins(winning map[int]struct{}, picked []int) int {
	numOfHits := 0
	for _, n := range picked {
		_, ok := winning[n]
		if ok {
			numOfHits++
		}
	}
	return numOfHits
}

func SolvePartOne() {
	file, err := os.Open("input/input")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		_, winning, picked := ParseLine(line)
		numOfHits := GetNumberOfWins(winning, picked)
		if numOfHits > 0 {
			result += 1 << (numOfHits - 1)
		}
	}
	fmt.Printf("4-1 result=%d\n", result)
}

func SolvePartTwo() {
	file, err := os.Open("input/input")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	result := 0
	cards := make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()
		cardNumber, winning, picked := ParseLine(line)
		numOfHits := GetNumberOfWins(winning, picked)

		cards[cardNumber] += 1
		result += cards[cardNumber]

		for i := 1; i <= numOfHits; i++ {
			cards[cardNumber+i] += cards[cardNumber]
		}
	}
	fmt.Printf("4-2 result=%d\n", result)
}

func main() {
	SolvePartOne()
	SolvePartTwo()
}
