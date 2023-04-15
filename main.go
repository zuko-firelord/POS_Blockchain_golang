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
}