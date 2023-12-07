package gimme

import (
	"context"
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/j13g/goutil/log"

	"github.com/gimme-repos/gimme/pkg/gimme/config"
	"github.com/gimme-repos/gimme/pkg/gimme/data"
	"github.com/gimme-repos/gimme/pkg/gimme/source"
)

const GimmeFile = "gimme.jsonnet"

type options struct {
	noDeps bool
}

type InstallOption func(o options)

func NoDeps() InstallOption {
	return func(o options) {
		o.noDeps = true
	}
}

func Install(ctx context.Context, pkgSpec string, installOptions ...InstallOption) error {
	opts := options{}
	for _, o := range installOptions {
		o(opts)
	}

	l := log.Get().With().Str("spec", pkgSpec).Logger()

	l.Info().Msg("installing")

	pkgSource, err := source.Parse(pkgSpec)
	if err != nil {
		return err
	}

	l.Trace().Msg("pulling source")
	Header("Pulling source for " + pkgSpec)
	err = pkgSource.Pull()
	if err != nil {
		return err
	}

	l.Trace().Msg("getting config")
	pkgConfig, err := pkgSource.GetConfig()
	if err != nil {
		return err
	}

	l.Trace().Str("path", pkgConfig.ConfigPath).Msg("saving config with updated spec")
	pkgConfig.Spec = pkgSpec
	err = config.Save(pkgConfig.ConfigPath, pkgConfig)
	if err != nil {
		return err
	}

	// TODO build dependency graph, break cycles, and install in topo-sorted order
	// This just installs everything in the order it is discovered. This will result in infinite loops if there are
	// circular dependencies and other unreasonable behaviors

	for _, dep := range pkgConfig.Dependencies {
		l.Info().Str("dep", dep).Msg("discovered dependency")
		err := Install(ctx, dep)
		if err != nil {
			return err
		}
	}

	installVars := data.GetInstallContext(pkgSpec, pkgSource.GetInstallDir())
	l.Debug().Interface("vars", installVars).Msg("install context")
	Header("Running install for %s", pkgSpec)

	l.Trace().Msg("running install steps")
	for i, setupStep := range pkgConfig.Setup {
		shouldRun, err := setupStep.ShouldRun()
		if err != nil {
			l.Error().Err(err).Msg("failed to determine if setup step should run")
			return err
		}
		if !shouldRun {
			l.Info().Msg("skipping setup step, 'when' condition not matched")
			continue
		}

		stepName := setupStep.GetName()
		if stepName == "" {
			stepName = fmt.Sprintf("#%d: %s", i, setupStep.GetType())
		}
		Section("Running setup step %s for package %s", stepName, pkgSpec)
		err = setupStep.Exec(ctx, installVars)
		if err != nil {
			Error("failed to execute setup step: %s", err.Error())
			l.Error().Err(err).Msg("failed to execute setup step")
			return err
		}
	}

	return nil
}

func ListInstalled() ([]*config.Config, error) {
	var found []*config.Config
	err := filepath.WalkDir(data.GetPkgDir(), func(p string, d fs.DirEntry, err error) error {
		if filepath.Base(p) == GimmeFile {
			cfg, err := config.LoadConfigFile(p)
			if err != nil {
				return err
			}
			found = append(found, cfg)
		}
		return nil
	})
	return found, err
}
