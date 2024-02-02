package cmd

import (
	"github.com/spf13/cobra"

	"github.com/jrdn/gimme/pkg/gimme"
)

func Some() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "some [flags] SPEC",
		Short:   "Install package specified by SPEC",
		Aliases: []string{"install"},
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			spec := args[0]
			return gimme.Install(cmd.Context(), spec)
		},
	}
	return cmd
}
