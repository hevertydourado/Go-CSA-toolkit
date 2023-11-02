package cmd

import (
	"fmt"

	"github.com/hevertydourado/Go-CSA-toolkit/cli"
	"github.com/spf13/cobra"
)

var cmdDir *cobra.Command

func runDir(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("requires one argument")
	}
	return cli.Dir(args[0])
}
