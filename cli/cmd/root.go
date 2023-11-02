// cli/cmd/root.go

package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "domain-ip-info"}

func init() {
	// Aqui, você pode configurar opções globais do rootCmd, se necessário
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		// Trate erros de execução aqui
	}
}
