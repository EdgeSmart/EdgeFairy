package commands

import (
	"github.com/spf13/cobra"
)

var rootCMD = &cobra.Command{
	Use: "EdgeFairy",
}

// GetCommand return root command
func GetCommand() *cobra.Command {
	return rootCMD
}

func addCommand(command *cobra.Command) {
	rootCMD.AddCommand(command)
}
