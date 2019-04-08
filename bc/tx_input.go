package bc

// 输入
type TxInput struct {
	Hash      []byte // 引用的上一笔交易的哈希
	Index     int    // 引用的上一笔交易的output索引号
	ScriptSig string // 脚本
}
