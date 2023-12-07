package taskfile

import (
	"context"
	"fmt"
	"os"

	"github.com/go-task/task/v3"
	"github.com/go-task/task/v3/taskfile"

	"github.com/j13g/goutil/log"

	"github.com/gimme-repos/gimme/pkg/gimme/data"
	"github.com/gimme-repos/gimme/pkg/gimme/step/base"
)

func New() *Step {
	return &Step{}
}

type Step struct {
	base.Step
	Target string `json:"target"`
}

func (s *Step) Exec(ctx context.Context, installContext data.InstallContext) error {
	l := log.Get()

	target := "default"
	if s.Target != "" {
		target = s.Target
	}

	installContext.SetEnv()

	dir := fmt.Sprintf("%v", installContext[data.InstallDirKey])
	err := os.Chdir(dir)
	if err != nil {
		return err
	}

	// TODO detect taskfile.yaml, taskfile.yml, Taskfile.yaml, etc
	e := task.Executor{
		Verbose:     true,
		AssumeYes:   true,
		Dir:         dir,
		Entrypoint:  "Taskfile.yml",
		Concurrency: 1,
		Silent:      false,

		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}
	err = e.Setup()
	if err != nil {
		return err
	}

	l.Trace().
		Str("target", target).
		Str("dir", dir).
		Interface("spec", installContext[data.SpecKey]).
		Msg("running task")
	return e.Run(ctx, taskfile.Call{
		Task:   target,
		Silent: false,
		Direct: true,
	})
}
