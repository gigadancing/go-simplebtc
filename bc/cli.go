package bc

import (
	"flag"
	"fmt"
	"log"
	"os"
)

//
type CLI struct {
	BC *BlockChain
}

// 展示用法
func Usage() {
	fmt.Println("Usage:")
	fmt.Printf("\tcreateblockchain -- 创建区块链\n")
	fmt.Printf("\tinsertblock -data DATA -- 交易数据\n")
	fmt.Printf("\tprintblockchain -- 输出区块链信息")
}

// 校验参数
// 只输入程序名称，就输出指令用法然后退出
func Validate() {
	if len(os.Args) < 2 {
		Usage()    // 打印用法
		os.Exit(1) // 退出程序
	}
}

// 添加区块
func (cli *CLI) insertBlock(data string) {
	cli.BC.InsertBlock([]byte(data))
}

// 输出区块链信息
func (cli *CLI) printBlockChain() {
	cli.BC.PrintChain()
}

// 创建区块链
func (cli *CLI) createBlockChain() {
	NewBlockChain()
}

//
func (cli *CLI) Run() {
	Validate()
	insertBlockCmd := flag.NewFlagSet("insertblock", flag.ExitOnError)
	printBlockChainCmd := flag.NewFlagSet("printblockchain", flag.ExitOnError)
	createBlockChainCmd := flag.NewFlagSet("newblockchain", flag.ExitOnError)

	flagInsertBlockArg := insertBlockCmd.String("data", "send 100 BTC to everyone", "交易数据")
	switch os.Args[1] {
	case "insertblock":
		if err := insertBlockCmd.Parse(os.Args[2:]); err != nil {
			log.Panicf("parse cmd of insertblock failed:%v\n", err)
		}
	case "printblockchain:":
		if err := printBlockChainCmd.Parse(os.Args[2:]); err != nil {
			log.Panicf("parse cmd of printblockchain failed:%v\n", err)
		}
	case "createblockchain":
		if err := createBlockChainCmd.Parse(os.Args[2:]); err != nil {
			log.Panicf("parse cmd of createblockchain failed:%v\n", err)
		}
	default:
		Usage()
		os.Exit(1)
	}

	if insertBlockCmd.Parsed() {
		if *flagInsertBlockArg == "" {
			Usage()
			os.Exit(1)
		}
		cli.insertBlock(*flagInsertBlockArg)
	}

	if printBlockChainCmd.Parsed() {
		cli.printBlockChain()
	}

	if createBlockChainCmd.Parsed() {
		cli.createBlockChain()
	}

}
