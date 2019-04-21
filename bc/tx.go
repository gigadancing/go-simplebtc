package bc

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

// 交易
type Transaction struct {
	Hash []byte   // 交易哈希
	Ins  []*TxIn  // 交易输入
	Outs []*TxOut // 交易输出
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
	in := &TxIn{
		Prevout: OutPoint{
			Hash:  nil,
			Index: 0,
		},
		ScriptSig: "mining award",
	}
	out := &TxOut{
		Value:        10,
		ScriptPubkey: address,
	}

	coinbaseTx := &Transaction{
		Hash: nil,
		Ins:  []*TxIn{in},
		Outs: []*TxOut{out},
	}
	coinbaseTx.TxHash()
	return coinbaseTx
}

// 转账交易
func NewSimpleTx(from, to string, amount int) *Transaction {
	var (
		txsIn  []*TxIn
		txsOut []*TxOut
	)
	in := &TxIn{ // 消费
		Prevout: OutPoint{
			Hash:  nil,
			Index: 0,
		},
		ScriptSig: from,
	}
	txsIn = append(txsIn, in)
	out := &TxOut{ // 转账
		Value:        int64(amount),
		ScriptPubkey: to,
	}
	txsOut = append(txsOut, out)
	out = &TxOut{ // 找零
		Value:        10 - int64(amount),
		ScriptPubkey: from,
	}
	txsOut = append(txsOut, out)

	tx := &Transaction{
		Hash: nil,
		Ins:  txsIn,
		Outs: txsOut,
	}
	tx.TxHash()
	return tx
}

// 判断交易是否是coinbase交易
func (tx *Transaction) IsCoinbaseTx() bool {
	return len(tx.Ins[0].Prevout.Hash) == 0 && tx.Ins[0].Prevout.Index == -1
}
