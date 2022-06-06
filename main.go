package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"net/http"
)

type Block struct {
	Hash []byte
	Data []byte
	PrevHash []byte
}

type BlockChain struct {
	blocks []*Block
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{ []byte {}, []byte(data), prevHash }

	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks) -1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Deploy Blockchain", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[] *Block{Genesis()}}
}

func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	chain := InitBlockChain()
	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	fmt.Fprintf(w, "GO GO GADGET BLOCKCHAIN \n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "\n")

	for _, block := range chain.blocks {
		fmt.Fprintf(w, "Previous hash: %x\n", block.PrevHash)
		fmt.Fprintf(w, "Data: %s\n", block.Data)
		fmt.Fprintf(w, "Current hash: %x\n", block.Hash)
		fmt.Fprintf(w, "NEXT BLOCK \n")
	}
}

func main() {
	http.HandleFunc("/", helloWorldPage)
	http.ListenAndServe("", nil)
}
