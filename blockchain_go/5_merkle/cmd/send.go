package cmd

import (
	"blockchain/core"
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "send FROM TO AMOUNT- Send AMOUNT of coins from FROM address to TO. Mine on the same node, when -mine is set.",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		from := args[0]
		to := args[1]
		amount, _ := strconv.Atoi(args[2])

		if !core.ValidateAddress(from) {
			log.Panic("ERROR: Sender address is not valid")
		}
		if !core.ValidateAddress(to) {
			log.Panic("ERROR: Recipient address is not valid")
		}

		bc := core.NewBlockchain()
		uTXOSet := core.UTXOSet{
			Blockchain: bc,
		}
		defer bc.DBClose()

		tx := core.NewUTXOTransaction(from, to, amount, &uTXOSet)
		cbTx := core.NewCoinbaseTX(from, "")
		txs := []*core.Transaction{cbTx, tx}

		newBlock := bc.MineBlock(txs)
		uTXOSet.Update(newBlock)
		fmt.Println("Success!")
	},
}
