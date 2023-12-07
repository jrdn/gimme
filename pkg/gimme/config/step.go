package config

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gimme-repos/gimme/pkg/gimme/data"
	"github.com/gimme-repos/gimme/pkg/gimme/step/base"
	"github.com/gimme-repos/gimme/pkg/gimme/step/pwsh"
	"github.com/gimme-repos/gimme/pkg/gimme/step/require_windows_admin"
	"github.com/gimme-repos/gimme/pkg/gimme/step/taskfile"
)

type Step interface {
	fmt.Stringer

	GetName() string
	GetType() string

	Exec(ctx context.Context, vars data.InstallContext) error
	ShouldRun() (bool, error)
}

type StepFactory func() Step

var stepTypes = map[string]StepFactory{
	"task":                  func() Step { return taskfile.New() },
	"pwsh":                  func() Step { return pwsh.New() },
	"windows-require-admin": func() Step { return require_windows_admin.New() },
}

type Steps []Step

func (s *Steps) UnmarshalJSON(data []byte) error {
	var raw []json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	for _, rawStep := range raw {
		base := &base.Step{}
		err := json.Unmarshal(rawStep, base)
		if err != nil {
			return err
		}

		if stepFactory, ok := stepTypes[base.Type]; ok {
			step := stepFactory()
			err := json.Unmarshal(rawStep, step)
			if err != nil {
				return err
			}
			*s = append(*s, step)
		} else {
			return fmt.Errorf("unknown step type: %s", base.Type)
		}
	}

	return nil
}
