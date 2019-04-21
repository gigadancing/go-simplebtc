package bc

import (
	"fmt"
	"os"
)

// 发送交易
func (cli *CLI) send(from, to, amount []string) {
	if !dbExist() {
		fmt.Println("db not exist.")
		os.Exit(1)
	}
	bc := BlockChainObject()
	defer bc.DB.Close()
	bc.MineNewBlock(from, to, amount)
}
