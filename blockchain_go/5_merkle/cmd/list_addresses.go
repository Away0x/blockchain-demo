package cmd

import (
	"blockchain/core"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var listAddressesCmd = &cobra.Command{
	Use:   "listaddresses",
	Short: "Lists all addresses from the wallet file",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		wallets, err := core.NewWallets()
		if err != nil {
			log.Panic(err)
		}
		addresses := wallets.GetAddresses()

		for _, address := range addresses {
			fmt.Println(address)
		}
	},
}
