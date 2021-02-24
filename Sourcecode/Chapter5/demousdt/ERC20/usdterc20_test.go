package ERC20

import (
	"testing"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/core"
	"math/big"
)

// Test USDTERC20 contract gets deployed correctly
func TestDeployUSDTERC20(t *testing.T) {
	//Setup simulated block chain
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	blockchain := backends.NewSimulatedBackend(alloc,100000000000)

	//Deploy contract
	var total, _ = new(big.Int).SetString("1000000000000000000000000000", 10)
	address, _, _, err := DeployUSDTERC20(auth, blockchain, total, "usdt erc20", "usdterc201", big.NewInt(18))

	// commit all pending transactions
	blockchain.Commit()

	if err != nil {
		t.Fatalf("Failed to deploy the USDT ERC20 contract: %v", err)
	}
	if len(address.Bytes()) == 0 {
		t.Error("Expected a valid deployment address. Received empty address byte array instead")
	}
}

//Test initial message gets set up correctly
func TestGetMessage(t *testing.T) {
	//Setup simulated block chain
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	blockchain := backends.NewSimulatedBackend(alloc,100000000000)

	//Deploy contract
	var total, _ = new(big.Int).SetString("1000000000000000000000000000", 10)
	_, _, contract, _ :=DeployUSDTERC20(auth, blockchain, total, "usdt erc20", "usdterc201", big.NewInt(18))
	// commit all pending transactions
	blockchain.Commit()
	if got, _ := contract.Message(nil); got != "This is a test" {
		t.Errorf("Expected message to be: Hello World. Go: %s", got)
	}
}

// Test message gets updated correctly
func TestSetMessage(t *testing.T) {
	//Setup simulated blockchain
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	blockchain := backends.NewSimulatedBackend(alloc,100000000000)

	//Deploy contract
	var total, _ = new(big.Int).SetString("1000000000000000000000000000", 10)
	_, _, contract, _ :=DeployUSDTERC20(auth, blockchain, total, "usdt erc20", "usdterc201", big.NewInt(18))

	// commit all pending transactions
	blockchain.Commit()
	contract.SetMessage(&bind.TransactOpts{
		From:auth.From,
		Signer:auth.Signer,
		Value: nil,
	}, "Hello World")
	blockchain.Commit()
	if got, _ := contract.Message(nil); got != "Hello World" {
		t.Errorf("Expected message to be: Hello World. Go: %s", got)
	}
}
