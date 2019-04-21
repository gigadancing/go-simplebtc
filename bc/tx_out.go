package bc

// 交易输出
type TxOut struct {
	Value        int64  // UTXO的金额
	ScriptPubkey string // 锁定脚本
}

//
func (out *TxOut) UnlockScripPubkeyWithAddress(addr string) bool {
	return addr == out.ScriptPubkey
}
