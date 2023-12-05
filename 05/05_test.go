package main

import (
	"reflect"
	"testing"
)

func TestParseLine(t *testing.T) {
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
