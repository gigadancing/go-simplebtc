package bc

import (
	"github.com/boltdb/bolt"
	"log"
)

// 区块链迭代器
type BlockChainIterator struct {
	DB   *bolt.DB //
	Hash []byte   // 当前区块哈希
}

// 创建迭代器
func (bc *BlockChain) Iterator() *BlockChainIterator {
	return &BlockChainIterator{
		DB:   bc.DB,
		Hash: bc.Tip,
	}
}

// 返回迭代器对应的区块
func (bcit *BlockChainIterator) Block() *Block {
	var block *Block
	if err := bcit.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockTableName))
		if bucket != nil {
			if data := bucket.Get(bcit.Hash); data != nil {
				block = Deserialize(data)
			}
		}
		return nil
	}); err != nil {
		log.Panicf("BlockChain iterator current view failed: %v\n", err)
	}
	return block
}

// 迭代器后移
func (bcit *BlockChainIterator) Next() {
	if err := bcit.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockTableName))
		if bucket != nil {
			block := bcit.Block()
			bcit.Hash = block.PrevHash // 更新迭代器
		}
		return nil
	}); err != nil {
		log.Panicf("blockchain iterator next view failed: %v\n", err)
	}
}
