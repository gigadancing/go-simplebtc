package bc

// 交易
type Transaction struct {
	Hash []byte      // 交易哈希
	In   []*TxInput  // 交易输入
	Out  []*TxOutput // 交易输出
}

// coinbase交易

// 转账交易
