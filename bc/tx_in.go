package bc

// 交易输入所指向的UTXO
type OutPoint struct {
	Hash  []byte // UTXO所在交易的哈希
	Index int    // UTXO的索引
}

// 交易输入
type TxIn struct {
	Prevout   OutPoint // 该输入引用的UTXO
	ScriptSig string   // 解锁脚本，用于解锁输入指向的UTXO
}

// 判断能否引用指定地址的output
func (txin *TxIn) UnlockWithAddress(addr string) bool {
	return txin.ScriptSig == addr
}
