package base

import (
	"errors"
	"fmt"

	"github.com/gimme-repos/gimme/pkg/condition"
	"github.com/gimme-repos/gimme/pkg/gimme/data"
)

type Step struct {
	Name      string `json:"name,omitempty"`
	Type      string `json:"type"`
	When      string `json:"when"`
	ErrorWhen string `json:"when!"`
}

func (s Step) ShouldRun() (bool, error) {
	installContext := data.GetInstallContext("", "")
	vars := installContext.ConditionVars()

	stmt, err := condition.Parse(s.ErrorWhen)
	if err != nil {
		return false, err
	}
	if !stmt.Evaluate(vars) {
		return false, errors.New("'when!' condition failed")
	}

	stmt, err = condition.Parse(s.When)
	if err != nil {
		return false, err
	}

	return stmt.Evaluate(vars), nil
}

func (s Step) GetName() string {
	return s.Name
}

func (s Step) GetType() string {
	return s.Type
}

func (s Step) String() string {
	name := s.GetName()
	if name == "" {
		name = s.GetType()
	}

	return fmt.Sprintf("%s when: %q, when!: %q", name, s.When, s.ErrorWhen)
}
