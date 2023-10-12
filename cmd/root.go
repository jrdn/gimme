package cmd

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"github.com/j13g/goutil/log"
)

func Root() *cobra.Command {
	cmd := &cobra.Command{
		TraverseChildren: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			levelString, err := cmd.Flags().GetString("level")
			if err != nil {
				return err
			}

			level, err := zerolog.ParseLevel(levelString)
			log.SetupLogging(
				log.WithLevel(level),
				log.WithStderr(),
				log.WithConsoleOutput(),
				log.WithAppName("gimme"),
			)

			l := log.Get()
			l.Info().Msg("Starting gimme")
			return nil
		},
	}
	cmd.PersistentFlags().StringP("level", "l", "warn", "log level")

	return cmd
}
