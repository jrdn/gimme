package cmd

import (
	"os"
	"path/filepath"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"

	"github.com/jrdn/goutil/cli/outputter"

	"github.com/jrdn/gimme/pkg/gimme"
)

func List() *cobra.Command {
	cmd := &cobra.Command{
		Aliases: []string{"ls"},
		RunE: func(cmd *cobra.Command, args []string) error {
			packages, err := gimme.ListInstalled()
			if err != nil {
				return err
			}

			jsonOutput, err := cmd.Flags().GetBool("json")
			if err != nil {
				return err
			}

			if jsonOutput {
				outputter.JSONOutputter{}.Print(packages)
			} else {
				t := table.NewWriter()
				t.SetOutputMirror(os.Stdout)
				t.AppendHeader(table.Row{"Name", "Path"})

				for _, p := range packages {
					t.AppendRow(table.Row{p.Spec, filepath.Dir(p.ConfigPath)})
				}
				t.SetStyle(table.StyleColoredDark)
				t.Render()
			}

			return nil
		},
	}

	cmd.Flags().BoolP("json", "j", false, "JSON output")
	return cmd
}
