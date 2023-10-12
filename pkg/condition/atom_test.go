package condition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestData() *Vars {
	return NewVarsWithData(map[string]any{
		"OS":           "windows",
		"Arch":         "amd64",
		"int1":         1,
		"float_pi":     3.14159,
		"string_space": "hello world",
	})
}

func TestAtomEvaluate(t *testing.T) {
	tests := []struct {
		name     string
		input    *atom
		expected bool
	}{
		{name: "int eq", input: &atom{varName: "int1", operator: Equal, value: 1}, expected: true},
		{name: "bad int eq", input: &atom{varName: "int1", operator: Equal, value: 2}, expected: false},
		{name: "int ne", input: &atom{varName: "int1", operator: NotEqual, value: 2}, expected: true},

		{name: "string eq", input: &atom{varName: "OS", operator: Equal, value: "windows"}, expected: true},
		{name: "bad string eq", input: &atom{varName: "OS", operator: Equal, value: "linux"}, expected: false},
		{name: "string ne", input: &atom{varName: "OS", operator: NotEqual, value: "linux"}, expected: true},
		{name: "string with spaces eq", input: &atom{varName: "string_space", operator: Equal, value: "hello world"}, expected: true},

		{name: "float_gt", input: &atom{varName: "float_pi", operator: GreaterThan, value: 3.0}, expected: true},
		{name: "float_lt", input: &atom{varName: "float_pi", operator: LessThan, value: 3.2}, expected: true},
	}
	vars := getTestData()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.input.Evaluate(vars)
			assert.Equal(t, test.expected, result)
		})
	}
}
