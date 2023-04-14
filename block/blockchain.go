package block

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

const (
	mining_difficulty = 3
	mining_sender = "the bockchain"
	mining_reward = 1.0

)
type Block struct {
	timeStamp    int64
	nouce        int
	previousHash [32]byte
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

func (b *Block) Print() {
	fmt.Printf("timestamp              	%d \n", b.timeStamp)
	fmt.Printf("nouce              		%d \n", b.nouce)
	fmt.Printf("previousHash            %x \n", b.previousHash)
	for _, t := range b.transactions {
		t.Print()
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))

}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json:"timestamp"`
		Nouce        int            `json:"nouce"`
		Previoushash [32]byte       `json:"previous_hash"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Timestamp:    b.timeStamp,
		Nouce:        b.nouce,
		Previoushash: b.previousHash,
		Transactions: b.transactions,
	})
}


type Blockchain struct {
	transactonPool []*Transaction
	chain          []*Block
	blockchainAddress string
}

func NewBlockchain(blockchainAddress string) *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.blockchainAddress = blockchainAddress
	bc.CreateBlock(0, b.Hash())
	return bc
}
func (bc *Blockchain) CreateBlock(nouce int, previousHash [32]byte) *Block {
	b := NewBlock(nouce, previousHash, bc.transactonPool)
	bc.chain = append(bc.chain, b)
	bc.transactonPool = []*Transaction{}
	return b
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func (bc *Blockchain) Lastbloc() *Block {
	return bc.chain[len(bc.chain)-1]

}

func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.transactonPool = append(bc.transactonPool, t)
}

func (bc *Blockchain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0)
	for _, t := range bc.transactonPool {
		transactions = append(transactions,
			NewTransaction(t.senderBlockchainAddress,
				t.recipientBlockchainAddress,
				t.value))
	}
	return transactions
}

func (bc *Blockchain) validProof(nouce int, previousHash [32]byte, transactions []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := Block{0, nouce, previousHash, transactions}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	//fmt.Println(guessHashStr)
	return guessHashStr[:difficulty] == zeros
	}   

func (bc *Blockchain) ProofOfWork() int  {
	transactions := bc.CopyTransactionPool() 
	previousHash := bc.Lastbloc().Hash()

	nouce := 0
	for !bc.validProof(nouce, previousHash,transactions,mining_difficulty){
		nouce +=1
	}
	return nouce
}

func (bc *Blockchain) mining() bool  {
	bc.AddTransaction(mining_sender,bc.blockchainAddress,mining_reward)
	nonce := bc.ProofOfWork()
	previousHash := bc.Lastbloc().Hash()
	bc.CreateBlock(nonce,previousHash)
	log.Println("action=mining, status=success")
	return true
}

func (bc *Blockchain) calTotalAmt(blockchainAddress string)  float32 {
	var totalAmount float32 = 0.0
	for _,b := range bc.chain{
		for _,t := range b.transactions{
			value := t.value
			if blockchainAddress == t.recipientBlockchainAddress{
				totalAmount += value
			}
			if blockchainAddress == t.senderBlockchainAddress{
				totalAmount -= value
			}
		}
	}
	return totalAmount
}

type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("_", 40))
	fmt.Printf("sender_blockchain_address     %s\n", t.senderBlockchainAddress)
	fmt.Printf("recipient_blockchain_address  %s\n", t.recipientBlockchainAddress)
	fmt.Printf("value                         %1f\n", t.value)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderBlockchainAddress,
		Recipient: t.recipientBlockchainAddress,
		Value:     t.value,
	})
}


