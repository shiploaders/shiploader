package shiploader

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wearewebera/shiploader/pkg/shiploader"
)

var migrateCmd = &cobra.Command{
	Use:     "migrate",
	Aliases: []string{"mig"},
	Short:   "Migrate is a command with multiple subcommands to perform various migration functions",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := shiploader.Utils(args[0])
		fmt.Println(res)
	},
}

func init() {
	utilsCmd.AddCommand(migrateCmd)
}
