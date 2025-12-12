package main

import (
	"maps"
	"testing"
)

func TestDictionary(t *testing.T) {
	// create a table of test cases
	table := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{
			name: "Empty stirng",
			input:    "",
			expected: map[string]int{},
		},

		{
			name: "Complex test",
			input: "solo_dolo,solo,Solo,dolo,Dolo-so-so-so_SO,So010101",
			expected: map[string]int{
				"solo": 3,
				"dolo": 3,
				"so": 5,
			},
		},

		{
			name: "Numbers",
			input: "123556879%#&H^^*(*)(*&^%h2H8l3L5O06a2)",
			expected: map[string]int{
				"h": 3,
				"l": 2,
				"a": 1,
				"o": 1,
			},
		},

		{
			name: "Complex test",
			input: "solo_dolo,solo,Solo",
			expected: map[string]int{
				"solo": 3,
				"dolo": 1,
			},
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {

			got := dictionary(tt.input)
			if  !maps.Equal(tt.expected, got){
				t.Errorf("Falied\ngot: %v\nexpected: %v", got, tt.expected)
			}
		})
	}
}
