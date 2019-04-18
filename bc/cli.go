package bc

import (
	"flag"
	"fmt"
	"log"
	"os"
	"simplebtc/util"
)

//
type CLI struct {
}

// 展示用法
func Usage() {
	fmt.Println("Usage:")
	fmt.Printf("\tcreateblockchain -address addr --地址\n")
	fmt.Printf("\tinsertblock -data DATA -- 交易数据\n")
	fmt.Printf("\tprintblockchain -- 输出区块链信息\n")
	fmt.Printf("\tsend -from FROM -to TO -amount AMOUNT -- 转账\n")
}

// 校验参数
// 只输入程序名称，就输出指令用法然后退出
func Validate() {
	if len(os.Args) < 2 {
		Usage()    // 打印用法
		os.Exit(1) // 退出程序
	}
}

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

// 添加区块
func (cli *CLI) insertBlock(txs []*Transaction) {
	if !dbExist() {
		fmt.Println("db not exist")
		os.Exit(1)
	}
	bc := BlockChainObject()
	defer bc.DB.Close()
	bc.InsertBlock(txs)
}

// 输出区块链信息
func (cli *CLI) printBlockChain() {
	if !dbExist() {
		fmt.Println("db not exist")
		os.Exit(1)
	}
	bc := BlockChainObject()
	defer bc.DB.Close()
	bc.PrintChain()
}

// 创建区块链
func (cli *CLI) createBlockChain(address string) {
	bc := NewBlockChain(address)
	defer bc.DB.Close()
}

//
func (cli *CLI) Run() {
	Validate()
	insertBlockCmd := flag.NewFlagSet("insertblock", flag.ExitOnError)
	printBlockChainCmd := flag.NewFlagSet("printblockchain", flag.ExitOnError)
	createBlockChainCmd := flag.NewFlagSet("newblockchain", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	// 获得命令行参数
	flagInsertBlockArg := insertBlockCmd.String("data", "send 100 BTC to everyone", "交易数据")
	flagCreateBlockChainArg := createBlockChainCmd.String("address", "", "地址")
	flagFromArg := sendCmd.String("from", "", "转账源地址")
	flagToArg := sendCmd.String("to", "", "转账目标地址")
	flagAmountArg := sendCmd.String("amount", "", "转账金额")

	switch os.Args[1] {
	case "send":
		if err := sendCmd.Parse(os.Args[2:]); err != nil {
			log.Panicf("parse cmd of send failed: %v\n", err)
		}
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

	if sendCmd.Parsed() {
		if *flagFromArg == "" {
			fmt.Println("源地址不能为空")
			Usage()
			os.Exit(1)
		}
		if *flagToArg == "" {
			fmt.Println("目标地址不能为空")
			Usage()
			os.Exit(1)
		}
		if *flagAmountArg == "" {
			fmt.Println("金额不能为空")
			Usage()
			os.Exit(1)
		}
		cli.send(util.JsonToSlice(*flagFromArg), util.JsonToSlice(*flagToArg), util.JsonToSlice(*flagAmountArg))
	}

	if insertBlockCmd.Parsed() {
		if *flagInsertBlockArg == "" {
			Usage()
			os.Exit(1)
		}
		cli.insertBlock([]*Transaction{})
	}

	if printBlockChainCmd.Parsed() {
		cli.printBlockChain()
	}

	if createBlockChainCmd.Parsed() {
		if *flagCreateBlockChainArg == "" {
			Usage()
			os.Exit(1)
		}
		cli.createBlockChain(*flagCreateBlockChainArg)
	}

}
