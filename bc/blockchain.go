package bc

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"simplebtc/util"
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

// 创建区块链
func NewBlockChain() *BlockChain {
	// 创建或打开数据库
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Panicf("open block.db failed:%v\n", err)
	}
	//defer db.Close()

	var latestBlockHash []byte

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockTableName))
		if bucket == nil { // 数据库不存在，创建数据库
			bucket, err = tx.CreateBucket([]byte(blockTableName))
			if err != nil {
				log.Panicf("create bucket [%s] failed:%v\n", blockTableName, err)
			}
		} else { // 数据库存在，更新数据
			genesisBlock := CreateGenesisBlock("today is saturday,2019/3/30")
			err = bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil {
				log.Panicf("put block data to db failed:%v\n", err)
			}
			err = bucket.Put([]byte("latest"), genesisBlock.Hash)
			if err != nil {
				log.Panicf("put the hash of latest block failed:%v\n", err)
			}
			latestBlockHash = genesisBlock.Hash
		}

		return nil
	})
	if err != nil {
		log.Panicf("update the data of genesis block failed:%v\n", err)
	}

	return &BlockChain{
		DB:  db,
		Tip: latestBlockHash,
	}
}

// 插入区块
func (blockchain *BlockChain) InsertBlock(data []byte) {
	// 更新数据
	err := blockchain.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockTableName)) // 获取表
		if bucket != nil {                          // 表存在
			var err error
			latestBlockHash := bucket.Get([]byte("latest"))
			latestBlockData := bucket.Get(latestBlockHash)
			latestBlock := Deserialize(latestBlockData)
			newBlock := NewBlock(latestBlock.Number+1, latestBlock.Hash, data)
			if err = bucket.Put(newBlock.Hash, newBlock.Serialize()); err != nil {
				log.Panicf("insert block into db failed:%v\n", err)
			}
			if err = bucket.Put([]byte("latest"), newBlock.Hash); err != nil {
				log.Panicf("update latest block failed:%v\n", err)
			}
			blockchain.Tip = newBlock.Hash

			fmt.Printf("blockNumber:%d parent:%v hash:%v tip:%v\n", newBlock.Number,
				util.HexToString(newBlock.ParentHash), util.HexToString(newBlock.Hash), util.HexToString(newBlock.Hash))

		}
		return nil
	})

	if err != nil {
		log.Panicf("blockchain insert block failed:%v\n", err)
	}
}
