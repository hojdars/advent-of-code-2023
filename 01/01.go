package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func ParseLine(line string) (int, error) {
	first := -1
	last := -1

	for _, char := range line {
		if !(char >= '0' && char <= '9') {
			continue
		}

		if first == -1 {
			first = int(char - '0')
			last = int(char - '0')
		} else {
			last = int(char - '0')
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
