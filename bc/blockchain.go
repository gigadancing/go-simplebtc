package bc

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"os"
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

//
func dbExist() bool {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		return false
	}
	return true
}

// 创建区块链
func NewBlockChain(txs []*Transaction) *BlockChain {
	var (
		db  *bolt.DB
		tip []byte
		err error
	)
	// 数据库存在，打开数据库，获得BlockChain实例
	if dbExist() {
		db, err = bolt.Open(dbName, 0600, nil)
		if err != nil {
			log.Panicf("open block.db failed: %v\n", err)
		}

		err = db.View(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte(blockTableName))
			tip = bucket.Get([]byte("tip"))
			return nil
		})
		if err != nil {
			log.Panicf("db view failed: %v\n", err)
		}
		return &BlockChain{
			DB:  db,
			Tip: tip,
		}
	}

	// 创建数据库
	db, err = bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Panicf("open block.db failed: %v\n", err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockTableName))
		bucket, err = tx.CreateBucket([]byte(blockTableName)) // 创建表
		if err != nil {
			log.Panicf("create bucket [%s] failed: %v\n", blockTableName, err)
		}
		genesisBlock := CreateGenesisBlock(txs)                       // 创建创世区块
		err = bucket.Put(genesisBlock.Hash, genesisBlock.Serialize()) // 存入创世区块
		if err != nil {
			log.Panicf("put genesis block data into db failed: %v\n", err)
		}
		err = bucket.Put([]byte("tip"), genesisBlock.Hash) // 存入最新区块哈希
		if err != nil {
			log.Panicf("put the hash of latest block failed: %v\n", err)
		}
		tip = genesisBlock.Hash
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
func (bc *BlockChain) InsertBlock(txs []*Transaction) {
	var err error
	// 更新数据
	err = bc.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockTableName)) // 获取表
		if bucket != nil {                          // 表存在
			if tipData := bucket.Get(bc.Tip); tipData != nil {
				tipBlock := Deserialize(tipData)
				newBlock := NewBlock(tipBlock.Number+1, tipBlock.Hash, txs)
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
		fmt.Printf("Height:%d,Timstamp:%d,Parent:%x,Hash:%x,Data:%v,Nonce:%d\n", curBlock.Number,
			curBlock.Timestamp, curBlock.Parent, curBlock.Hash, curBlock.Txs, curBlock.Nonce)
		hashInt := big.NewInt(0).SetBytes(curBlock.Parent)
		if big.NewInt(0).Cmp(hashInt) == 0 { // 到达创世块
			break
		}
		itr.Next()
	}
}

// 返回BlockChain对象
func BlockChainObject() *BlockChain {
	var (
		db  *bolt.DB
		tip []byte
		err error
	)
	if db, err = bolt.Open(dbName, 0600, nil); err != nil {
		log.Panicf("get the object of blockchain failed: %v", err)
	}
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockTableName))
		if bucket != nil {
			tip = bucket.Get([]byte("tip"))
		}
		return nil
	})

	return &BlockChain{
		DB:  db,
		Tip: tip,
	}
}
