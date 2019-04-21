package bc

import (
	"fmt"
	"os"
)

// 打印区块链信息
func (cli *CLI) printBlockChain() {
	if !dbExist() {
		fmt.Println("db not exist")
		os.Exit(1)
	}
	bc := BlockChainObject()
	defer bc.DB.Close()
	bc.PrintChain()
}
