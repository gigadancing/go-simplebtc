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
	BC *BlockChain
}

// 展示用法
func Usage() {
	fmt.Println("Usage:")
	fmt.Printf("\tcreateblockchain -address addr --地址\n")
	fmt.Printf("\tprintblockchain -- 输出区块链信息\n")
	fmt.Printf("\tsend -from FROM -to TO -amount AMOUNT -- 转账\n")
	fmt.Printf("\tgetbalance -address FROM -- 查询余额\n")
}

// 校验参数
// 只输入程序名称，就输出指令用法然后退出
func Validate() {
	if len(os.Args) < 2 {
		Usage()    // 打印用法
		os.Exit(1) // 退出程序
	}
}

//
func (cli *CLI) Run() {
	Validate()
	printBlockChainCmd := flag.NewFlagSet("printblockchain", flag.ExitOnError)
	createBlockChainCmd := flag.NewFlagSet("newblockchain", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	// 获得命令行参数
	flagCreateBlockChainArg := createBlockChainCmd.String("address", "", "地址")
	flagFromArg := sendCmd.String("from", "", "转账源地址")
	flagToArg := sendCmd.String("to", "", "转账目标地址")
	flagAmountArg := sendCmd.String("amount", "", "转账金额")
	flagGetBalanceArg := getBalanceCmd.String("from", "", "查询地址")
	switch os.Args[1] {
	case "send":
		if err := sendCmd.Parse(os.Args[2:]); err != nil {
			log.Panicf("parse cmd of send failed: %v\n", err)
		}
	case "printblockchain:":
		if err := printBlockChainCmd.Parse(os.Args[2:]); err != nil {
			log.Panicf("parse cmd of printblockchain failed: %v\n", err)
		}
	case "createblockchain":
		if err := createBlockChainCmd.Parse(os.Args[2:]); err != nil {
			log.Panicf("parse cmd of createblockchain failed: %v\n", err)
		}
	case "getbalance":
		if err := getBalanceCmd.Parse(os.Args[2:]); err != nil {
			log.Panicf("get balance error: %v\n", err)
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

	if getBalanceCmd.Parsed() {
		if *flagGetBalanceArg == "" {
			fmt.Println("未指定地址")
			Usage()
			os.Exit(1)
		}
		cli.getBalance(*flagGetBalanceArg)
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
