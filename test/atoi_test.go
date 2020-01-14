package tests

import (
	"testing"

	student ".."
)

func TestAtoi(t *testing.T) {
	tables := []struct {
		s      string
		answer int
	}{
		{"", 0},
		{"-", 0},
		{"--123", 0},
		{"1", 1},
		{"-3", -3},
		{"8292", 8292},
		{"000000123", 123},
		{"9223372036854775807", 9223372036854775807},
		{"-9223372036854775808", -9223372036854775808},
		{"000000000000000009223372036854775807", 9223372036854775807},
	}

	for _, table := range tables {
		res := student.Atoi(table.s)
		if res != table.answer {
			t.Errorf("Atoi (%s): [your answer - (%d)] != [(%d) - test's answer]", table.s, res, table.answer)
		}
	}
}
