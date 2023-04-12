package main

import (
	"fmt"
	"log"
	"strings"
	"time"
	"encoding/json"
	"crypto/sha256"
)

type Block struct {
	nouce        int
	previousHash [32]byte
	timeStamp    int64
	transactions []*Transaction
}

func NewBlock(nouce int, previousHash [32]byte, transactions []*Transaction) *Block {
	b := new(Block)
	b.timeStamp = time.Now().UnixNano()
	b.nouce = nouce
	b.previousHash = previousHash
	b.transactions = transactions
	return b
}

func (b *Block) Print()  {
	fmt.Printf("timestamp              	%d \n",b.timeStamp)
	fmt.Printf("nouce              		%d \n",b.nouce)
	fmt.Printf("previousHash            %x \n",b.previousHash)
	for _ ,t := range b.transactions{
		t.Print()
	}
}

func (b *Block) Hash() [32]byte {
	m,_ := json.Marshal(b)
	return sha256.Sum256([]byte(m))

}

func (b *Block) MarshalJSON() ([]byte,error)  {
	return json.Marshal( struct{
		Timestamp int64 `json:"timestamp"`
		Nouce int		`json:"nouce"`
		Previoushash [32]byte	`json:"previous_hash"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Timestamp: b.timeStamp,
		Nouce: b.nouce,
		Previoushash: b.previousHash,
		Transactions: b.transactions,
	})
}
func init()  {
	log.SetPrefix("Blockchain:  ")
}

type Blockchain struct {
	transactonPool []*Transaction 
	chain          []*Block
}

func NewBlockchain() *Blockchain  {
	b:= &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0,b.Hash())
	return bc
}
func (bc *Blockchain) CreateBlock(nouce int, previousHash [32]byte) *Block  {
	b:= NewBlock(nouce,previousHash,bc.transactonPool)
	bc.chain = append(bc.chain, b)
	bc.transactonPool = []*Transaction{}
	return b
}

func (bc *Blockchain) Print()  {
	for i, block := range bc.chain{
		fmt.Printf("%s Chain %d %s\n",strings.Repeat("=",25),i,strings.Repeat("=",25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func (bc *Blockchain) Lastbloc() *Block {
	return bc.chain[len(bc.chain)-1]

}

func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32)  {
	t := NewTransaction(sender,recipient,value)
	bc.transactonPool = append(bc.transactonPool, t)
}

type Transaction struct{
	senderBlockchainAddress string
	recipientBlockchainAddress string
	value float32
}

func NewTransaction(sender string, recipient string, value float32) *Transaction  {
	return &Transaction{sender,recipient,value}
}

func (t *Transaction) Print()  {
	fmt.Printf("%s\n",strings.Repeat("_",40))
	fmt.Printf("sender_blockchain_address     %s\n", t.senderBlockchainAddress)
	fmt.Printf("recipient_blockchain_address  %s\n", t.recipientBlockchainAddress)
	fmt.Printf("value                         %1f\n", t.value)
}

func (t * Transaction) MarshalJSON() ([]byte, error)  {
	return json.Marshal(struct{
		Sender string `json:"sender_blockchain_address"`
		Recipient string `json:"recipient_blockchain_address"`
		Value float32 `json:"value"`
	}{
		Sender: t.senderBlockchainAddress,
		Recipient: t.recipientBlockchainAddress,
		Value: t.value,
	})
}








func main()  {
	blockchain := NewBlockchain()
	blockchain.Print()

	blockchain.AddTransaction("a","b",1.0)
	previousHash := blockchain.Lastbloc().Hash()
	blockchain.CreateBlock(1,previousHash)
	blockchain.Print()

	blockchain.AddTransaction("b","c",2.0)
	blockchain.AddTransaction("c","d",2.0)
	previousHash = blockchain.Lastbloc().Hash()
	blockchain.CreateBlock(2, previousHash)
	blockchain.Print()
}