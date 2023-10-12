package condition_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/gimme-repos/gimme/pkg/condition"
)

func TestConditionDSL(t *testing.T) {
	testData := condition.NewVarsWithData(map[string]any{
		"OS":         "windows",
		"Arch":       "amd64",
		"int1":       1,
		"true":       true,
		"has_spaces": "foo bar",
		"float":      3.1415926,
	})

	tests := []struct {
		input        string
		expected     bool
		wantParseErr bool
	}{
		{input: "os eq windows", expected: true},
		{input: "[os eq windows arch eq amd64]", expected: true},
		{input: "[os eq linux arch eq amd64]", expected: false},
		{input: "{os eq linux arch eq amd64}", expected: true},
		{input: "has_spaces eq \"foo bar\"", expected: true},
		{input: `
[
  {os eq windows arch eq badvalue}
  {os eq windows arch eq badvalue}
]
`, expected: true},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			s, err := condition.Parse(test.input)
			if test.wantParseErr {
				require.Error(t, err)
				return
			}
			assert.NoError(t, err)

			result := s.Evaluate(testData)
			assert.Equal(t, test.expected, result)
		})
	}
}
