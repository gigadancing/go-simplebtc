package bc

// 查询余额
func (cli *CLI) getBalance(from string) {
	// 获取指定地址的余额
	bc := BlockChainObject()
	defer bc.DB.Close()
	bc.GetBalance(from)
}
