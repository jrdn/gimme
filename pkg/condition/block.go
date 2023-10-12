package condition

import (
	"regexp"
	"strings"
)

type blockType int

const (
	And blockType = iota
	Or
	Comment
)

type blockDelimiter struct {
	block   blockType
	begin   string
	end     string
	matcher *regexp.Regexp
}

func (bd blockDelimiter) Prefix(x string) bool {
	return strings.HasPrefix(x, bd.begin)
}

var blockDelimiters = map[blockType]blockDelimiter{
	And:     {And, "[", "]", regexp.MustCompile(`^\[(?P<block>[\w\W]*?)\](?P<remaining>[\w\W]*)$`)},
	Or:      {Or, "{", "}", regexp.MustCompile(`^\{(?P<block>[\w\W]*?)\}(?P<remaining>[\w\W]*)$`)},
	Comment: {Comment, "(", ")", regexp.MustCompile(`^\((?P<block>[\w\W]*?)\)`)},
}

var blockStartDelimiter = []string{"{", "[", "("}

type block struct {
	blockType blockType
	contents  []Statement
}

func (b *block) Evaluate(m *Vars) bool {
	if b.blockType == And {
		return b.evaluateAnd(m)
	} else if b.blockType == Or {
		return b.evaluateOr(m)
	} else {
		panic("unknown block type")
	}
}

func (b *block) evaluateAnd(m *Vars) bool {
	for _, c := range b.contents {
		if !c.Evaluate(m) {
			return false
		}
	}
	return true
}

func (b *block) evaluateOr(m *Vars) bool {
	for _, c := range b.contents {
		if c.Evaluate(m) {
			return true
		}
	}
	return false
}
