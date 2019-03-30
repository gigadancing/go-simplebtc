package bc

import (
	"fmt"
	"simplebtc/util"
	"testing"
)

func TestBlockChain(t *testing.T) {
	bc := NewBlockChain()
	bc.InsertBlock(int64(len(bc.Blocks)), bc.Blocks[len(bc.Blocks)-1].Hash, []byte("Alice send 100 to Bob"))
	bc.InsertBlock(int64(len(bc.Blocks)), bc.Blocks[len(bc.Blocks)-1].Hash, []byte("Alice send 1 to Kevin"))
	bc.InsertBlock(int64(len(bc.Blocks)), bc.Blocks[len(bc.Blocks)-1].Hash, []byte("Kevin send 2 to Tom"))
	bc.InsertBlock(int64(len(bc.Blocks)), bc.Blocks[len(bc.Blocks)-1].Hash, []byte("Alice send 3 to Bob"))
	bc.InsertBlock(int64(len(bc.Blocks)), bc.Blocks[len(bc.Blocks)-1].Hash, []byte("Alice send 4 to Bob"))
	fmt.Println("len:", len(bc.Blocks))
	for _, b := range bc.Blocks {
		fmt.Printf("%d : %v : %v\n", b.Number, util.HexToString(b.ParentHash), util.HexToString(b.Hash))
	}
}
