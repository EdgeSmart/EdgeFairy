package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deployCMD = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy EdgeFairy",
	Long:  `Deploy EdgeFairy`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("OK")
	},
}

func init() {
	addCommand(deployCMD)
}
