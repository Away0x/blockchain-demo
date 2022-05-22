package cmd

import (
	"blockchain/core"
	"fmt"

	"github.com/spf13/cobra"
)

var getBalanceCmd = &cobra.Command{
	Use:   "getbalance",
	Short: "Get balance of address",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		address := args[0]

		bc := core.NewBlockchain(address)
		defer bc.DBClose()

		balance := 0
		UTXOs := bc.FindUTXO(address)

		for _, out := range UTXOs {
			balance += out.Value
		}

		fmt.Printf("Balance of '%s': %d\n", address, balance)
	},
}
