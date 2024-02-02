package cmd

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/jrdn/gimme/pkg/gimme"
	"github.com/jrdn/gimme/pkg/gimme/config"
)

func Edit() *cobra.Command {
	cmd := &cobra.Command{
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			spec := args[0]

			installed, err := gimme.ListInstalled()
			if err != nil {
				return err
			}

			for _, pkg := range installed {
				if strings.Contains(pkg.Spec, spec) {
					editor, err := cmd.Flags().GetString("editor")
					if err != nil {
						return err
					}
					return edit(pkg, editor)
				}
			}
			return nil
		},
	}

	cmd.Flags().StringP("editor", "e", "", "Editor to use")
	return cmd
}

func edit(pkg *config.Config, editor string) error {
	if editor == "" {
		editor = "code -n -w"
	}

	// TODO detect appropriate editor

	// TODO handle arguments with spaces, etc
	parts := strings.Fields(strings.TrimSpace(editor))
	cmd := parts[0]
	args := parts[1:]
	args = append(args, filepath.Dir(pkg.ConfigPath))

	c := exec.Command(cmd, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stdout
	c.Stdin = os.Stdin
	return c.Run()
}
