package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func main() {
	bc := NewBlockchain()

	bc.AddTransaction("transaction1")
	bc.AddTransaction("transaction2")
	bc.CreateBlock()

	bc.AddTransaction("transaction3")
	bc.AddTransaction("transaction4")
	bc.CreateBlock()

	bc.Print()
}

type Block struct {
	previousHash [32]byte
	transaction  []string
	timestamp    int64
}

func NewBlock(previousHash [32]byte, transaction []string) *Block {
	block := &Block{}

	block.previousHash = previousHash
	block.transaction = transaction
	block.timestamp = time.Now().UnixNano()

	return block
}

func (b *Block) Hash() [32]byte {
	m, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}

	return sha256.Sum256(m)
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64    `json:"timestamp"`
		PreviousHash [32]byte `json:"previous_hash"`
		Transaction  []string `json:"transaction"`
	}{
		Timestamp:    b.timestamp,
		PreviousHash: b.previousHash,
		Transaction:  b.transaction,
	})
}

func (b *Block) Print() {
	fmt.Printf("Timestamp       %d\n", b.timestamp)
	fmt.Printf("PreviousHash    %x\n", b.previousHash)
	fmt.Printf("Transaction     %s\n", b.transaction)
}

////////////////////////////////////////////////////////////////////////////////

type Blockchain struct {
	transactionPool []string
	blocks          []*Block
}

func NewBlockchain() *Blockchain {
	bc := &Blockchain{}

	// Genesis Blockの作成 . 最初のブロックは直前のブロックのハッシュを持たないため0値を入れる
	bc.CreateBlock()

	return bc
}

func (bc *Blockchain) CreateBlock() {
	if len(bc.blocks) == 0 {
		b := NewBlock([32]byte{}, nil)
		bc.blocks = append(bc.blocks, b)
	} else {
		ph := bc.blocks[len(bc.blocks)-1].Hash()

		b := NewBlock(ph, bc.transactionPool)
		bc.blocks = append(bc.blocks, b)
	}

	bc.transactionPool = nil // プールは一時的なもののためトランザクションをクリア
}

func (bc *Blockchain) AddTransaction(transaction string) {
	bc.transactionPool = append(bc.transactionPool, transaction)
}

func (bc *Blockchain) Print() {
	for i, b := range bc.blocks {
		fmt.Printf("%s Block %d %s\n", strings.Repeat("=", 15), i, strings.Repeat("=", 15))
		b.Print()
	}
}
