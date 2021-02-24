package main

import (
	"demousdt/ERC20"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"time"
)

//var (
//	srcAddress    = common.HexToAddress("0xA4e71aEeAdaA01FD0f15455f67D7d2CAD32EbFcA") // Keystore
//	targetAddress = common.HexToAddress("0x40362C2A6E4ADdB1388dd38022CA4ebAf1200BA5") // FF
//)
var (
	CoinAddr = common.HexToAddress("0xc4dbf323c3c2ae821e4e02afe5a0542c78423e8b") // USDT

	// CoinHash superCoin 合约地址Hash
	CoinHash = CoinAddr.Hash()
)

func main() {
	keyin := "{\"address\":\"a4e71aeeadaa01fd0f15455f67d7d2cad32ebfca\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"7573755cb636a8712c595b63a69e69c968ff7be194e8757b1122584a63ad2457\",\"cipherparams\":{\"iv\":\"2afb66d1c584804b8f3265fff3b220cb\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"bad691204ff97a29742329338a56c91df858ee383733d1357e95f7b4736ea5a6\"},\"mac\":\"061dae9ec6409181dea894afd52c5acf779e320a970159a91ab334a54a55b4d1\"},\"id\":\"d4a01c73-baeb-44b8-9489-2c6a7df5aa45\",\"version\":3}"
	srcTransactOpt := ERC20.AuthAccount(keyin, "geth123")

	//s := ERC20.NewConnecter()

	conn, err := ethclient.Dial("https://ropsten.infura.io/v3/fdda941e249941418aeb27b0323d03c6")
	if err != nil {
		panic(err)
	}
	coin, err := ERC20.NewUSDTERC20(CoinAddr, conn)
	if err != nil {
		panic(err)
	}


	str,_ := coin.Message(nil)
	log.Printf("Before : message = %s,\n", str )

	//	s.WatchTransferAndMint()
	coin.SetMessage(srcTransactOpt,"USDT ERC20")
	str,_ = coin.Message(nil)
	time.Sleep(30 * time.Second)
	log.Printf("After : message = %s,\n", str )
}