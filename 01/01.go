package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func checkForWord(line string, i int, find string) bool {
	return len(line) >= i+len(find) && line[i:i+len(find)] == find
}

func updateFirstLast(value, firstIn int) (first, last int) {
	if firstIn == -1 {
		first = value
		last = value
	} else {
		first = firstIn
		last = value
	}
	return
}

var letters = [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func ParseLine(line string) (int, error) {
	first := -1
	last := -1

	for i := 0; i < len(line); {
		if line[i] >= '0' && line[i] <= '9' {
			first, last = updateFirstLast(int(line[i]-'0'), first)
			i += 1
			continue
		} else {
			found := false
			for index, number := range letters {
				if checkForWord(line, i, number) {
					first, last = updateFirstLast(index+1, first)
					i += len(number) - 1
					found = true
					break
				}
			}
			if !found {
				i += 1
			}
		}
	}

	if first == -1 {
		return -1, errors.New("no numbers found")
	} else {
		return 10*first + last, nil
	}
}

func main() {
	file, err := os.Open("input/input.txt")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		num, err := ParseLine(line)
		if err != nil {
			fmt.Println("error encountered, exit")
			return
		}
		sum += num
	}

	fmt.Println(sum)
}
