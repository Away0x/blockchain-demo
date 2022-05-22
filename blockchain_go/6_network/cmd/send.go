package cmd

import (
	"blockchain/core"
	"blockchain/server"
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var sendMine bool // 块会立刻被同一节点挖出来 (必须要有这个标志，因为初始状态时，网络中没有矿工节点)

func init() {
	sendCmd.Flags().BoolVarP(&sendMine, "mine", "m", false, "Mine immediately on the same node")
}

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "send FROM TO AMOUNT - Send AMOUNT of coins from FROM address to TO. Mine on the same node, when MINE is set.",
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
		if amount <= 0 {
			log.Panic("ERROR: Amount must be greater than zero")
		}

		bc := core.NewBlockchain(nodeID)
		uTXOSet := core.UTXOSet{
			Blockchain: bc,
		}
		defer bc.DBClose()

		wallets, err := core.NewWallets(nodeID)
		if err != nil {
			log.Panic(err)
		}
		wallet := wallets.GetWallet(from)

		tx := core.NewUTXOTransaction(&wallet, to, amount, &uTXOSet)
		if sendMine {
			cbTx := core.NewCoinbaseTX(from, "")
			txs := []*core.Transaction{cbTx, tx}

			newBlock := bc.MineBlock(txs)
			uTXOSet.Update(newBlock)
		} else {
			server.SendTx(server.KnownNodes[0], tx)
		}
		fmt.Println("Success!")
	},
}
