package input_test

import (
	"testing"
	"github.com/Kuroashi1995/go-pokedex/cmd/input"
)


func TestCleanInput(t *testing.T) {
	//define table struct
	cases := []struct {
		input		string
		expected	[]string
	} {
		{
		input: " hello world ",
		expected: []string{"hello", "world"},
		},
		{
		input: "helloworld ",
		expected: []string{"helloworld"},
		},
		{
		input: "Hello World",
		expected: []string{"hello", "world"},
		},
		{
		input: "HELLO WORLD",
		expected: []string{"hello", "world"},
		},
		{
		input: "hello   world",
		expected: []string{"hello", "world"},
		},

	}

	for _, c := range cases {
		actual := input.CleanInput(c.input)
		if e, a := len(c.expected), len(actual); e != a {
			t.Errorf("expected length and actual length mismatch")
			t.Fatalf("expected: %v, got: %v", e, a)
		}
		for i := range actual {
			if expectedWord, actualWord := c.expected[i], actual[i]; expectedWord != actualWord {
				t.Errorf("expected word and actual word mismatch")
				t.Fatalf("expected: %v, got: %v", expectedWord, actualWord)
			}
		}
	}
}
