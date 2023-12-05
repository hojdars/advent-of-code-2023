package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	seedStr        = "seed-to-soil map:"
	soilStr        = "soil-to-fertilizer map:"
	fertilizerStr  = "fertilizer-to-water map:"
	waterStr       = "water-to-light map:"
	lightStr       = "light-to-temperature map:"
	temperatureStr = "temperature-to-humidity map:"
	humidityStr    = "humidity-to-location map:"
)

const (
	seed        = iota
	soil        = iota
	fertilizer  = iota
	water       = iota
	light       = iota
	temperature = iota
	humidity    = iota
)

type Map struct {
	numbers []string
}

func main() {
	file, err := os.Open("input/input.test")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	seeds := scanner.Text()
	fmt.Println(seeds)

	maps := [7]Map{}

	status := -1
	for scanner.Scan() {
		fmt.Println(status)
		line := scanner.Text()
		if line == "" {
			continue
		}

		if status == -1 && line == seedStr {
			status = seed
		} else if status == seed && line == soilStr {
			status = soil
		} else if status == soil && line == fertilizerStr {
			status = fertilizer
		} else if status == fertilizer && line == waterStr {
			status = water
		} else if status == water && line == lightStr {
			status = light
		} else if status == light && line == temperatureStr {
			status = temperature
		} else if status == temperature && line == humidityStr {
			status = humidity
		} else {
			maps[status].numbers = append(maps[status].numbers, line)
		}
	}

	for _, m := range maps {
		fmt.Println(m)
	}
}
