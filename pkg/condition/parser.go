package condition

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/j13g/goutil/regex"
)

var parseError = errors.New("parsing failed")

func parseErr(msg, context string) error {
	return fmt.Errorf("%w: %s\nINPUT: %s\n", parseError, msg, context)
}

func Parse(input string) (Statement, error) {
	if strings.TrimSpace(input) == "" {
		return &alwaysTrueAtom{}, nil
	}

	s, remaining, err := parse(input)
	if err != nil {
		return nil, err
	}

	if len(remaining) != 0 {
		return nil, parseErr("parsing incorrectly had leftover data", remaining)
	}

	return s, nil
}

func parse(input string) (Statement, string, error) {
	input = strings.TrimSpace(input)
	for _, bd := range blockDelimiters {
		if bd.Prefix(input) {
			return parseBlock(bd, input)
		}
	}

	return parseAtom(input)
}

var atomRegex = regexp.MustCompile(`^\s*(?P<var>\w*)\s+(?P<op>\w*)\s+(?P<val>(\".*\")|\S*)\s*(?P<remaining>.*)$`)

func parseAtom(input string) (*atom, string, error) {
	matches := regex.Match(atomRegex, input)
	if len(matches) == 0 {
		return nil, "", parseErr("failed to match atom", input)
	}

	remaining := matches[0]["remaining"]

	a := &atom{
		varName:  strings.ToLower(matches[0]["var"]),
		operator: parseOperator(matches[0]["op"]),
		value:    parseValue(matches[0]["val"]),
	}

	return a, remaining, nil
}

func parseValue(value string) any {
	if i, err := strconv.Atoi(value); err == nil {
		return i
	}

	if f, err := strconv.ParseFloat(value, 64); err == nil {
		return f
	}

	if b, err := strconv.ParseBool(value); err == nil {
		return b
	}

	if strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`) {
		return strings.Trim(value, `"`)
	}

	return value
}

func parseBlock(delim blockDelimiter, input string) (*block, string, error) {
	matches := regex.Match(delim.matcher, input)
	if len(matches) == 0 {
		return nil, "", parseErr("failed to match block", input)
	}
	topLevelRemaining := matches[0]["remaining"]

	b := &block{
		blockType: delim.block,
	}

	var s Statement
	var err error
	remaining := matches[0]["block"]
	lastLen := len(remaining)
	for remaining != "" {
		s, remaining, err = parse(remaining)
		if err != nil {
			return nil, "", err
		}
		curLen := len(remaining)
		if curLen == lastLen {
			// running through the parser again hasn't consumed any of the remaining data, so something is wrong
			return nil, "", parseErr("remaining block data not parsing", remaining)
		}
		lastLen = curLen
		b.contents = append(b.contents, s)
	}

	return b, topLevelRemaining, nil
}
