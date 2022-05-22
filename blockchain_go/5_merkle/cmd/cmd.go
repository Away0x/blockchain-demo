package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "blockchain",
}

func init() {
	rootCmd.AddCommand(
		printChainCmd,
		createBlockchainCmd,
		getBalanceCmd,
		sendCmd,
		listAddressesCmd,
		createWalletCmd,
		reindexUTXOCmd,
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
