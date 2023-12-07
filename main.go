package main

import (
	"fmt"
	"os"

	"github.com/j13g/goutil/cli"

	"github.com/gimme-repos/gimme/cmd"
)

func main() {
	app := cli.NewCLI("gimme")
	cmd.SetupGimmeCLI(app)

	if err := app.Run(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
