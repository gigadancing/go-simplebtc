package bc

// 创建区块链
func (cli *CLI) createBlockChain(address string) {
	bc := NewBlockChain(address)
	defer bc.DB.Close()
}
