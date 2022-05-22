package cmd

import (
	"blockchain/core"
	"fmt"

	"github.com/spf13/cobra"
)

var createWalletCmd = &cobra.Command{
	Use:   "createwallet",
	Short: "Generates a new key-pair and saves it into the wallet file",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		wallets, _ := core.NewWallets()
		address := wallets.CreateWallet()
		wallets.SaveToFile()

		fmt.Printf("Your new address: %s\n", address)
	},
}
