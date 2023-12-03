package main

import (
	"reflect"
	"testing"
)

func TestSkipEmpty(t *testing.T) {
	line := "467..114.."
	compare(t, SkipEmpty(line, 3), 5)
	compare(t, SkipEmpty(line, 0), 0)
	compare(t, SkipEmpty(line, 8), -1)

	pos := 0
	correctResults := []int{0, 1, 2, 5, 5, 6, 7, -1}
	for _, correct := range correctResults {
		next := SkipEmpty(line, pos)
		compare(t, next, correct)
		if pos == next {
			pos++
		} else {
			pos = next
		}
	}
}

func TestParseNumber(t *testing.T) {
	line := "..35..633."
	num, nextEmpty, len := ParseNumber(line, 2)
	compare(t, num, 35)
	compare(t, nextEmpty, 4)
	compare(t, len, 2)
	num, nextEmpty, len = ParseNumber(line, 6)
	compare(t, num, 633)
	compare(t, nextEmpty, 9)
	compare(t, len, 3)
}

func TestParseLine(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		line := "..35.&..633.*"
		symbols, numbers, gears := ParseLine(line, 0)
		correctSymbols := []Coord{{5, 0}, {12, 0}}
		correctNumbers := []Number{{35, Coord{2, 0}, 2}, {633, Coord{8, 0}, 3}}
		compareLine(t, symbols, correctSymbols, numbers, correctNumbers)

		correctGears := []Coord{{12, 0}}
		if !reflect.DeepEqual(gears, correctGears) {
			t.Errorf("got %v, want %v", gears, correctGears)
		}
	})

	t.Run("test 2", func(t *testing.T) {
		line := "...$.*...."
		symbols, numbers, gears := ParseLine(line, 0)
		correctSymbols := []Coord{{3, 0}, {5, 0}}
		var correctNumbers []Number = nil
		compareLine(t, symbols, correctSymbols, numbers, correctNumbers)

		correctGears := []Coord{{5, 0}}
		if !reflect.DeepEqual(gears, correctGears) {
			t.Errorf("got %v, want %v", gears, correctGears)
		}
	})

	t.Run("test 3", func(t *testing.T) {
		line := ".....+.58."
		symbols, numbers, _ := ParseLine(line, 0)
		correctSymbols := []Coord{{5, 0}}
		correctNumbers := []Number{{58, Coord{7, 0}, 2}}
		compareLine(t, symbols, correctSymbols, numbers, correctNumbers)
	})

	t.Run("test 4", func(t *testing.T) {
		line := "..592....."
		symbols, numbers, _ := ParseLine(line, 0)
		var correctSymbols []Coord = nil
		correctNumbers := []Number{{592, Coord{2, 0}, 3}}
		compareLine(t, symbols, correctSymbols, numbers, correctNumbers)
	})

	t.Run("test 5", func(t *testing.T) {
		line := "1/592=...."
		symbols, numbers, _ := ParseLine(line, 0)
		correctSymbols := []Coord{{1, 0}, {5, 0}}
		correctNumbers := []Number{{1, Coord{0, 0}, 1}, {592, Coord{2, 0}, 3}}
		compareLine(t, symbols, correctSymbols, numbers, correctNumbers)
	})

	t.Run("test 6", func(t *testing.T) {
		line := ".....+.58"
		symbols, numbers, _ := ParseLine(line, 0)
		correctSymbols := []Coord{{5, 0}}
		correctNumbers := []Number{{58, Coord{7, 0}, 2}}
		compareLine(t, symbols, correctSymbols, numbers, correctNumbers)
	})
}

func TestGetNeighbourCoords(t *testing.T) {
	t.Run("length 1", func(t *testing.T) {
		got := GetNeighbourCoords(Coord{1, 1}, 1)
		want := make(map[Coord]struct{}, 0)
		for _, c := range GetNeighbours(Coord{1, 1}) {
			want[c] = struct{}{}
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("length 2", func(t *testing.T) {
		got := GetNeighbourCoords(Coord{1, 1}, 2)
		want := map[Coord]struct{}{
			{0, 0}: {}, {0, 1}: {}, {0, 2}: {}, {1, 0}: {}, {1, 2}: {}, {2, 0}: {}, {2, 2}: {}, {3, 0}: {}, {3, 1}: {}, {3, 2}: {},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func compare(t testing.TB, got int, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func compareLine(t testing.TB, symbols, correctSymbols []Coord, numbers, correctNumbers []Number) {
	t.Helper()
	if !reflect.DeepEqual(correctSymbols, symbols) {
		t.Errorf("got %v, want %v", symbols, correctSymbols)
	}
	if !reflect.DeepEqual(correctNumbers, numbers) {
		t.Errorf("got %v, want %v", numbers, correctNumbers)
	}
}
