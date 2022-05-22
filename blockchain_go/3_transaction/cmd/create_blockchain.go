package cmd

import (
	"blockchain/core"
	"fmt"

	"github.com/spf13/cobra"
)

var createBlockchainCmd = &cobra.Command{
	Use:   "createblockchain",
	Short: "Create a blockchain and send genesis block reward to address",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		address := args[0]
		// if !ValidateAddress(address) {
		// 	log.Panic("ERROR: Address is not valid")
		// }
		bc := core.CreateBlockchain(address)
		defer bc.DBClose()

		// UTXOSet := UTXOSet{bc}
		// UTXOSet.Reindex()

		fmt.Println("Done!")
	},
}
