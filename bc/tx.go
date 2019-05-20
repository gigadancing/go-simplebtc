package bc

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"log"
)

// 交易
type Transaction struct {
	Hash []byte   // 交易哈希
	Vin  []*TxIn  // 交易输入
	Vout []*TxOut // 交易输出
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
		Vin:  []*TxIn{in},
		Vout: []*TxOut{out},
	}
	coinbaseTx.TxHash()
	return coinbaseTx
}

// 转账交易
func NewSimpleTx(from, to string, amount int, chain *BlockChain) *Transaction {
	var (
		txIns  []*TxIn
		txOuts []*TxOut
	)

	// 查找指定地址的可用UTXO
	value, spendableUTXO := chain.FindSpendableUTXO(from, int64(amount))
	for txhash, indexArray := range spendableUTXO {
		hashBytes, _ := hex.DecodeString(txhash)
		for _, index := range indexArray {
			// 此处的输出是需要消费的，必然会被其他的交易输入所引用
			txin := &TxIn{
				Prevout:   OutPoint{hashBytes, index},
				ScriptSig: from,
			}
			txIns = append(txIns, txin)
		}
	}

	// 转账
	out := &TxOut{Value: int64(amount), ScriptPubkey: to}
	txOuts = append(txOuts, out)

	// 找零
	out = &TxOut{Value: value - int64(amount), ScriptPubkey: from}
	txOuts = append(txOuts, out)

	// 生成交易
	tx := &Transaction{Hash: nil, Vin: txIns, Vout: txOuts}
	tx.TxHash()

	return tx
}

// 判断交易是否是coinbase交易
func (tx *Transaction) IsCoinbaseTx() bool {
	return len(tx.Vin[0].Prevout.Hash) == 0 && tx.Vin[0].Prevout.Index == -1
}
