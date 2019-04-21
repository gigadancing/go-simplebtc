package bc

// 输入
type TxInput struct {
	Hash      []byte // 引用的上一笔交易的哈希
	Vout      int    // 引用的上一笔交易的output索引号
	ScriptSig string // 脚本
}

// 判断能否引用指定地址的output
func (txin *TxInput) UnlockWithAddress(addr string) bool {
	return txin.ScriptSig == addr
}
