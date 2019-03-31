package bc

import (
	"bytes"
	"crypto/sha256"
	"math/big"
	"simplebtc/util"
)

const targetBit = 20 // 目标难度值，代表哈希前targetBit位为0才能满足条件

// 工作量证明
type POW struct {
	Block  *Block   // 对指定的区块进行验证
	Target *big.Int //
}

// 创建新的POW对象
func NewPOW(block *Block) *POW {
	target := big.NewInt(1)
	target.Lsh(target, 256-targetBit) // 左移256-targetBit位
	target.Sub(target, big.NewInt(1)) // 减1
	return &POW{
		Block:  block,
		Target: target,
	}
}

// 进行工作量证明
func (pow *POW) Run() ([]byte, int64) {
	var (
		nonce   int64    // 随机数（碰撞次数）
		hashInt big.Int  // 生成的哈希值对应的整数
		hash    [32]byte // 生成的哈希值
	)
	for {
		data := pow.PrepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])
		if pow.Target.Cmp(&hashInt) == 1 { // 生成的哈希值小于目标值
			break
		}
		nonce++
	}

	return hash[:], nonce
}

// 准备数据，将区块相关属性拼接返回一个字节数组
func (pow *POW) PrepareData(nonce int64) []byte {
	data := bytes.Join([][]byte{
		pow.Block.ParentHash,
		pow.Block.Data,
		util.IntToHex(pow.Block.Timestamp),
		util.IntToHex(pow.Block.Number),
		util.IntToHex(nonce),
		util.IntToHex(targetBit),
	}, []byte{})

	return data
}
