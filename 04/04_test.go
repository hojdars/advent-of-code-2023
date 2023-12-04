package main

import (
	"reflect"
	"testing"
)

func TestParseLine(t *testing.T) {
	line := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	cardNumber, winning, selected := ParseLine(line)
	compareInt(t, cardNumber, 1)
	compareMap(t, winning, map[int]struct{}{41: {}, 48: {}, 83: {}, 86: {}, 17: {}})
	compareLine(t, selected, []int{83, 86, 6, 31, 17, 9, 48, 53})
}

func compareInt(t testing.TB, got int, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func compareLine(t testing.TB, numbers, correctNumbers []int) {
	t.Helper()
	if !reflect.DeepEqual(correctNumbers, numbers) {
		t.Errorf("got %v, want %v", numbers, correctNumbers)
	}
}
func compareMap(t testing.TB, numbers, correctNumbers map[int]struct{}) {
	t.Helper()
	if !reflect.DeepEqual(correctNumbers, numbers) {
		t.Errorf("got %v, want %v", numbers, correctNumbers)
	}
}
