package bc

import (
	"fmt"
	"github.com/boltdb/bolt"
	"simplebtc/util"
	"testing"
)

func TestBlockChain(t *testing.T) {
	bc := NewBlockChain()
	fmt.Println(bc.DB)
	fmt.Println("tip:", util.HexToString(bc.Tip))
	defer bc.DB.Close()
	err := bc.DB.View(func(tx *bolt.Tx) error {
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

	bc.InsertBlock([]byte("A transfer to B 100 BTC"))
	bc.InsertBlock([]byte("A transfer to D 100 BTC"))
	bc.InsertBlock([]byte("B transfer to C 13 BTC"))
	bc.InsertBlock([]byte("C transfer to E 3 BTC"))

}

func TestBlockChain_PrintChain(t *testing.T) {
	bc := NewBlockChain()
	defer bc.DB.Close()
	bc.InsertBlock([]byte("A transfer to B 100 BTC"))
	bc.InsertBlock([]byte("A transfer to D 100 BTC"))
	bc.InsertBlock([]byte("B transfer to C 13 BTC"))
	bc.PrintChain()
}
