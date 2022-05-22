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

		bc := core.NewBlockchain(from)
		defer bc.DBClose()

		tx := core.NewUTXOTransaction(from, to, amount, bc)
		bc.MineBlock([]*core.Transaction{tx})
		fmt.Println("Success!")
	},
}
