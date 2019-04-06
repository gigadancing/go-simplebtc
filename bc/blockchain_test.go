package bc

import (
	"fmt"
	"github.com/boltdb/bolt"
	"simplebtc/util"
	"testing"
)

func TestBlockChain(t *testing.T) {
	blockchain := NewBlockChain()
	fmt.Println(blockchain.DB)
	fmt.Println("tip:", util.HexToString(blockchain.Tip))
	defer blockchain.DB.Close()
	err := blockchain.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockTableName))
		if bucket != nil {
			value := bucket.Get([]byte("tip"))
			fmt.Printf("value:%v\n", util.HexToString(value))
		} else {
			fmt.Println("the bucket is nil")
		}
		return nil
	})

	if err != nil {
		t.Fatalf("blockchain view failed:%v\n", err)
	}

	blockchain.InsertBlock([]byte("A transfer to B 100 BTC"))
	fmt.Println("tip:", util.HexToString(blockchain.Tip))
	blockchain.InsertBlock([]byte("A transfer to D 100 BTC"))
	fmt.Println("tip:", util.HexToString(blockchain.Tip))
	blockchain.InsertBlock([]byte("B transfer to C 13 BTC"))
	fmt.Println("tip:", util.HexToString(blockchain.Tip))
	blockchain.InsertBlock([]byte("C transfer to E 3 BTC"))
	fmt.Println("tip:", util.HexToString(blockchain.Tip))
}
