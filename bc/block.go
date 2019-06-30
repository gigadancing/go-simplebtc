package bc

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"time"
)

// 区块
type Block struct {
	Timestamp int64  // 时间戳，区块产生的时间
	Number    int64  // 区块高度（索引、ID）
	Nonce     int64  //
	PrevHash  []byte // 父区块哈希
	Hash      []byte // 当前区块哈希
	Txs       []*Transaction
}

// 创建区块
func NewBlock(num int64, parentHash []byte, txs []*Transaction) *Block {
	block := Block{
		Number:    num,
		PrevHash:  parentHash,
		Txs:       txs,
		Timestamp: time.Now().Unix(),
	}
	pow := NewPOW(&block)
	hash, nonce := pow.Run() // 进行工作量证明
	block.Hash = hash
	block.Nonce = nonce
	return &block
}

// 生成创世区块
func CreateGenesisBlock(txs []*Transaction) *Block {
	return NewBlock(0, nil, txs)
}

// 序列化，把区块结构序列化为字节数组([]byte)
func (b *Block) Serialize() []byte {
	var data bytes.Buffer
	encoder := gob.NewEncoder(&data)          // 创建encoder对象
	if err := encoder.Encode(b); err != nil { // 编码
		log.Panicf("serialize block failed:%v\n", err)
	}
	return data.Bytes()
}

// 反序列化，把字节数组结构化为区块
func Deserialize(data []byte) *Block {
	b := Block{}
	decoder := gob.NewDecoder(bytes.NewReader(data))
	if err := decoder.Decode(&b); err != nil {
		log.Panicf("deserialize block failed: %v\n", err)
	}
	return &b
}

// 区块中所有交易的哈希
func (b *Block) TxsHash() []byte {
	var txHash [][]byte
	for _, tx := range b.Txs {
		txHash = append(txHash, tx.Hash)
	}
	h := sha256.Sum256(bytes.Join(txHash, []byte{}))
	return h[:]
}
