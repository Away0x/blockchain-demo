package cmd

import (
	"blockchain/core"
	"blockchain/tools"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var getBalanceCmd = &cobra.Command{
	Use:   "getbalance",
	Short: "Get balance of address",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		address := args[0]

		if !core.ValidateAddress(address) {
			log.Panic("ERROR: Address is not valid")
		}

		bc := core.NewBlockchain(nodeID)
		defer bc.DBClose()

		balance := 0
		pubKeyHash := tools.Base58Decode([]byte(address))
		pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]

		uTXOSet := core.UTXOSet{
			Blockchain: bc,
		}
		uTXOs := uTXOSet.FindUTXO(pubKeyHash)

		for _, out := range uTXOs {
			balance += out.Value
		}

		fmt.Printf("Balance of '%s': %d\n", address, balance)
	},
}
