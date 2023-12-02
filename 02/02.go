package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	id    int
	red   []int
	green []int
	blue  []int
}

func (g Game) String() string {
	return fmt.Sprintf("Game[id=%v, lens={%v, %v, %v}]", g.id, len(g.red), len(g.green), len(g.blue))
}

func (g Game) GetMax() (result [3]int) {
	result = [3]int{0, 0, 0}
	for _, r := range g.red {
		if r > result[0] {
			result[0] = r
		}
	}
	for _, g := range g.green {
		if g > result[1] {
			result[1] = g
		}
	}
	for _, b := range g.blue {
		if b > result[2] {
			result[2] = b
		}
	}
	return
}

func ParseLine(line string) (result Game) {
	if line[0:5] != "Game " {
		panic("invalid line")
	}

	start := strings.Index(line, ":")
	idStr := line[5:start]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic("invalid game id")
	}
	result.id = id

	pulls := strings.Split(line[start+1:], ";")
	result.red = make([]int, len(pulls))
	result.green = make([]int, len(pulls))
	result.blue = make([]int, len(pulls))

	for i, pull := range pulls {
		for _, colorLine := range strings.Split(pull, ",") {
			colorSlice := strings.Trim(colorLine, " ")
			color := strings.Split(colorSlice, " ")
			num, err := strconv.Atoi(color[0])
			if err != nil {
				panic("invalid number in color")
			}
			if color[1] == "red" {
				result.red[i] = num
			} else if color[1] == "green" {
				result.green[i] = num
			} else if color[1] == "blue" {
				result.blue[i] = num
			} else {
				panic("unknown color")
			}
		}
	}

	return
}

func main() {
	file, err := os.Open("input/input")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sumPart1 := 0
	sumPart2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		game := ParseLine(line)
		maxes := game.GetMax()

		// Part 1
		if maxes[0] <= 12 && maxes[1] <= 13 && maxes[2] <= 14 {
			sumPart1 += game.id
		}

		// Part 2
		num := maxes[0] * maxes[1] * maxes[2]
		sumPart2 += num
	}

	fmt.Printf("part 1 answer=%v, part 2 answer=%v\n", sumPart1, sumPart2)
}
