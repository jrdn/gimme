package condition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Statement
		wantErr  bool
	}{
		{name: "always true", input: "", expected: &alwaysTrueAtom{}},
		{name: "case insensitive", input: "OS eq windows", expected: &atom{varName: "os", operator: Equal, value: "windows"}},
		{name: "case insensitive", input: "oS Eq windows", expected: &atom{varName: "os", operator: Equal, value: "windows"}},
		{name: "value with spaces", input: `something EQ "long string with spaces"`, expected: &atom{varName: "something", operator: Equal, value: "long string with spaces"}},

		{name: "int value", input: "var EQ 1", expected: &atom{varName: "var", operator: Equal, value: 1}},
		{name: "bool value", input: "var EQ true", expected: &atom{varName: "var", operator: Equal, value: true}},
		{name: "float value", input: "var EQ 3.14159", expected: &atom{varName: "var", operator: Equal, value: 3.14159}},

		{name: "AND block", input: "[os eq windows arch eq amd64]", expected: &block{
			blockType: And,
			contents: []Statement{
				&atom{varName: "os", operator: Equal, value: "windows"},
				&atom{varName: "arch", operator: Equal, value: "amd64"},
			},
		}},
		{
			name: "complex nested block",
			input: `
[
  {os eq windows arch eq badvalue}
  {foo eq bar baz eq quux}
]`,
			expected: &block{
				blockType: And,
				contents: []Statement{
					&block{blockType: Or, contents: []Statement{
						&atom{"os", Equal, "windows"},
						&atom{"arch", Equal, "badvalue"},
					}},
					&block{blockType: Or, contents: []Statement{
						&atom{"foo", Equal, "bar"},
						&atom{"baz", Equal, "quux"},
					}},
				},
			},
		},
		{
			name: "complex nested block with types and long strings",
			input: `
[
  {os eq "foo bar baz" arch eq 1.2}
  {foo eq false baz eq true}
]`,
			expected: &block{
				blockType: And,
				contents: []Statement{
					&block{blockType: Or, contents: []Statement{
						&atom{"os", Equal, `foo bar baz`},
						&atom{"arch", Equal, 1.2},
					}},
					&block{blockType: Or, contents: []Statement{
						&atom{"foo", Equal, false},
						&atom{"baz", Equal, true},
					}},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := Parse(test.input)

			if test.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Equal(t, test.expected, result)
		})
	}
}
