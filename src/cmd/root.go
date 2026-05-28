package cmd

import (
	"os"

	"github.com/AsTR0I/sipcV2/internal/infra/appinfo"
	"github.com/spf13/cobra"
)

var logLevel string

var rootCmd = &cobra.Command{
	Use:     appinfo.Service,
	Version: appinfo.Version,
	Short:   "SIPC/v2. SIP CLI client",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(
		&logLevel,
		"log-level",
		"info",
		"debug|info|warn|error",
	)
}
