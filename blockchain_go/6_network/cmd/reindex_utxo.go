package cmd

import (
	"blockchain/core"
	"fmt"

	"github.com/spf13/cobra"
)

var reindexUTXOCmd = &cobra.Command{
	Use:   "listaddresses",
	Short: "Lists all addresses from the wallet file",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		bc := core.NewBlockchain(nodeID)
		uTXOSet := core.UTXOSet{
			Blockchain: bc,
		}
		uTXOSet.Reindex()

		count := uTXOSet.CountTransactions()
		fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
	},
}
