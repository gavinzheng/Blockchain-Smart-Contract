package ERC20

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	//"discountedBackend/coinlib/ERC20"
)

var (
	CoinAddr = common.HexToAddress("0xc4dbf323c3c2ae821e4e02afe5a0542c78423e8b") // USDT

	// CoinHash superCoin 合约地址Hash
	CoinHash = CoinAddr.Hash()
)

// Connecter SuperCoin连接者
type Connecter struct {
	ctx  context.Context
	conn *ethclient.Client
	coin *USDTERC20
}

// NewConnecter 生成一个USDT ERC20连接者
func NewConnecter() *Connecter {
	// Dial这里支持传入 ws、http、ipc的多种链接
	// 如果是经常需要调用最好还是使用 ws 方式保持通讯状态
	//conn, err := ethclient.Dial("ws://127.0.0.1:8546")
	conn, err := ethclient.Dial("https://ropsten.infura.io/v3/fdda941e249941418aeb27b0323d03c6")
	if err != nil {
		panic(err)
	}
	coin, err := NewUSDTERC20(CoinAddr, conn)
	if err != nil {
		panic(err)
	}
	return &Connecter{
		ctx:  context.Background(),
		conn: conn,
		coin: coin,
	}
}

// NewConnecterWithDeploy 部署合约，并创建一个connecter
func NewConnecterWithDeploy(ownerAuth *bind.TransactOpts) *Connecter {
	//conn, err := ethclient.Dial("ws://127.0.0.1:8546")
	conn, err := ethclient.Dial("https://ropsten.infura.io/v3/fdda941e249941418aeb27b0323d03c6")
	if err != nil {
		panic(err)
	}

	var total, _ = new(big.Int).SetString("1000000000000000000000000000", 10)
	_, tx, coin, err := DeployUSDTERC20(ownerAuth, conn, total, "usdt erc20", "usdterc201", big.NewInt(18))
	//_, tx, coin, err := DeploySuperCoin(ownerAuth, conn)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	CoinAddr, err = bind.WaitDeployed(ctx, conn, tx)
	if err != nil {
		panic(err)
	}
	CoinHash = CoinAddr.Hash()
	return &Connecter{
		ctx:  ctx,
		conn: conn,
		coin: coin,
	}
}

func (c *Connecter) SetMessage(opts *bind.TransactOpts, newMessage string)(*types.Transaction, error){
	aa, err:=c.SetMessage(opts, newMessage)
	return aa, err
}
//func (c *Connecter) Message(opts *bind.TransactOpts, newMessage string)(*types.Transaction, error){
//	aa, err:=c.coin.USDTERC20Caller.contract.
//	return aa, err
//}

// BlockNumber 当前块高度
func (c *Connecter) BlockNumber() *big.Int {
	blockNumber, err := c.conn.BlockByNumber(c.ctx, nil)
	if err != nil {
		log.Fatal("Get block number error", err)
		return big.NewInt(0)
	}
	return blockNumber.Number()
}

// ContractName SuperCoin名称
func (c *Connecter) ContractName() string {
	name, err := c.coin.Name(nil)
	if err != nil {
		log.Fatal("Get contract name error", err)
		return ""
	}
	return name
}

// BalanceOfEth 以太币余额
func (c *Connecter) BalanceOfEth(addr common.Address) *big.Int {
	balance, err := c.conn.BalanceAt(c.ctx, addr, nil)
	if err != nil {
		log.Fatal("Get address eth balance error", err)
		return big.NewInt(0)
	}
	return balance
}

// BalanceOfCoin SuperCoin余额
func (c *Connecter) BalanceOfCoin(addr common.Address) *big.Int {
	balance, err := c.coin.BalanceOf(nil, addr)
	if err != nil {
		log.Fatal("Get address SuperCoin balance error", err)
		return big.NewInt(0)
	}
	return balance
}

// TotalSupply SuperCoin已铸币量
func (c *Connecter) TotalSupply() *big.Int {
	totalSupply, err := c.coin.TotalSupply(nil)
	if err != nil {
		log.Fatal("Get SuperCoin totalSuply error", err)
		return big.NewInt(0)
	}
	return totalSupply
}

// AuthAccount 解锁账户
// 正式使用时候，此处应该限制GasPrice和GasLimit
func AuthAccount(key, passphrase string) *bind.TransactOpts {
	auth, err := bind.NewTransactor(strings.NewReader(key), passphrase)
	if err != nil {
		log.Fatalf("Auth account error: %v", err)
		return nil
	}
	log.Printf("Gas price is: %v, Gas limit is: %v", auth.GasPrice, auth.GasLimit)
	return auth
}

// TransferCoin SuperCoin转账
func (c *Connecter) TransferCoin(fromAuth *bind.TransactOpts, to common.Address, amount *big.Int) bool {
	tx, err := c.coin.Transfer(fromAuth, to, amount)
	if err != nil {
		log.Fatalf("Transfer SuperCoin from %s to %s amount %s error: %v", fromAuth.From.String(), to.String(), amount.String(), err)
		return false
	}
	// 等待执行
	receipt, err := bind.WaitMined(c.ctx, c.conn, tx)
	if err != nil {
		log.Fatalf("Wait Transfer Mined error: %v", err)
		return false
	}
	return receipt.Status == 1
}

// TransferLogs SuperCoin转账记录
func (c *Connecter) TransferLogs(froms []common.Address, tos []common.Address) {
	iter, err := c.coin.FilterTransfer(&bind.FilterOpts{}, froms, tos)
	defer iter.Close()
	if err != nil {
		panic(err)
	}
	for {
		if !iter.Next() {
			break
		}
		log.Printf("Transfer log: %s", iter.Event.String())
	}
}

func (s *USDTERC20Transfer) String() string {
	return fmt.Sprintf("From: %s, To: %s, Amount: %s", s.From.Hex(), s.To.Hex(), s.Value)
}


// WatchTransferAndMint 监听transfer和mint事件
func (c *Connecter) WatchTransferAndMint() {
	transferCh := make(chan *USDTERC20Transfer)
	transferSub, err := c.coin.WatchTransfer(&bind.WatchOpts{}, transferCh, nil, nil)
	if err != nil {
		panic(err)
	}

	//mintCh := make(chan *SuperCoinMint)
	//mintSub, err := c.coin.WatchMint(&bind.WatchOpts{}, mintCh, nil)
	//if err != nil {
	//	panic(err)
	//}

	for {
		select {
		case t := <-transferCh:
			log.Printf("Transfer From: %s, To: %s, Amount: %d", t.From.String(), t.To.String(), t.Value)
		case e := <-transferSub.Err():
			panic(e)
		//case m := <-mintCh:
		//	log.Printf("Mint To: %s, Amount: %d", m.To.String(), m.Amount)
		//case e := <-mintSub.Err():
		//	panic(e)
		}
	}
}
