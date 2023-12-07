package require_windows_admin

import (
	"context"
	"errors"
	"os/exec"
	"strings"

	"github.com/gimme-repos/gimme/pkg/gimme/data"
	"github.com/gimme-repos/gimme/pkg/gimme/step/base"
)

func New() *Step {
	return &Step{}
}

type Step struct {
	base.Step
}

func (r Step) Exec(ctx context.Context, vars data.InstallContext) error {
	cmd := exec.Command("net", "session")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	if strings.Contains(string(out), "Access is denied") {
		return errors.New("not windows admin")
	}

	return nil
}
