package day_one

import "errors"

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
