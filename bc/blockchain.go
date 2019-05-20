package bc

import (
	"encoding/hex"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"os"
	"simplebtc/util"
	"strconv"
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
func NewBlockChain(address string) *BlockChain {
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
		coinbaseTx := NewCoinbaseTx(address)
		genesisBlock := CreateGenesisBlock([]*Transaction{coinbaseTx}) // 创建创世区块
		err = bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())  // 存入创世区块
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
	if err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockTableName))
		if bucket != nil {
			tip = bucket.Get([]byte("tip"))
		}
		return nil
	}); err != nil {
		log.Panicf("BlockChainObject: view db error: %v\n", err)
	}

	return &BlockChain{
		DB:  db,
		Tip: tip,
	}
}

// 挖矿
func (bc *BlockChain) MineNewBlock(from, to, amount []string) {
	var (
		txs   []*Transaction // 要打包的交易
		block *Block
	)
	value, _ := strconv.Atoi(amount[0])
	tx := NewSimpleTx(from[0], to[0], value)
	txs = append(txs, tx)

	if err := bc.DB.View(func(tx *bolt.Tx) error { // 获取当前最新区块
		bucket := tx.Bucket([]byte(blockTableName))
		if nil != nil {
			h := bucket.Get([]byte("tip"))
			data := bucket.Get(h)
			block = Deserialize(data)
		}
		return nil
	}); err != nil {
		log.Panicf("MinBlock view db error: %v\n", err)
	}

	newBlock := NewBlock(block.Number+1, block.Hash, txs) // 打包的新区快
	if err := bc.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockTableName))
		if bucket != nil {
			if err := bucket.Put(newBlock.Hash, newBlock.Serialize()); err != nil { // 将打包的区块持久化
				log.Panicf("MineBlock: put new block error: %v\n", err)
			}
			if err := bucket.Put([]byte("tip"), newBlock.Hash); err != nil { // 更新最新区块
				log.Panicf("MineBlock: put tip error: %v\n", err)
			}
			bc.Tip = newBlock.Hash
		}
		return nil
	}); err != nil {
		log.Panicf("MineBlock: update error: %v\n", err)
	}
}

// 返回指定地址的utxo
func (bc *BlockChain) UnspentUTXO(addr string) []*UTXO {
	var utxos []*UTXO         // 未花费的交易输出
	blockItr := bc.Iterator() // 区块迭代器

	// key：每个input所引用交易的哈希
	// value：output索引列表
	spentOutputs := make(map[string][]int) // 已花费的交易输出

	for {
		block := blockItr.Block()   // 返回迭代器对应的区块
		if len(block.Parent) == 0 { // 到达创世块
			break
		}
		for _, tx := range block.Txs { // 遍历每个区块的交易
			// 查找输入
			if !tx.IsCoinbaseTx() { // 普通交易
				for _, in := range tx.Vin {
					if in.UnlockWithAddress(addr) { // 验证地址
						key := util.HexToString(in.Prevout.Hash)
						spentOutputs[key] = append(spentOutputs[key], in.Prevout.Index)
					}
				}
			}
			// 查找输出
			for index, out := range tx.Vout {
				if out.UnlockScripPubkeyWithAddress(addr) { // 验证输出是否属于传入地址
					if len(spentOutputs) != 0 { // 已花费输出不为空
						for txHash, indexArray := range spentOutputs {
							for _, i := range indexArray {
								if txHash == util.HexToString(tx.Hash) && i == index {
									continue
								} else {
									utxo := &UTXO{Hash: tx.Hash, Index: index, Output: out}
									utxos = append(utxos, utxo)
								}
							}
						}
					} else { // 已花费输出为空
						utxo := &UTXO{Hash: tx.Hash, Index: index, Output: out}
						utxos = append(utxos, utxo)
					}
				}
			}
		}
		blockItr.Next() // 迭代器后移
	}

	return utxos
}

// 查询指定地址的余额
func (bc *BlockChain) GetBalance(addr string) int64 {
	utxos := bc.UnspentUTXO(addr)
	var amount int64
	for _, utxo := range utxos {
		amount += utxo.Output.Value
	}
	return amount
}

// 转账
// 通过查找可用的UTXO，超过需要的资金即可中断查找
func (bc *BlockChain) FindSpendableUTXO(from string, amount int64) (int64, map[string][]int) {
	var (
		value         int64
		spendableUTXO = make(map[string][]int)
	)
	utxos := bc.UnspentUTXO(from)
	for _, utxo := range utxos {
		value += utxo.Output.Value
		h := hex.EncodeToString(utxo.Hash)
		spendableUTXO[h] = append(spendableUTXO[h], utxo.Index)
		if value >= amount {
			break
		}
	}

	if value < amount {
		fmt.Printf("%s 余额不足\n", from)
		os.Exit(1)
	}

	return value, spendableUTXO
}
