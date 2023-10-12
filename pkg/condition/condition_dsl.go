/*
Package condition implements a simple DSL for expressing the conditions under which steps should be run.
Conditions are statements with three parts: "VARIABLE OPERATOR EXPECTED"

These simple statements can be joined to form more complex conditions.
Enclosing conditions in square brackets means all enclosed conditions must be true (AND)
Enclosing conditions in curly braces means the block is true if any statement inside it is true (OR)
Anything enclosed with parenthesis is a comment.

Whitespace is meaningless.

Example:
(this condition matches running on windows on AMD64 or ARM64)
[
  OS eq WINDOWS
  {ARCH eq AMD64 ARCH eq ARM64}
]
*/

package condition

type Statement interface {
	Evaluate(vars *Vars) bool
}
