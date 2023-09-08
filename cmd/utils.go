package shiploader

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wearewebera/shiploader/pkg/shiploader"
)

var utilsCmd = &cobra.Command{
	Use:     "utils",
	Aliases: []string{"util"},
	Short:   "Utils is a command with multiple subcommands to perform various utility functions",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := shiploader.Utils(args[0])
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(utilsCmd)
}
