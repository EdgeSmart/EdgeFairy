package commands

import (
	"github.com/EdgeSmart/EdgeFairy/deploy"
	"github.com/spf13/cobra"
)

var deployCMD = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy EdgeFairy",
	Long:  `Deploy EdgeFairy`,
	Run: func(cmd *cobra.Command, args []string) {
		deploy.Run()
	},
}

func init() {
	addCommand(deployCMD)
}
