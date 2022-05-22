package cmd

import (
	"blockchain/core"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var printChainCmd = &cobra.Command{
	Use:   "printchain",
	Short: "Print all the blocks of the blockchain",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		bc := core.NewBlockchain("")
		defer bc.DBClose()

		bci := bc.Iterator()

		for {
			block := bci.Next()

			fmt.Printf("============ Block %x ============\n", block.Hash)
			// fmt.Printf("Height: %d\n", block.Height)
			fmt.Printf("Prev. block: %x\n", block.PrevBlockHash)
			pow := core.NewProofOfWork(block)
			fmt.Printf("PoW: %s\n\n", strconv.FormatBool(pow.Validate()))
			for _, tx := range block.Transactions {
				fmt.Println(tx)
			}
			fmt.Printf("\n\n")

			if len(block.PrevBlockHash) == 0 {
				break
			}
		}
	},
}
