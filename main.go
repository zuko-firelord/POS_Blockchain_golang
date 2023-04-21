package main

import (
	"fmt"
	"log"
	"github.com/zuko-firelord/POW_Blockchain_golang/block"
	"github.com/zuko-firelord/POW_Blockchain_golang/wallet"
)

func init() {
	log.SetPrefix("Blockchain:  ")
}

func main() {
	wM := wallet.NewWallet()
	wA := wallet.NewWallet()
	wB := wallet.NewWallet()
	//transaction
	t := wallet.NewTransaction(wA.PrivateKey(),wA.PublicKey(),wA.BlockchainAddress(),wB.BlockchainAddress(),1.0)

	blockchain := block.NewBlockchain(wM.BlockchainAddress())
	isadded := blockchain.AddTransaction(wA.BlockchainAddress(),wB.BlockchainAddress(),1.0,wA.PublicKey(),t.GenerateSig())
	fmt.Println("Added?", isadded)

	blockchain.Mining()
	blockchain.Print()

	fmt.Printf("A %.1f\n", blockchain.CalTotalAmt(wA.BlockchainAddress()))
	fmt.Printf("B %.1f\n", blockchain.CalTotalAmt(wB.BlockchainAddress()))
	fmt.Printf("M %.1f\n", blockchain.CalTotalAmt(wM.BlockchainAddress()))

}