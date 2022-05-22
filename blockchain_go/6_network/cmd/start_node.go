package cmd

import (
	"blockchain/core"
	"blockchain/server"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var startNodeMiner string

func init() {
	startNodeCmd.Flags().StringVarP(&startNodeMiner, "miner", "m", "", "Enable mining mode and send reward to ADDRESS")
}

var startNodeCmd = &cobra.Command{
	Use:   "startnode",
	Short: "startnode ADDRESS - Start a node with ID specified in nodeID",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Starting node %s\n", nodeID)
		if len(startNodeMiner) > 0 {
			if core.ValidateAddress(startNodeMiner) {
				fmt.Println("Mining is on. Address to receive rewards: ", startNodeMiner)
			} else {
				log.Panic("Wrong miner address!")
			}
		}
		server.StartServer(nodeID, startNodeMiner)
	},
}
