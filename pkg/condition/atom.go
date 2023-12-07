package condition

import (
	"strings"

	"github.com/stretchr/testify/assert"
)

func NewAtom(varName string, op operator, expected any) *atom {
	return &atom{
		varName:  strings.ToLower(varName),
		operator: op,
		value:    expected,
	}
}

type atom struct {
	varName  string
	operator operator
	value    any
}

func (a *atom) Evaluate(vars *Vars) bool {
	// stealing stretchr/testify for this is gross, but easy
	varValue := vars.Get(a.varName)

	switch a.operator {
	case Equal:
		return assert.Equal(ft, varValue, a.value)
	case NotEqual:
		return assert.NotEqual(ft, varValue, a.value)
	case LessThan:
		return assert.Less(ft, varValue, a.value)
	case LessOrEqual:
		return assert.LessOrEqual(ft, varValue, a.value)
	case GreaterThan:
		return assert.Greater(ft, varValue, a.value)
	case GreaterOrEqual:
		return assert.GreaterOrEqual(ft, varValue, a.value)
	}
	return false
}

func (a *atom) eq(m map[string]any) bool {
	return assert.ObjectsAreEqualValues(a.value, m[a.varName])
}

// a fake testing.T implementation to make it easier to hijack stretchr/testify because I am a bad person
type fakeT struct{}

func (f fakeT) Errorf(format string, args ...interface{}) {}

var ft = fakeT{}

type alwaysTrueAtom struct{}

func (a alwaysTrueAtom) Evaluate(vars *Vars) bool {
	return true
}
