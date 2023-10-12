package source

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/gimme-repos/gimme/pkg/gimme/config"
	"github.com/gimme-repos/gimme/pkg/gimme/data"
)

type gitSource struct {
	*sourceBase

	pullURL string
	version string
}

func (g *gitSource) Pull() error {
	// TODO if packages are already cloned, pull instead
	installDir := g.GetInstallDir()
	_, err := os.Stat(filepath.Join(installDir, ".git"))
	if os.IsNotExist(err) {
		return g.gitClone()
	}

	return g.gitPull()
}

func (g *gitSource) gitPull() error {
	cmd := exec.Command("git", "fetch", "--all")
	cmd.Dir = g.GetInstallDir()
	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))

	cmd = exec.Command("git", "checkout", "origin/"+g.version)
	cmd.Dir = g.GetInstallDir()
	out, err = cmd.CombinedOutput()
	fmt.Println(string(out))

	return err
}

func (g *gitSource) gitClone() error {
	cmd := exec.Command("git", "clone", "-b", g.version, g.pullURL, g.GetInstallDir())
	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))
	return err
}

func (g *gitSource) GetInstallDir() string {
	idHash := sha256.Sum256([]byte(g.pullURL))
	id := hex.EncodeToString(idHash[:])
	installDir := path.Join(data.GetPkgDir(), id)
	err := os.MkdirAll(installDir, 0o755)
	if err != nil {
		panic(err)
	}
	return installDir
}

func (g *gitSource) GetConfig() (*config.Config, error) {
	path := filepath.Join(g.GetInstallDir(), "gimme.jsonnet")
	return config.LoadConfigFile(path)
}

func (g *gitSource) Status() SourceStatus {
	return "TODO"
}
