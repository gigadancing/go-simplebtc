package bc

import (
	"fmt"
	"simplebtc/util"
	"testing"
)

func TestBlock(t *testing.T) {
	genesis := NewBlock(0, nil, []byte("This is the genesis block"))
	fmt.Printf("%v %v %v \n", genesis.Number, util.HexToString(genesis.ParentHash), util.HexToString(genesis.Hash))
	block1 := NewBlock(1, genesis.Hash, []byte("This is the No.1 block"))
	fmt.Printf("%v %v %v \n", block1.Number, util.HexToString(block1.ParentHash), util.HexToString(block1.Hash))
	block2 := NewBlock(2, block1.Hash, []byte("This is the No.2 block"))
	fmt.Printf("%v %v %v \n", block2.Number, util.HexToString(block2.ParentHash), util.HexToString(block2.Hash))
}
