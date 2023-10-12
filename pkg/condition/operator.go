package condition

import "strings"

type operator string

const (
	Equal          operator = "EQ"
	NotEqual       operator = "NE"
	LessThan       operator = "LT"
	LessOrEqual    operator = "LE"
	GreaterThan    operator = "GT"
	GreaterOrEqual operator = "GE"
)

func parseOperator(x string) operator {
	return operator(strings.ToUpper(x))
}
