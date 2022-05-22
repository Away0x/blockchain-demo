package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var nodeID string

var rootCmd = &cobra.Command{
	Use: "blockchain",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&nodeID, "node", "n", "3000", "set node id")

	rootCmd.AddCommand(
		printChainCmd,
		createBlockchainCmd,
		getBalanceCmd,
		sendCmd,
		listAddressesCmd,
		createWalletCmd,
		reindexUTXOCmd,
		startNodeCmd,
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
