package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"github.com/sabbas/inbox/contracts"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)


func main(){
	// connect to an ethereum node  hosted by infura
	blockchain, err := ethclient.Dial("https://rinkeby.infura.io/fYe8qCnWiM9gh&ZAXOVoff")

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

	// Create a new instance of the Inbox contract bound to a specific deployed contract
  contract, err :=contracts.NewInbox(common.HexToAddress("0x491c7fd67ac1f0afeceae79447cd98d2a0e6a9ff"), blockchain)
	if err != nil {
		log.Fatalf("Unable to bind to deployed instance of contract:%v\n")
	}

	fmt.Println(contract.Message(nil))

}