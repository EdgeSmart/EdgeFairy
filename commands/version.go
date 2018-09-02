package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	version = "0.0.0.1 beta"
)

var versionCMD = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Long:  `Print version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: V%s", version)
	},
}

func init() {
	addCommand(versionCMD)
}
