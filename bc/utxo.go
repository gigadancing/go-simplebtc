package bc

type UTXO struct {
	Hash   []byte
	Vout   int
	Output *TxOutput
}
