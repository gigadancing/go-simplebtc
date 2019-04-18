package bc

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

// 交易
type Transaction struct {
	Hash []byte      // 交易哈希
	In   []*TxInput  // 交易输入
	Out  []*TxOutput // 交易输出
}

// 生成交易哈希
func (tx *Transaction) TxHash() {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	if err := encoder.Encode(tx); err != nil {
		log.Panicf("generate tx hash failed:%v\n", err)
	}
	h := sha256.Sum256(res.Bytes())
	tx.Hash = h[:]
}

// coinbase交易
func NewCoinbaseTx(address string) *Transaction {
	in := &TxInput{
		Hash:      []byte{},
		Index:     -1,
		ScriptSig: "mining award",
	}
	out := &TxOutput{
		Value:        10,
		ScriptPubkey: address,
	}

	coinbaseTx := &Transaction{
		Hash: nil,
		In:   []*TxInput{in},
		Out:  []*TxOutput{out},
	}
	coinbaseTx.TxHash()
	return coinbaseTx
}

// 转账交易
func NewSimpleTx(from, to string, amount int) *Transaction {
	var (
		txsIn  []*TxInput
		txsOut []*TxOutput
	)
	in := &TxInput{ // 消费
		Hash:      nil,
		Index:     0,
		ScriptSig: from,
	}
	txsIn = append(txsIn, in)
	out := &TxOutput{ // 转账
		Value:        int64(amount),
		ScriptPubkey: to,
	}
	txsOut = append(txsOut, out)
	out = &TxOutput{ // 找零
		Value:        10 - int64(amount),
		ScriptPubkey: from,
	}
	txsOut = append(txsOut, out)

	tx := &Transaction{
		Hash: nil,
		In:   txsIn,
		Out:  txsOut,
	}
	tx.TxHash()
	return tx
}
