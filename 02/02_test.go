package main

import (
	"reflect"
	"testing"
)

func TestLineParse(t *testing.T) {
	t.Run("line parse", func(t *testing.T) {
		got := ParseLine("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
		want := Game{1, []int{4, 1, 0}, []int{0, 2, 2}, []int{3, 6, 0}}

		if !reflect.DeepEqual(got.id, want.id) {
			t.Errorf("got %v, want %v", got.id, want.id)
		}
		if !reflect.DeepEqual(got.red, want.red) {
			t.Errorf("got %v, want %v", got.red, want.red)
		}
		if !reflect.DeepEqual(got.green, want.green) {
			t.Errorf("got %v, want %v", got.green, want.green)
		}
		if !reflect.DeepEqual(got.blue, want.blue) {
			t.Errorf("got %v, want %v", got.blue, want.blue)
		}
	})

}
