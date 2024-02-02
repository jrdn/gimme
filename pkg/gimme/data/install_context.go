package data

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/jrdn/goutil/log"

	"github.com/jrdn/gimme/pkg/condition"
)

type InstallContext map[string]any

type InstallContextKey string

const (
	SpecKey       = "SPEC"
	InstallDirKey = "INSTALL_DIR"
	PkgDirKey     = "PKG_DIR"
	OSKey         = "OS"
	ArchKey       = "ARCH"
	HomeKey       = "HOME"
)

func (c InstallContext) ConditionVars() *condition.Vars {
	vars := condition.NewVars()

	for _, env := range os.Environ() {
		parts := strings.SplitN(env, "=", 2)
		vars.Set(parts[0], parts[1])
	}

	for k, v := range c {
		vars.Set(k, v)
	}

	return vars
}

func (c InstallContext) SetEnv() {
	l := log.Get()
	for k, v := range c {
		val := fmt.Sprintf("%v", v)
		l.Trace().Str("k", k).Str("v", val).Msg("setting up env")
		err := os.Setenv(k, val)
		if err != nil {
			return
		}
	}
}

func GetInstallContext(spec, installDir string) InstallContext {
	c := InstallContext{
		SpecKey:       spec,
		OSKey:         runtime.GOOS,
		ArchKey:       runtime.GOARCH,
		InstallDirKey: installDir,
		PkgDirKey:     GetPkgDir(),
	}

	if home, err := os.UserHomeDir(); err == nil {
		c[HomeKey] = home
	}

	return c
}

func GetPkgDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	pkgDir := path.Join(home, ".config", "gimme", "pkg")
	err = os.MkdirAll(pkgDir, 0o755)
	if err != nil {
		panic(err)
	}
	return pkgDir
}
