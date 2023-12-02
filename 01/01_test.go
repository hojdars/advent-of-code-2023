package main

import "testing"

func TestLineRead(t *testing.T) {
	t.Run("two numbers", func(t *testing.T) {
		compare(t, "1abc2", 12)
		compare(t, "pqr3stu8vwx", 38)
		compare(t, "a1b2c3d4e5f", 15)
	})

	t.Run("one number", func(t *testing.T) {
		compare(t, "treb7uchet", 77)
	})

	t.Run("numbers only", func(t *testing.T) {
		compare(t, "739", 79)
	})

	t.Run("no numbers", func(t *testing.T) {
		input := "abcd"
		got, err := ParseLine(input)
		if err == nil {
			t.Errorf("wanted error, got nil and value %d", got)
		}
	})

	t.Run("numbers as words mixed in", func(t *testing.T) {
		compare(t, "one2", 12)
		compare(t, "two1nine", 29)
		compare(t, "eightwothree", 83)
		compare(t, "abcone2threexyz", 13)
		compare(t, "xtwone3four", 24)
		compare(t, "4nineeightseven2", 42)
		compare(t, "zoneight234", 14)
		compare(t, "7pqrstsixteen", 76)
		compare(t, "1onebfjtdslkdbthree4jvvonezqdthreesrghnnbsix", 16)
		compare(t, "oneight", 18)
		compare(t, "oneightwoneight", 18)
	})
}

func compare(t *testing.T, input string, want int) {
	t.Helper()
	got, err := ParseLine(input)

	if err != nil {
		t.Fatal("number not found")
	}

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
