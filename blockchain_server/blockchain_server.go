package main

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/zuko-firelord/POW_Blockchain_golang/block"
	"github.com/zuko-firelord/POW_Blockchain_golang/wallet"
)
var cache map[string]*block.Blockchain = make(map[string]*block.Blockchain)
type BlockchainServer struct {
	port uint16
}

func NewBlockchainServer(port uint16) *BlockchainServer {
	return &BlockchainServer{port}
}

func (bcs *BlockchainServer) Port() uint16 {
	return bcs.port
}

func (bcs *BlockchainServer) GetBlockchain() *block.Blockchain  {
	bc, ok := cache["blockchain"]
	if !ok {
		minerWallet := wallet.NewWallet()
		bc = block.NewBlockchain(minerWallet.BlockchainAddress(),bcs.Port())
		cache["blockchain"] = bc
		log.Printf("Private_key %v", minerWallet.PrivateKeyStr())
		log.Printf("PublicKeyStr %v", minerWallet.PublicKeyStr())
		log.Printf("BlockchainAddress %v", minerWallet.BlockchainAddress())
	}
	return bc
}

func (bcs *BlockchainServer) Getchain(w http.ResponseWriter, req *http.Request)  {
	switch req.Method{
	case http.MethodGet:
		w.Header().Add("Content-Type","application/json")
		bc := bcs.GetBlockchain()
		m,_:=bc.MarshalJSON()
		io.WriteString(w,string(m[:]))		 
	default:
		log.Print("Error: Invaild Http Method")
	}
}

func (bcs *BlockchainServer) Run() {
	http.HandleFunc("/", bcs.Getchain)
	log.Fatal(http.ListenAndServe("127.0.0.1:"+strconv.Itoa(int(bcs.Port())), nil))
}
