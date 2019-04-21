package bc

// 输出
type TxOutput struct {
	Value        int64  // 金额
	ScriptPubkey string // 脚本
}

//
func (out *TxOutput) UnlockScripPubkeyWithAddress(addr string) bool {
	return addr == out.ScriptPubkey
}
