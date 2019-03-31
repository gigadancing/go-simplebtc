package bc

import (
	"bytes"
	"crypto/sha256"
	"simplebtc/util"
	"time"
)

// 区块
type Block struct {
	Timestamp  int64  // 时间戳，区块产生的时间
	Number     int64  // 区块高度（索引、ID）
	Nonce      int64  //
	ParentHash []byte // 父区块哈希
	Hash       []byte // 当前区块哈希
	Data       []byte // 交易数据
}

// 创建区块
func NewBlock(num int64, parentHash []byte, data []byte) *Block {
	block := Block{
		Number:     num,
		ParentHash: parentHash,
		Data:       data,
		Timestamp:  time.Now().Unix(),
	}
	//b.SetHash()
	pow := NewPOW(&block)
	hash, nonce := pow.Run() // 进行工作量证明
	block.Hash = hash
	block.Nonce = nonce
	return &block
}

// 计算区块哈希
func (b *Block) SetHash() {
	h := util.IntToHex(b.Number) // 将整数转字节数组
	t := util.IntToHex(b.Timestamp)
	data := bytes.Join([][]byte{h, t, b.ParentHash, b.Data}, []byte{}) // 拼接所有字节数组
	hash := sha256.Sum256(data)                                        // 进行哈希
	b.Hash = hash[:]
}

// 生成创世区块
func CreateGenesisBlock(data string) *Block {
	return NewBlock(0, nil, []byte(data))
}
