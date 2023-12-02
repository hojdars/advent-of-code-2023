package day_one

import "testing"

func TestLineRead(t *testing.T) {
	t.Run("two numbers", func(t *testing.T) {
		compare(t, "1abc2", 12)
		compare(t, "pqr3stu8vwx", 38)
		compare(t, "a1b2c3d4e5f", 15)
	})

	t.Run("one number", func(t *testing.T) {
		input := "treb7uchet"
		compare(t, input, 77)
	})

	t.Run("no numbers", func(t *testing.T) {
		input := "abcd"
		got, err := ParseLine(input)
		if err == nil {
			t.Errorf("wanted error, got nil and value %d", got)
		}
	})
}

func compare(t *testing.T, input string, want int) {
	got, err := ParseLine(input)

	if err != nil {
		t.Errorf("number not found")
	}

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
