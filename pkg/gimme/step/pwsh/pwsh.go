package pwsh

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/gimme-repos/gimme/pkg/gimme/data"
	"github.com/gimme-repos/gimme/pkg/gimme/step/base"
)

func New() *Step {
	return &Step{}
}

type Step struct {
	base.Step
	Path    string `json:"path,omitempty"`
	Command string `json:"command,omitempty"`
}

func (s Step) Exec(ctx context.Context, installContext data.InstallContext) error {
	dir := fmt.Sprintf("%v", installContext[data.InstallDirKey])
	installContext.SetEnv()

	if s.Command != "" {
		cmd := exec.Command("pwsh", "-nologo", "-noprofile", "-command", s.Command)
		cmd.Dir = dir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stdout
		err := cmd.Run()
		if err != nil {
			return err
		}
	}

	if s.Path != "" {
		cmd := exec.Command("pwsh", "-nologo", "-noprofile", s.Path)
		cmd.Dir = dir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stdout
		err := cmd.Run()
		if err != nil {
			return err
		}
	}

	return nil
}
