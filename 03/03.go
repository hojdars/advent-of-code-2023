package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x, y int
}

type Number struct {
	value  int
	coords Coord
	len    int
}

func SkipEmpty(line string, start int) (nextNonEmpty int) {
	notEmpty := func(c rune) bool {
		return c != '.'
	}
	nextPos := strings.IndexFunc(line[start:], notEmpty)
	if nextPos == -1 {
		return -1
	} else {
		return nextPos + start
	}
}

func ParseNumber(line string, start int) (result, nextEmpty, length int) {
	notNumber := func(c rune) bool {
		return c < '0' || c > '9'
	}
	nextPos := strings.IndexFunc(line[start:], notNumber)
	var numberString string
	if nextPos == -1 {
		numberString = line[start:]
		nextEmpty = -1
	} else {
		numberString = line[start : start+nextPos]
		nextEmpty = start + nextPos
	}

	result, err := strconv.Atoi(numberString)
	if err != nil {
		errString := fmt.Sprintf("cannot parse number, str=%s", numberString)
		panic(errString)
	}

	length = len(numberString)

	return
}

func ParseLine(line string, y int) (symbols []Coord, numbers []Number) {
	isNumber := func(c byte) bool { return c >= '0' && c <= '9' }

	pos := 0
	for pos != -1 && pos < len(line) {
		if isNumber(line[pos]) {
			num, next, length := ParseNumber(line, pos)
			numbers = append(numbers, Number{num, Coord{pos, y}, length})
			pos = next
		} else if line[pos] == '.' {
			pos = SkipEmpty(line, pos)
		} else {
			symbols = append(symbols, Coord{pos, y})
			pos++
		}
	}
	return
}

func GetNeighbours(target Coord) [8]Coord {
	return [8]Coord{
		{target.x + 1, target.y},
		{target.x + 1, target.y + 1},
		{target.x, target.y + 1},
		{target.x - 1, target.y + 1},
		{target.x - 1, target.y},
		{target.x - 1, target.y - 1},
		{target.x, target.y - 1},
		{target.x + 1, target.y - 1},
	}
}

func GetNeighbourCoords(target Coord, len int) map[Coord]struct{} {
	result := make(map[Coord]struct{}, 0)

	for i := 0; i < len; i++ {
		c := Coord{target.x + i, target.y}
		neighbours := GetNeighbours(c)
		for _, n := range neighbours {
			result[n] = struct{}{}
		}
	}

	for i := 0; i < len; i++ {
		c := Coord{target.x + i, target.y}
		delete(result, c)
	}

	return result
}

func main() {
	file, err := os.Open("input/input")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	symbolMap := make(map[Coord]struct{}, 0)
	numbersList := make([]Number, 0)

	result := 0
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		symbols, numbers := ParseLine(line, lineNumber)
		for _, s := range symbols {
			symbolMap[s] = struct{}{}
		}
		numbersList = append(numbersList, numbers...)
		lineNumber++
	}

	for _, number := range numbersList {
		neighbours := GetNeighbourCoords(number.coords, number.len)
		isPart := false
		for n := range neighbours {
			_, ok := symbolMap[n]
			if ok {
				isPart = true
				break
			}
		}

		if isPart {
			result += number.value
		}
	}

	fmt.Printf("%d\n", result)
}
