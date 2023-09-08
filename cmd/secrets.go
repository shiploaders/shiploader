// path: cmd/shiploader/secrets.go

package shiploader

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wearewebera/shiploader/pkg/shiploader"
)

var secretsCmd = &cobra.Command{
	Use:     "secrets",
	Aliases: []string{"secrets"},
	Short:   "Secrets is a command with arguments to perform secret migration task",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := shiploader.Utils(args[0])
		fmt.Println(res)
	},
}

func init() {
	migrateCmd.AddCommand(secretsCmd)
}
