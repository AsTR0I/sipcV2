package cmd

import (
	"os"

	"github.com/AsTR0I/sipcV2/internal/infra/appinfo"
	"github.com/spf13/cobra"
)

var (
	logLevel string
)

var rootCmd = &cobra.Command{
	Use:     appinfo.Service,
	Version: appinfo.Version,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
		// DisableNoDescFlag: true,
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().SortFlags = false
	rootCmd.PersistentFlags().SortFlags = false
	rootCmd.PersistentFlags().StringVar(
		&logLevel,
		"log-level",
		"info",
		"debug|info|warn|error",
	)
}
