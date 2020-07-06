package main

import (
	"fmt"
	"log"
	"project_go/blockchain/block"
	"project_go/blockchain/wallet"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {

	//w := wallet.NewWallet()
	walletM := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()
	//fmt.Println(w.PrivateKey())
	//fmt.Println(w.PublicKey())
	//fmt.Println(w.PrivateKeyStr())
	//fmt.Println(w.PublicKeyStr())
	//fmt.Println(w.BlockchainAddress())

	t := wallet.NewTransaction(walletA.PrivateKey(), walletA.PublicKey(),
		walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0)
	//fmt.Printf("signature %s\n", t.GenerateSignature())

	blockchain := block.NewBlockchain(walletM.BlockchainAddress())
	isAdded := blockchain.AddTransaction(walletA.BlockchainAddress(), walletB.BlockchainAddress(),
		1.0, walletA.PublicKey(), t.GenerateSignature())
	fmt.Println("Added? ", isAdded)

	blockchain.Mining()
	blockchain.Print()

	fmt.Printf("A %.1f\n", blockchain.CalculateTotalAmount(walletA.BlockchainAddress()))
	fmt.Printf("B %.1f\n", blockchain.CalculateTotalAmount(walletB.BlockchainAddress()))
	fmt.Printf("M %.1f\n", blockchain.CalculateTotalAmount(walletM.BlockchainAddress()))





	//log.Println("test")
	//fmt.Println("test2")
	//b := NewBlock(0, "init hash")
	//b.Print()

	//myBlockchainAddress := "my_blockchain_address"
	//blockChain := NewBlockchain(myBlockchainAddress)
	//blockChain.Print()
	//
	//blockChain.AddTransaction("A", "B", 1.0)
	//blockChain.Mining()

	//previousHash := blockChain.LastBlock().Hash()
	//nonce := blockChain.ProofOfWork()
	//blockChain.CreateBlock(nonce, previousHash)
	//blockChain.Print()
	//
	//blockChain.AddTransaction("C", "D", 2.0)
	//blockChain.AddTransaction("X", "Y", 3.0)
	//previousHash = blockChain.LastBlock().Hash()
	//nonce = blockChain.ProofOfWork()
	//blockChain.CreateBlock(nonce, previousHash)
	//blockChain.Mining()
	//blockChain.Print()
	//blockChain.AddTransaction("C", "X", 10.0)
	//blockChain.Mining()
	//blockChain.Print()
	//
	//fmt.Printf("D %.1f\n", blockChain.CalculateTotalAmount("D"))
	//fmt.Printf("X %.1f\n", blockChain.CalculateTotalAmount("X"))

	//fmt.Println(blockchain)
	//block := &Block{nonce: 1}
	//fmt.Printf("%x\n", block.Hash())

}

