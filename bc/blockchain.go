package bc

// 区块链
type BlockChain struct {
	Blocks []*Block // 存储有序的区块
}

// 创建区块链
func NewBlockChain() *BlockChain {
	genesis := CreateGenesisBlock("today is saturday")
	return &BlockChain{
		Blocks: []*Block{genesis},
	}
}

// 插入区块
func (bc *BlockChain) InsertBlock(num int64, parentHash []byte, data []byte) {
	b := NewBlock(num, parentHash, data)
	bc.Blocks = append(bc.Blocks, b)
}
