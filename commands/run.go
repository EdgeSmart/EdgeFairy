package commands

import (
	"github.com/EdgeSmart/EdgeFairy/run"
	"github.com/spf13/cobra"
)

var (
	clusterKey   string
	clusterToken string
	server       string
)

var runCMD = &cobra.Command{
	Use:   "run",
	Short: "Run EdgeFairy deamon",
	Long:  `Run EdgeFairy deamon`,
	Run:   runProcess,
}

func init() {
	runCMD.Flags().StringVarP(&clusterKey, "cluster_key", "k", "", "Cluster key")
	runCMD.Flags().StringVarP(&clusterToken, "cluster_token", "t", "", "Cluster token")
	runCMD.Flags().StringVarP(&server, "server", "s", "tcp://192.168.1.175:1883", "Server config")
	addCommand(runCMD)
}

func runProcess(cmd *cobra.Command, args []string) {
	run.Run(clusterKey, clusterToken, server)
}
