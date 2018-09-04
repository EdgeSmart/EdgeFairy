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
		fmt.Println("run command")
	},
}

func init() {
	addCommand(runCMD)
}
