package shiploader

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "shiploader",
	Version: version,
	Short:   "A powerful CLI to ship applications to the cloud, migrate cloud services, bootstrap your cloud infrastructure, or maintain your existing cloud infrastructure",
	Long:    `Shiploader is a powerful CLI to ship applications to the cloud, migrate cloud services, bootstrap your cloud infrastructure, or maintain your existing cloud infrastructure.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
