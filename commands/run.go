package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var runCMD = &cobra.Command{
	Use:   "run",
	Short: "Run EdgeFairy deamon",
	Long:  `Run EdgeFairy deamon`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run")
	},
}

func init() {
	addCommand(runCMD)
}
