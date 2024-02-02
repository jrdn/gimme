package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/jrdn/gimme/pkg/gimme"
)

func Dir() *cobra.Command {
	cmd := &cobra.Command{
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			installed, err := gimme.ListInstalled()
			if err != nil {
				return err
			}
			for _, i := range installed {
				if strings.Contains(i.Spec, name) {
					fmt.Println(filepath.Dir(i.ConfigPath))
					return nil
				}
			}

			return nil
		},
	}
	return cmd
}
