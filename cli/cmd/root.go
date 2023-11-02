// cli/cmd/root.go

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "GoDomainInfo"}
var verbose bool
var outputFormat string

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose mode")
	rootCmd.PersistentFlags().StringVar(&outputFormat, "output", "text", "Output format (text, json, etc.)")
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		// Implement outras lógicas de tratamento de erros aqui, como registrar em logs
		os.Exit(1) // Opcional: você pode sair com um código de saída não zero
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		handleError(err)
	}
}
