package bc

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
)

const (
	dbName         = "block.db"
	blockTableName = "blocks"
)

// 区块链
type BlockChain struct {
	DB  *bolt.DB // 数据库
	Tip []byte   //最新区块哈希
}

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
	var (
		block *Block
		err   error
	)

	err = bcit.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockTableName))
		if bucket != nil {
			if data := bucket.Get(bcit.Hash); data != nil {
				block = Deserialize(data)
			}
		}
		return nil
	})
	if err != nil {
		log.Panicf("BlockChain iterator current view failed: %v\n", err)
	}
	return block
}

// 迭代器后移
func (bcit *BlockChainIterator) Next() {
	var (
		block *Block
		err   error
	)
	err = bcit.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockTableName))
		if bucket != nil {
			block = bcit.Block()
			bcit.Hash = block.Parent // 更新迭代器
		}
		return nil
	})

	if err != nil {
		log.Panicf("blockchain iterator next view failed: %v\n", err)
	}

}

// 创建区块链
func NewBlockChain() *BlockChain {
	// 创建或打开数据库，如果dbName不存在则创建，否则打开dbName
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Panicf("open block.db failed:%v\n", err)
	}

	var tip []byte // 最新区块哈希

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockTableName))
		if bucket == nil { // 表不存在
			bucket, err = tx.CreateBucket([]byte(blockTableName)) // 创建表
			if err != nil {
				log.Panicf("create bucket [%s] failed: %v\n", blockTableName, err)
			}
			genesisBlock := CreateGenesisBlock("today is saturday, 2019/3/30.") // 创建创世区块
			err = bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())       // 存入创世区块
			if err != nil {
				log.Panicf("put genesis block data into db failed: %v\n", err)
			}
			err = bucket.Put([]byte("tip"), genesisBlock.Hash) // 存入最新区块哈希
			if err != nil {
				log.Panicf("put the hash of latest block failed: %v\n", err)
			}
			tip = genesisBlock.Hash
		} else { // 表存在
			tip = bucket.Get([]byte("tip")) // 更新最先区块
		}
		return nil
	})

	if err != nil {
		log.Panicf("update the data of genesis block failed:%v\n", err)
	}

	return &BlockChain{
		DB:  db,
		Tip: tip,
	}
}

// 插入区块
func (bc *BlockChain) InsertBlock(data []byte) {
	var err error
	// 更新数据
	err = bc.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockTableName)) // 获取表
		if bucket != nil {                          // 表存在
			if tipData := bucket.Get(bc.Tip); tipData != nil {
				tipBlock := Deserialize(tipData)
				newBlock := NewBlock(tipBlock.Number+1, tipBlock.Hash, data)
				if err = bucket.Put(newBlock.Hash, newBlock.Serialize()); err != nil {
					log.Panicf("insert block into db failed: %v\n", err)
				}
				if err = bucket.Put([]byte("tip"), newBlock.Hash); err != nil {
					log.Panicf("update latest block failed: %v\n", err)
				}
				bc.Tip = newBlock.Hash
			}
		}
		return nil
	})

	if err != nil {
		log.Panicf("blockchain insert block failed:%v\n", err)
	}
}

// 遍历输出所有区块信息
func (bc *BlockChain) PrintChain() {
	var (
		curBlock *Block
		itr      = bc.Iterator()
	)
	fmt.Println("==========BLOCKCHAIN INFO==========")
	for {
		curBlock = itr.Block()
		fmt.Printf("Height:%d,Timstamp:%d,Parent:%x,Hash:%x,Data:%s,Nonce:%d\n", curBlock.Number,
			curBlock.Timestamp, curBlock.Parent, curBlock.Hash, string(curBlock.Data), curBlock.Nonce)
		hashInt := big.NewInt(0).SetBytes(curBlock.Parent)
		if big.NewInt(0).Cmp(hashInt) == 0 { // 到达创世块
			break
		}
		itr.Next()
	}
}
