package bc

// 未花费交易输出
type UTXO struct {
	Hash   []byte
	Index  int
	Output *TxOut
}
