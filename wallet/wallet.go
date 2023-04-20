package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/zuko-firelord/POW_Blockchain_golang/utils"
	"golang.org/x/crypto/ripemd160"
)

type wallet struct{
	privateKey *ecdsa.PrivateKey
	publicKey *ecdsa.PublicKey
	blockchainAddress string
}

func NewWallet () *wallet  {
	//1. Creating ECDSA private key (32 bytes) public key (64bytes)
	w := new(wallet)
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(),rand.Reader)
	w.privateKey = privateKey
	w.publicKey = &w.privateKey.PublicKey

	//2. Perform SHA-256 hashing on the public key (32 bytes)
	h2 := sha256.New()
	h2.Write(w.publicKey.X.Bytes())
	h2.Write(w.publicKey.Y.Bytes())
	digest2 := h2.Sum(nil)

	//3. Perform RIPEMD-160 hashing on the result of the SHA(20bytes)
	h3 := ripemd160.New()
	h3.Write(digest2)
	digest3 := h3.Sum(nil)

	//4. Add version byte in front of RIPEMD-160 hash (0x00 for main network)
	vd4:= make([]byte,21)
	vd4[0] = 0x00
	copy(vd4[1:],digest3[:])
	
	//5. Perform SHA-256 hash on the extended RIPEMD-160 result
	h5 := sha256.New()
	h5.Write(vd4)
	digest5 := h5.Sum(nil)

	//6. Perform SHA-256 hash on the result of the pervious SHA-256 hash
	h6 := sha256.New()
	h6.Write(digest5)
	digest6 := h6.Sum(nil)
	
	//7. Take the first 4 bytes of the second SHA-256 hash for checksum
	chsum := digest6[:4]
	
	//8. Add the 4 checksum bytes from 7 at the end of extended RIPEMD-160 hash from 4 (25 bytes)
	dc8:= make([]byte,25)
	copy(dc8[:21],vd4[:])
	copy(dc8[21:], chsum[:])
		
	//9. Convert the result from a byte string into base58
	address := base58.Encode(dc8)
	w.blockchainAddress = address
	
	return w
}

func (w *wallet)  PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey	
}

func (w *wallet) PrivateKeyStr() string  {
	return fmt.Sprintf("%x",w.privateKey.D.Bytes())
}

func (w *wallet)  PublicKey() *ecdsa.PublicKey {
	return w.publicKey
}

func (w *wallet) PublicKeyStr() string  {
	return fmt.Sprintf("%x%x",w.publicKey.X.Bytes(), w.publicKey.Y.Bytes())
}

func (w *wallet)  BlockchainAddress() string {
	return w.blockchainAddress
}

//------------TRANSACTION----------------
type Transaction struct {
	senderPrivateKey 	*ecdsa.PrivateKey
	senderPublicKey 	*ecdsa.PublicKey
	senderBlockchainAddress string
	recipientBlockchainAddress string
	value 				float32
}

func NewTransaction(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, 
	sender string,recipient string, value float32) *Transaction  {
	return &Transaction{privateKey,publicKey,sender,recipient,value}
}

func (t *Transaction) GenerateSig() *utils.Signature  {
	m,_ := json.Marshal(t)
	h := sha256.Sum256([]byte(m))
	r,s,_ := ecdsa.Sign(rand.Reader,t.senderPrivateKey,h[:])
	return &utils.Signature {R: r,S: s}
}

func (t *Transaction) MarshalJSON() ([]byte,error)  {
	return json.Marshal(struct {
		Sender string `json:"sender_blockchain_address"`
		Recipient string `json:"recipient_blockchain_address"`
		Value float32 `json:"value"`
	}{
		Sender: t.senderBlockchainAddress,
		Recipient: t.recipientBlockchainAddress,
		Value: t.value,
	})
}
