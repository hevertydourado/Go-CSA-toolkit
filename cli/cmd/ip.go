// cli/cmd/ip.go

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Get information about an IP address",
	Run: func(cmd *cobra.Command, args []string) {
		ip := args[0]
		// Implement IP information retrieval here
		fmt.Printf("IP Address Information for %s\n", ip)
	},
}

func init() {
	rootCmd.AddCommand(ipCmd)
}
