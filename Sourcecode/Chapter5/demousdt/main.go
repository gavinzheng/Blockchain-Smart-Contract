package main

import (
	"demousdt/ERC20"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
)

var (
	//服务端口号
	port          string
	srcAddress    = common.HexToAddress("0xA4e71aEeAdaA01FD0f15455f67D7d2CAD32EbFcA") // Keystore
	targetAddress = common.HexToAddress("0x40362C2A6E4ADdB1388dd38022CA4ebAf1200BA5") // FF
)

func main() {
	keyin := "{\"address\":\"a4e71aeeadaa01fd0f15455f67d7d2cad32ebfca\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"7573755cb636a8712c595b63a69e69c968ff7be194e8757b1122584a63ad2457\",\"cipherparams\":{\"iv\":\"2afb66d1c584804b8f3265fff3b220cb\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"bad691204ff97a29742329338a56c91df858ee383733d1357e95f7b4736ea5a6\"},\"mac\":\"061dae9ec6409181dea894afd52c5acf779e320a970159a91ab334a54a55b4d1\"},\"id\":\"d4a01c73-baeb-44b8-9489-2c6a7df5aa45\",\"version\":3}"
	srcTransactOpt := ERC20.AuthAccount(keyin, "geth123")

	//eth.DCSETHsend()
	//s := ERC20.NewConnecter()
	s:=ERC20.NewConnecterWithDeploy(srcTransactOpt)

	//	s.WatchTransferAndMint()
	log.Printf("Contract Name = %s, BlockNo= %d, BalanceofETH=%d, BalanceofCoin=%d\n", s.ContractName(), s.BlockNumber(), s.BalanceOfEth(srcAddress), s.BalanceOfCoin(srcAddress))

	num, _ := new(big.Int).SetString("100000000000000000000000", 10)
	s.TransferCoin(srcTransactOpt, targetAddress, num)
	log.Printf("Contract Name = %s, BlockNo= %d, BalanceofETH=%d, BalanceofCoin=%d\n", s.ContractName(), s.BlockNumber(), s.BalanceOfEth(srcAddress), s.BalanceOfCoin(srcAddress))
}