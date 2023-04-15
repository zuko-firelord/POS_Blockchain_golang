package main

import (
	"fmt"
	"log"
	"github.com/zuko-firelord/POW_Blockchain_golang/wallet"


)

func init() {
	log.SetPrefix("Blockchain:  ")
}

func main() {
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PublicKeyStr())
	fmt.Println(w.BlockchainAddress())


	//transaction
	t := wallet.NewTransaction(w.PrivateKey(),w.PublicKey(),w.BlockchainAddress(),"B",1.0)
	fmt.Printf("signature %s \n", t.GenerateSig())
}