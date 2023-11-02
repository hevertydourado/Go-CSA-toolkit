// cli/cmd/domain.go

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var domainCmd = &cobra.Command{
	Use:   "domain",
	Short: "Get information about a domain",
	Run: func(cmd *cobra.Command, args []string) {
		domain := args[0]
		// Implement domain information retrieval here
		fmt.Printf("Domain Information for %s\n", domain)
	},
}

func init() {
	rootCmd.AddCommand(domainCmd)
}
