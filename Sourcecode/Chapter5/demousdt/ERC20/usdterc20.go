// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC20

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// USDTERC20ABI is the input ABI used to generate the binding from.
const USDTERC20ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_upgradedAddress\",\"type\":\"address\"}],\"name\":\"deprecate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"deprecated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_evilUser\",\"type\":\"address\"}],\"name\":\"addBlackList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"upgradedAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maximumFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newMessage\",\"type\":\"string\"}],\"name\":\"setMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_maker\",\"type\":\"address\"}],\"name\":\"getBlackListStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newBasisPoints\",\"type\":\"uint256\"},{\"name\":\"newMaxFee\",\"type\":\"uint256\"}],\"name\":\"setParams\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"issue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"basisPointsRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"message\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isBlackListed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_clearedUser\",\"type\":\"address\"}],\"name\":\"removeBlackList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_UINT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_blackListedUser\",\"type\":\"address\"}],\"name\":\"destroyBlackFunds\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_initialSupply\",\"type\":\"uint256\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_decimals\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Issue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"Deprecate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"feeBasisPoints\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"maxFee\",\"type\":\"uint256\"}],\"name\":\"Params\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_blackListedUser\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"DestroyedBlackFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"AddedBlackList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"RemovedBlackList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"}]"

// USDTERC20Bin is the compiled bytecode used for deploying new contracts.
var USDTERC20Bin = "0x608060405260008060146101000a81548160ff021916908315150217905550600060035560006004556040805190810160405280600e81526020017f5468697320697320612074657374000000000000000000000000000000000000815250600b908051906020019062000075929190620001d3565b503480156200008357600080fd5b50604051620031673803806200316783398101806040528101908080519060200190929190805182019291906020018051820192919060200180519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600181905550826007908051906020019062000128929190620001d3565b50816008908051906020019062000141929190620001d3565b508060098190555083600260008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506000600a60146101000a81548160ff0219169083151502179055505050505062000282565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200021657805160ff191683800117855562000247565b8280016001018555821562000247579182015b828111156200024657825182559160200191906001019062000229565b5b5090506200025691906200025a565b5090565b6200027f91905b808211156200027b57600081600090555060010162000261565b5090565b90565b612ed580620002926000396000f3006080604052600436106101ac576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde03146101b15780630753c30c14610241578063095ea7b3146102845780630e136b19146102d15780630ecb93c01461030057806318160ddd1461034357806323b872dd1461036e57806326976e3f146103db57806327e235e314610432578063313ce5671461048957806335390714146104b4578063368b8772146104df5780633eaaf86b146105485780633f4ba83a1461057357806359bf1abe1461058a5780635c658165146105e55780635c975abb1461065c57806370a082311461068b5780638456cb59146106e2578063893d20e8146106f95780638da5cb5b1461075057806395d89b41146107a7578063a9059cbb14610837578063c0324c7714610884578063cc872b66146108bb578063db006a75146108e8578063dd62ed3e14610915578063dd644f721461098c578063e21f37ce146109b7578063e47d606014610a47578063e4997dc514610aa2578063e5b5019a14610ae5578063f2fde38b14610b10578063f3bdc22814610b53575b600080fd5b3480156101bd57600080fd5b506101c6610b96565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102065780820151818401526020810190506101eb565b50505050905090810190601f1680156102335780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561024d57600080fd5b50610282600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610c34565b005b34801561029057600080fd5b506102cf600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610d51565b005b3480156102dd57600080fd5b506102e6610ea4565b604051808215151515815260200191505060405180910390f35b34801561030c57600080fd5b50610341600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610eb7565b005b34801561034f57600080fd5b50610358610fd0565b6040518082815260200191505060405180910390f35b34801561037a57600080fd5b506103d9600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506110b8565b005b3480156103e757600080fd5b506103f061129d565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561043e57600080fd5b50610473600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506112c3565b6040518082815260200191505060405180910390f35b34801561049557600080fd5b5061049e6112db565b6040518082815260200191505060405180910390f35b3480156104c057600080fd5b506104c96112e1565b6040518082815260200191505060405180910390f35b3480156104eb57600080fd5b50610546600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506112e7565b005b34801561055457600080fd5b5061055d611301565b6040518082815260200191505060405180910390f35b34801561057f57600080fd5b50610588611307565b005b34801561059657600080fd5b506105cb600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506113c5565b604051808215151515815260200191505060405180910390f35b3480156105f157600080fd5b50610646600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061141b565b6040518082815260200191505060405180910390f35b34801561066857600080fd5b50610671611440565b604051808215151515815260200191505060405180910390f35b34801561069757600080fd5b506106cc600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611453565b6040518082815260200191505060405180910390f35b3480156106ee57600080fd5b506106f761157a565b005b34801561070557600080fd5b5061070e61163a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561075c57600080fd5b50610765611663565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156107b357600080fd5b506107bc611688565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156107fc5780820151818401526020810190506107e1565b50505050905090810190601f1680156108295780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561084357600080fd5b50610882600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611726565b005b34801561089057600080fd5b506108b960048036038101908080359060200190929190803590602001909291905050506118d5565b005b3480156108c757600080fd5b506108e6600480360381019080803590602001909291905050506119ba565b005b3480156108f457600080fd5b5061091360048036038101908080359060200190929190505050611bb1565b005b34801561092157600080fd5b50610976600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611d44565b6040518082815260200191505060405180910390f35b34801561099857600080fd5b506109a1611ea1565b6040518082815260200191505060405180910390f35b3480156109c357600080fd5b506109cc611ea7565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610a0c5780820151818401526020810190506109f1565b50505050905090810190601f168015610a395780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b348015610a5357600080fd5b50610a88600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611f45565b604051808215151515815260200191505060405180910390f35b348015610aae57600080fd5b50610ae3600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611f65565b005b348015610af157600080fd5b50610afa61207e565b6040518082815260200191505060405180910390f35b348015610b1c57600080fd5b50610b51600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506120a2565b005b348015610b5f57600080fd5b50610b94600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612177565b005b60078054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610c2c5780601f10610c0157610100808354040283529160200191610c2c565b820191906000526020600020905b815481529060010190602001808311610c0f57829003601f168201915b505050505081565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610c8f57600080fd5b6001600a60146101000a81548160ff02191690831515021790555080600a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fcc358699805e9a8b7f77b522628c7cb9abd07d9efb86b6fb616af1609036a99e81604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b604060048101600036905010151515610d6957600080fd5b600a60149054906101000a900460ff1615610e9457600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663aee92d333385856040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b158015610e7757600080fd5b505af1158015610e8b573d6000803e3d6000fd5b50505050610e9f565b610e9e83836122fb565b5b505050565b600a60149054906101000a900460ff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610f1257600080fd5b6001600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc81604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b6000600a60149054906101000a900460ff16156110af57600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561106d57600080fd5b505af1158015611081573d6000803e3d6000fd5b505050506040513d602081101561109757600080fd5b810190808051906020019092919050505090506110b5565b60015490505b90565b600060149054906101000a900460ff161515156110d457600080fd5b600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615151561112d57600080fd5b600a60149054906101000a900460ff161561128c57600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638b477adb338585856040518563ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001945050505050600060405180830381600087803b15801561126f57600080fd5b505af1158015611283573d6000803e3d6000fd5b50505050611298565b611297838383612498565b5b505050565b600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60026020528060005260406000206000915090505481565b60095481565b60045481565b80600b90805190602001906112fd929190612e04565b5050565b60015481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561136257600080fd5b600060149054906101000a900460ff16151561137d57600080fd5b60008060146101000a81548160ff0219169083151502179055507f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b6000600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b6005602052816000526040600020602052806000526040600020600091509150505481565b600060149054906101000a900460ff1681565b6000600a60149054906101000a900460ff161561156957600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b15801561152757600080fd5b505af115801561153b573d6000803e3d6000fd5b505050506040513d602081101561155157600080fd5b81019080805190602001909291905050509050611575565b6115728261293f565b90505b919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156115d557600080fd5b600060149054906101000a900460ff161515156115f157600080fd5b6001600060146101000a81548160ff0219169083151502179055507f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60088054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561171e5780601f106116f35761010080835404028352916020019161171e565b820191906000526020600020905b81548152906001019060200180831161170157829003601f168201915b505050505081565b600060149054906101000a900460ff1615151561174257600080fd5b600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615151561179b57600080fd5b600a60149054906101000a900460ff16156118c657600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636e18980a3384846040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b1580156118a957600080fd5b505af11580156118bd573d6000803e3d6000fd5b505050506118d1565b6118d08282612988565b5b5050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561193057600080fd5b60148210151561193f57600080fd5b60328110151561194e57600080fd5b8160038190555061196d600954600a0a82612cf090919063ffffffff16565b6004819055507fb044a1e409eac5c48e5af22d4af52670dd1a99059537a78b31b48c6500a6354e600354600454604051808381526020018281526020019250505060405180910390a15050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611a1557600080fd5b6001548160015401111515611a2957600080fd5b600260008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481600260008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205401111515611af957600080fd5b80600260008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540192505081905550806001600082825401925050819055507fcb8241adb0c3fdb35b70c24ce35c5eb0c17af7431c99f827d44a445ca624176a816040518082815260200191505060405180910390a150565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611c0c57600080fd5b8060015410151515611c1d57600080fd5b80600260008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410151515611c8c57600080fd5b8060016000828254039250508190555080600260008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055507f702d5967f45f6513a38ffc42d6ba9bf230bd40e8f53b16363c7eb4fd2deb9a44816040518082815260200191505060405180910390a150565b6000600a60149054906101000a900460ff1615611e8e57600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e84846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b158015611e4c57600080fd5b505af1158015611e60573d6000803e3d6000fd5b505050506040513d6020811015611e7657600080fd5b81019080805190602001909291905050509050611e9b565b611e988383612d2b565b90505b92915050565b60035481565b600b8054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611f3d5780601f10611f1257610100808354040283529160200191611f3d565b820191906000526020600020905b815481529060010190602001808311611f2057829003601f168201915b505050505081565b60066020528060005260406000206000915054906101000a900460ff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611fc057600080fd5b6000600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507fd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c81604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156120fd57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151561217457806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156121d457600080fd5b600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151561222c57600080fd5b61223582611453565b90506000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550806001600082825403925050819055507f61e6e66b0d6339b2980aecc6ccc0039736791f0ccde9ed512e789a7fbdd698c68282604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a15050565b60406004810160003690501015151561231357600080fd5b600082141580156123a157506000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414155b1515156123ad57600080fd5b81600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040518082815260200191505060405180910390a3505050565b60008060006060600481016000369050101515156124b557600080fd5b600560008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054935061255d61271061254f60035488612cf090919063ffffffff16565b612db290919063ffffffff16565b925060045483111561256f5760045492505b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84101561262b576125aa8585612dcd90919063ffffffff16565b600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b61263e8386612dcd90919063ffffffff16565b915061269285600260008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054612dcd90919063ffffffff16565b600260008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061272782600260008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054612de690919063ffffffff16565b600260008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060008311156128d1576127e683600260008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054612de690919063ffffffff16565b600260008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef856040518082815260200191505060405180910390a35b8573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a350505050505050565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000806040600481016000369050101515156129a357600080fd5b6129cc6127106129be60035487612cf090919063ffffffff16565b612db290919063ffffffff16565b92506004548311156129de5760045492505b6129f18385612dcd90919063ffffffff16565b9150612a4584600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054612dcd90919063ffffffff16565b600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550612ada82600260008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054612de690919063ffffffff16565b600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506000831115612c8457612b9983600260008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054612de690919063ffffffff16565b600260008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef856040518082815260200191505060405180910390a35b8473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a35050505050565b6000806000841415612d055760009150612d24565b8284029050828482811515612d1657fe5b04141515612d2057fe5b8091505b5092915050565b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6000808284811515612dc057fe5b0490508091505092915050565b6000828211151515612ddb57fe5b818303905092915050565b6000808284019050838110151515612dfa57fe5b8091505092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10612e4557805160ff1916838001178555612e73565b82800160010185558215612e73579182015b82811115612e72578251825591602001919060010190612e57565b5b509050612e809190612e84565b5090565b612ea691905b80821115612ea2576000816000905550600101612e8a565b5090565b905600a165627a7a723058200d04620b048cd4dea6407fad3a85f5438b2f5bc473af9945aa4efee67a6f6cdf0029"

// DeployUSDTERC20 deploys a new Ethereum contract, binding an instance of USDTERC20 to it.
func DeployUSDTERC20(auth *bind.TransactOpts, backend bind.ContractBackend, _initialSupply *big.Int, _name string, _symbol string, _decimals *big.Int) (common.Address, *types.Transaction, *USDTERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(USDTERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(USDTERC20Bin), backend, _initialSupply, _name, _symbol, _decimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &USDTERC20{USDTERC20Caller: USDTERC20Caller{contract: contract}, USDTERC20Transactor: USDTERC20Transactor{contract: contract}, USDTERC20Filterer: USDTERC20Filterer{contract: contract}}, nil
}

// USDTERC20 is an auto generated Go binding around an Ethereum contract.
type USDTERC20 struct {
	USDTERC20Caller     // Read-only binding to the contract
	USDTERC20Transactor // Write-only binding to the contract
	USDTERC20Filterer   // Log filterer for contract events
}

// USDTERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type USDTERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDTERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type USDTERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDTERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type USDTERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDTERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type USDTERC20Session struct {
	Contract     *USDTERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// USDTERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type USDTERC20CallerSession struct {
	Contract *USDTERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// USDTERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type USDTERC20TransactorSession struct {
	Contract     *USDTERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// USDTERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type USDTERC20Raw struct {
	Contract *USDTERC20 // Generic contract binding to access the raw methods on
}

// USDTERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type USDTERC20CallerRaw struct {
	Contract *USDTERC20Caller // Generic read-only contract binding to access the raw methods on
}

// USDTERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type USDTERC20TransactorRaw struct {
	Contract *USDTERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewUSDTERC20 creates a new instance of USDTERC20, bound to a specific deployed contract.
func NewUSDTERC20(address common.Address, backend bind.ContractBackend) (*USDTERC20, error) {
	contract, err := bindUSDTERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &USDTERC20{USDTERC20Caller: USDTERC20Caller{contract: contract}, USDTERC20Transactor: USDTERC20Transactor{contract: contract}, USDTERC20Filterer: USDTERC20Filterer{contract: contract}}, nil
}

// NewUSDTERC20Caller creates a new read-only instance of USDTERC20, bound to a specific deployed contract.
func NewUSDTERC20Caller(address common.Address, caller bind.ContractCaller) (*USDTERC20Caller, error) {
	contract, err := bindUSDTERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &USDTERC20Caller{contract: contract}, nil
}

// NewUSDTERC20Transactor creates a new write-only instance of USDTERC20, bound to a specific deployed contract.
func NewUSDTERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*USDTERC20Transactor, error) {
	contract, err := bindUSDTERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &USDTERC20Transactor{contract: contract}, nil
}

// NewUSDTERC20Filterer creates a new log filterer instance of USDTERC20, bound to a specific deployed contract.
func NewUSDTERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*USDTERC20Filterer, error) {
	contract, err := bindUSDTERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &USDTERC20Filterer{contract: contract}, nil
}

// bindUSDTERC20 binds a generic wrapper to an already deployed contract.
func bindUSDTERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(USDTERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_USDTERC20 *USDTERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _USDTERC20.Contract.USDTERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_USDTERC20 *USDTERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDTERC20.Contract.USDTERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_USDTERC20 *USDTERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDTERC20.Contract.USDTERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_USDTERC20 *USDTERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _USDTERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_USDTERC20 *USDTERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDTERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_USDTERC20 *USDTERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDTERC20.Contract.contract.Transact(opts, method, params...)
}

// MAXUINT is a free data retrieval call binding the contract method 0xe5b5019a.
//
// Solidity: function MAX_UINT() constant returns(uint256)
func (_USDTERC20 *USDTERC20Caller) MAXUINT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _USDTERC20.contract.Call(opts, out, "MAX_UINT")
	return *ret0, err
}

// MAXUINT is a free data retrieval call binding the contract method 0xe5b5019a.
//
// Solidity: function MAX_UINT() constant returns(uint256)
func (_USDTERC20 *USDTERC20Session) MAXUINT() (*big.Int, error) {
	return _USDTERC20.Contract.MAXUINT(&_USDTERC20.CallOpts)
}

// MAXUINT is a free data retrieval call binding the contract method 0xe5b5019a.
//
// Solidity: function MAX_UINT() constant returns(uint256)
func (_USDTERC20 *USDTERC20CallerSession) MAXUINT() (*big.Int, error) {
	return _USDTERC20.Contract.MAXUINT(&_USDTERC20.CallOpts)
}

// UsdttotalSupply is a free data retrieval call binding the contract method 0x3eaaf86b.
//
// Solidity: function _totalSupply() constant returns(uint256)
func (_USDTERC20 *USDTERC20Caller) UsdttotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _USDTERC20.contract.Call(opts, out, "_totalSupply")
	return *ret0, err
}

// UsdttotalSupply is a free data retrieval call binding the contract method 0x3eaaf86b.
//
// Solidity: function _totalSupply() constant returns(uint256)
func (_USDTERC20 *USDTERC20Session) UsdttotalSupply() (*big.Int, error) {
	return _USDTERC20.Contract.UsdttotalSupply(&_USDTERC20.CallOpts)
}

// UsdttotalSupply is a free data retrieval call binding the contract method 0x3eaaf86b.
//
// Solidity: function _totalSupply() constant returns(uint256)
func (_USDTERC20 *USDTERC20CallerSession) UsdttotalSupply() (*big.Int, error) {
	return _USDTERC20.Contract.UsdttotalSupply(&_USDTERC20.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256 remaining)
func (_USDTERC20 *USDTERC20Caller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _USDTERC20.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256 remaining)
func (_USDTERC20 *USDTERC20Session) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _USDTERC20.Contract.Allowance(&_USDTERC20.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256 remaining)
func (_USDTERC20 *USDTERC20CallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _USDTERC20.Contract.Allowance(&_USDTERC20.CallOpts, _owner, _spender)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) constant returns(uint256)
func (_USDTERC20 *USDTERC20Caller) Allowed(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _USDTERC20.contract.Call(opts, out, "allowed", arg0, arg1)
	return *ret0, err
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) constant returns(uint256)
func (_USDTERC20 *USDTERC20Session) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _USDTERC20.Contract.Allowed(&_USDTERC20.CallOpts, arg0, arg1)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) constant returns(uint256)
func (_USDTERC20 *USDTERC20CallerSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _USDTERC20.Contract.Allowed(&_USDTERC20.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) constant returns(uint256)
func (_USDTERC20 *USDTERC20Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _USDTERC20.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) constant returns(uint256)
func (_USDTERC20 *USDTERC20Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _USDTERC20.Contract.BalanceOf(&_USDTERC20.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) constant returns(uint256)
func (_USDTERC20 *USDTERC20CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _USDTERC20.Contract.BalanceOf(&_USDTERC20.CallOpts, who)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) constant returns(uint256)
func (_USDTERC20 *USDTERC20Caller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _USDTERC20.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) constant returns(uint256)
func (_USDTERC20 *USDTERC20Session) Balances(arg0 common.Address) (*big.Int, error) {
	return _USDTERC20.Contract.Balances(&_USDTERC20.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) constant returns(uint256)
func (_USDTERC20 *USDTERC20CallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _USDTERC20.Contract.Balances(&_USDTERC20.CallOpts, arg0)
}

// BasisPointsRate is a free data retrieval call binding the contract method 0xdd644f72.
//
// Solidity: function basisPointsRate() constant returns(uint256)
func (_USDTERC20 *USDTERC20Caller) BasisPointsRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _USDTERC20.contract.Call(opts, out, "basisPointsRate")
	return *ret0, err
}

// BasisPointsRate is a free data retrieval call binding the contract method 0xdd644f72.
//
// Solidity: function basisPointsRate() constant returns(uint256)
func (_USDTERC20 *USDTERC20Session) BasisPointsRate() (*big.Int, error) {
	return _USDTERC20.Contract.BasisPointsRate(&_USDTERC20.CallOpts)
}

// BasisPointsRate is a free data retrieval call binding the contract method 0xdd644f72.
//
// Solidity: function basisPointsRate() constant returns(uint256)
func (_USDTERC20 *USDTERC20CallerSession) BasisPointsRate() (*big.Int, error) {
	return _USDTERC20.Contract.BasisPointsRate(&_USDTERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_USDTERC20 *USDTERC20Caller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _USDTERC20.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_USDTERC20 *USDTERC20Session) Decimals() (*big.Int, error) {
	return _USDTERC20.Contract.Decimals(&_USDTERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_USDTERC20 *USDTERC20CallerSession) Decimals() (*big.Int, error) {
	return _USDTERC20.Contract.Decimals(&_USDTERC20.CallOpts)
}

// Deprecated is a free data retrieval call binding the contract method 0x0e136b19.
//
// Solidity: function deprecated() constant returns(bool)
func (_USDTERC20 *USDTERC20Caller) Deprecated(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _USDTERC20.contract.Call(opts, out, "deprecated")
	return *ret0, err
}

// Deprecated is a free data retrieval call binding the contract method 0x0e136b19.
//
// Solidity: function deprecated() constant returns(bool)
func (_USDTERC20 *USDTERC20Session) Deprecated() (bool, error) {
	return _USDTERC20.Contract.Deprecated(&_USDTERC20.CallOpts)
}

// Deprecated is a free data retrieval call binding the contract method 0x0e136b19.
//
// Solidity: function deprecated() constant returns(bool)
func (_USDTERC20 *USDTERC20CallerSession) Deprecated() (bool, error) {
	return _USDTERC20.Contract.Deprecated(&_USDTERC20.CallOpts)
}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) constant returns(bool)
func (_USDTERC20 *USDTERC20Caller) GetBlackListStatus(opts *bind.CallOpts, _maker common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _USDTERC20.contract.Call(opts, out, "getBlackListStatus", _maker)
	return *ret0, err
}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) constant returns(bool)
func (_USDTERC20 *USDTERC20Session) GetBlackListStatus(_maker common.Address) (bool, error) {
	return _USDTERC20.Contract.GetBlackListStatus(&_USDTERC20.CallOpts, _maker)
}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) constant returns(bool)
func (_USDTERC20 *USDTERC20CallerSession) GetBlackListStatus(_maker common.Address) (bool, error) {
	return _USDTERC20.Contract.GetBlackListStatus(&_USDTERC20.CallOpts, _maker)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_USDTERC20 *USDTERC20Caller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _USDTERC20.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_USDTERC20 *USDTERC20Session) GetOwner() (common.Address, error) {
	return _USDTERC20.Contract.GetOwner(&_USDTERC20.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_USDTERC20 *USDTERC20CallerSession) GetOwner() (common.Address, error) {
	return _USDTERC20.Contract.GetOwner(&_USDTERC20.CallOpts)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) constant returns(bool)
func (_USDTERC20 *USDTERC20Caller) IsBlackListed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _USDTERC20.contract.Call(opts, out, "isBlackListed", arg0)
	return *ret0, err
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) constant returns(bool)
func (_USDTERC20 *USDTERC20Session) IsBlackListed(arg0 common.Address) (bool, error) {
	return _USDTERC20.Contract.IsBlackListed(&_USDTERC20.CallOpts, arg0)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) constant returns(bool)
func (_USDTERC20 *USDTERC20CallerSession) IsBlackListed(arg0 common.Address) (bool, error) {
	return _USDTERC20.Contract.IsBlackListed(&_USDTERC20.CallOpts, arg0)
}

// MaximumFee is a free data retrieval call binding the contract method 0x35390714.
//
// Solidity: function maximumFee() constant returns(uint256)
func (_USDTERC20 *USDTERC20Caller) MaximumFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _USDTERC20.contract.Call(opts, out, "maximumFee")
	return *ret0, err
}

// MaximumFee is a free data retrieval call binding the contract method 0x35390714.
//
// Solidity: function maximumFee() constant returns(uint256)
func (_USDTERC20 *USDTERC20Session) MaximumFee() (*big.Int, error) {
	return _USDTERC20.Contract.MaximumFee(&_USDTERC20.CallOpts)
}

// MaximumFee is a free data retrieval call binding the contract method 0x35390714.
//
// Solidity: function maximumFee() constant returns(uint256)
func (_USDTERC20 *USDTERC20CallerSession) MaximumFee() (*big.Int, error) {
	return _USDTERC20.Contract.MaximumFee(&_USDTERC20.CallOpts)
}

// Message is a free data retrieval call binding the contract method 0xe21f37ce.
//
// Solidity: function message() constant returns(string)
func (_USDTERC20 *USDTERC20Caller) Message(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _USDTERC20.contract.Call(opts, out, "message")
	return *ret0, err
}

// Message is a free data retrieval call binding the contract method 0xe21f37ce.
//
// Solidity: function message() constant returns(string)
func (_USDTERC20 *USDTERC20Session) Message() (string, error) {
	return _USDTERC20.Contract.Message(&_USDTERC20.CallOpts)
}

// Message is a free data retrieval call binding the contract method 0xe21f37ce.
//
// Solidity: function message() constant returns(string)
func (_USDTERC20 *USDTERC20CallerSession) Message() (string, error) {
	return _USDTERC20.Contract.Message(&_USDTERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_USDTERC20 *USDTERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _USDTERC20.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_USDTERC20 *USDTERC20Session) Name() (string, error) {
	return _USDTERC20.Contract.Name(&_USDTERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_USDTERC20 *USDTERC20CallerSession) Name() (string, error) {
	return _USDTERC20.Contract.Name(&_USDTERC20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_USDTERC20 *USDTERC20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _USDTERC20.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_USDTERC20 *USDTERC20Session) Owner() (common.Address, error) {
	return _USDTERC20.Contract.Owner(&_USDTERC20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_USDTERC20 *USDTERC20CallerSession) Owner() (common.Address, error) {
	return _USDTERC20.Contract.Owner(&_USDTERC20.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_USDTERC20 *USDTERC20Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _USDTERC20.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_USDTERC20 *USDTERC20Session) Paused() (bool, error) {
	return _USDTERC20.Contract.Paused(&_USDTERC20.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_USDTERC20 *USDTERC20CallerSession) Paused() (bool, error) {
	return _USDTERC20.Contract.Paused(&_USDTERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_USDTERC20 *USDTERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _USDTERC20.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_USDTERC20 *USDTERC20Session) Symbol() (string, error) {
	return _USDTERC20.Contract.Symbol(&_USDTERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_USDTERC20 *USDTERC20CallerSession) Symbol() (string, error) {
	return _USDTERC20.Contract.Symbol(&_USDTERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_USDTERC20 *USDTERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _USDTERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_USDTERC20 *USDTERC20Session) TotalSupply() (*big.Int, error) {
	return _USDTERC20.Contract.TotalSupply(&_USDTERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_USDTERC20 *USDTERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _USDTERC20.Contract.TotalSupply(&_USDTERC20.CallOpts)
}

// UpgradedAddress is a free data retrieval call binding the contract method 0x26976e3f.
//
// Solidity: function upgradedAddress() constant returns(address)
func (_USDTERC20 *USDTERC20Caller) UpgradedAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _USDTERC20.contract.Call(opts, out, "upgradedAddress")
	return *ret0, err
}

// UpgradedAddress is a free data retrieval call binding the contract method 0x26976e3f.
//
// Solidity: function upgradedAddress() constant returns(address)
func (_USDTERC20 *USDTERC20Session) UpgradedAddress() (common.Address, error) {
	return _USDTERC20.Contract.UpgradedAddress(&_USDTERC20.CallOpts)
}

// UpgradedAddress is a free data retrieval call binding the contract method 0x26976e3f.
//
// Solidity: function upgradedAddress() constant returns(address)
func (_USDTERC20 *USDTERC20CallerSession) UpgradedAddress() (common.Address, error) {
	return _USDTERC20.Contract.UpgradedAddress(&_USDTERC20.CallOpts)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_USDTERC20 *USDTERC20Transactor) AddBlackList(opts *bind.TransactOpts, _evilUser common.Address) (*types.Transaction, error) {
	return _USDTERC20.contract.Transact(opts, "addBlackList", _evilUser)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_USDTERC20 *USDTERC20Session) AddBlackList(_evilUser common.Address) (*types.Transaction, error) {
	return _USDTERC20.Contract.AddBlackList(&_USDTERC20.TransactOpts, _evilUser)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_USDTERC20 *USDTERC20TransactorSession) AddBlackList(_evilUser common.Address) (*types.Transaction, error) {
	return _USDTERC20.Contract.AddBlackList(&_USDTERC20.TransactOpts, _evilUser)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns()
func (_USDTERC20 *USDTERC20Transactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _USDTERC20.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns()
func (_USDTERC20 *USDTERC20Session) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _USDTERC20.Contract.Approve(&_USDTERC20.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns()
func (_USDTERC20 *USDTERC20TransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _USDTERC20.Contract.Approve(&_USDTERC20.TransactOpts, _spender, _value)
}

// Deprecate is a paid mutator transaction binding the contract method 0x0753c30c.
//
// Solidity: function deprecate(address _upgradedAddress) returns()
func (_USDTERC20 *USDTERC20Transactor) Deprecate(opts *bind.TransactOpts, _upgradedAddress common.Address) (*types.Transaction, error) {
	return _USDTERC20.contract.Transact(opts, "deprecate", _upgradedAddress)
}

// Deprecate is a paid mutator transaction binding the contract method 0x0753c30c.
//
// Solidity: function deprecate(address _upgradedAddress) returns()
func (_USDTERC20 *USDTERC20Session) Deprecate(_upgradedAddress common.Address) (*types.Transaction, error) {
	return _USDTERC20.Contract.Deprecate(&_USDTERC20.TransactOpts, _upgradedAddress)
}

// Deprecate is a paid mutator transaction binding the contract method 0x0753c30c.
//
// Solidity: function deprecate(address _upgradedAddress) returns()
func (_USDTERC20 *USDTERC20TransactorSession) Deprecate(_upgradedAddress common.Address) (*types.Transaction, error) {
	return _USDTERC20.Contract.Deprecate(&_USDTERC20.TransactOpts, _upgradedAddress)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address _blackListedUser) returns()
func (_USDTERC20 *USDTERC20Transactor) DestroyBlackFunds(opts *bind.TransactOpts, _blackListedUser common.Address) (*types.Transaction, error) {
	return _USDTERC20.contract.Transact(opts, "destroyBlackFunds", _blackListedUser)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address _blackListedUser) returns()
func (_USDTERC20 *USDTERC20Session) DestroyBlackFunds(_blackListedUser common.Address) (*types.Transaction, error) {
	return _USDTERC20.Contract.DestroyBlackFunds(&_USDTERC20.TransactOpts, _blackListedUser)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address _blackListedUser) returns()
func (_USDTERC20 *USDTERC20TransactorSession) DestroyBlackFunds(_blackListedUser common.Address) (*types.Transaction, error) {
	return _USDTERC20.Contract.DestroyBlackFunds(&_USDTERC20.TransactOpts, _blackListedUser)
}

// Issue is a paid mutator transaction binding the contract method 0xcc872b66.
//
// Solidity: function issue(uint256 amount) returns()
func (_USDTERC20 *USDTERC20Transactor) Issue(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _USDTERC20.contract.Transact(opts, "issue", amount)
}

// Issue is a paid mutator transaction binding the contract method 0xcc872b66.
//
// Solidity: function issue(uint256 amount) returns()
func (_USDTERC20 *USDTERC20Session) Issue(amount *big.Int) (*types.Transaction, error) {
	return _USDTERC20.Contract.Issue(&_USDTERC20.TransactOpts, amount)
}

// Issue is a paid mutator transaction binding the contract method 0xcc872b66.
//
// Solidity: function issue(uint256 amount) returns()
func (_USDTERC20 *USDTERC20TransactorSession) Issue(amount *big.Int) (*types.Transaction, error) {
	return _USDTERC20.Contract.Issue(&_USDTERC20.TransactOpts, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_USDTERC20 *USDTERC20Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDTERC20.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_USDTERC20 *USDTERC20Session) Pause() (*types.Transaction, error) {
	return _USDTERC20.Contract.Pause(&_USDTERC20.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_USDTERC20 *USDTERC20TransactorSession) Pause() (*types.Transaction, error) {
	return _USDTERC20.Contract.Pause(&_USDTERC20.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 amount) returns()
func (_USDTERC20 *USDTERC20Transactor) Redeem(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _USDTERC20.contract.Transact(opts, "redeem", amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 amount) returns()
func (_USDTERC20 *USDTERC20Session) Redeem(amount *big.Int) (*types.Transaction, error) {
	return _USDTERC20.Contract.Redeem(&_USDTERC20.TransactOpts, amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 amount) returns()
func (_USDTERC20 *USDTERC20TransactorSession) Redeem(amount *big.Int) (*types.Transaction, error) {
	return _USDTERC20.Contract.Redeem(&_USDTERC20.TransactOpts, amount)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_USDTERC20 *USDTERC20Transactor) RemoveBlackList(opts *bind.TransactOpts, _clearedUser common.Address) (*types.Transaction, error) {
	return _USDTERC20.contract.Transact(opts, "removeBlackList", _clearedUser)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_USDTERC20 *USDTERC20Session) RemoveBlackList(_clearedUser common.Address) (*types.Transaction, error) {
	return _USDTERC20.Contract.RemoveBlackList(&_USDTERC20.TransactOpts, _clearedUser)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_USDTERC20 *USDTERC20TransactorSession) RemoveBlackList(_clearedUser common.Address) (*types.Transaction, error) {
	return _USDTERC20.Contract.RemoveBlackList(&_USDTERC20.TransactOpts, _clearedUser)
}

// SetMessage is a paid mutator transaction binding the contract method 0x368b8772.
//
// Solidity: function setMessage(string newMessage) returns()
func (_USDTERC20 *USDTERC20Transactor) SetMessage(opts *bind.TransactOpts, newMessage string) (*types.Transaction, error) {
	return _USDTERC20.contract.Transact(opts, "setMessage", newMessage)
}

// SetMessage is a paid mutator transaction binding the contract method 0x368b8772.
//
// Solidity: function setMessage(string newMessage) returns()
func (_USDTERC20 *USDTERC20Session) SetMessage(newMessage string) (*types.Transaction, error) {
	return _USDTERC20.Contract.SetMessage(&_USDTERC20.TransactOpts, newMessage)
}

// SetMessage is a paid mutator transaction binding the contract method 0x368b8772.
//
// Solidity: function setMessage(string newMessage) returns()
func (_USDTERC20 *USDTERC20TransactorSession) SetMessage(newMessage string) (*types.Transaction, error) {
	return _USDTERC20.Contract.SetMessage(&_USDTERC20.TransactOpts, newMessage)
}

// SetParams is a paid mutator transaction binding the contract method 0xc0324c77.
//
// Solidity: function setParams(uint256 newBasisPoints, uint256 newMaxFee) returns()
func (_USDTERC20 *USDTERC20Transactor) SetParams(opts *bind.TransactOpts, newBasisPoints *big.Int, newMaxFee *big.Int) (*types.Transaction, error) {
	return _USDTERC20.contract.Transact(opts, "setParams", newBasisPoints, newMaxFee)
}

// SetParams is a paid mutator transaction binding the contract method 0xc0324c77.
//
// Solidity: function setParams(uint256 newBasisPoints, uint256 newMaxFee) returns()
func (_USDTERC20 *USDTERC20Session) SetParams(newBasisPoints *big.Int, newMaxFee *big.Int) (*types.Transaction, error) {
	return _USDTERC20.Contract.SetParams(&_USDTERC20.TransactOpts, newBasisPoints, newMaxFee)
}

// SetParams is a paid mutator transaction binding the contract method 0xc0324c77.
//
// Solidity: function setParams(uint256 newBasisPoints, uint256 newMaxFee) returns()
func (_USDTERC20 *USDTERC20TransactorSession) SetParams(newBasisPoints *big.Int, newMaxFee *big.Int) (*types.Transaction, error) {
	return _USDTERC20.Contract.SetParams(&_USDTERC20.TransactOpts, newBasisPoints, newMaxFee)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_USDTERC20 *USDTERC20Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _USDTERC20.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_USDTERC20 *USDTERC20Session) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _USDTERC20.Contract.Transfer(&_USDTERC20.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_USDTERC20 *USDTERC20TransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _USDTERC20.Contract.Transfer(&_USDTERC20.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns()
func (_USDTERC20 *USDTERC20Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _USDTERC20.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns()
func (_USDTERC20 *USDTERC20Session) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _USDTERC20.Contract.TransferFrom(&_USDTERC20.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns()
func (_USDTERC20 *USDTERC20TransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _USDTERC20.Contract.TransferFrom(&_USDTERC20.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_USDTERC20 *USDTERC20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _USDTERC20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_USDTERC20 *USDTERC20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _USDTERC20.Contract.TransferOwnership(&_USDTERC20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_USDTERC20 *USDTERC20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _USDTERC20.Contract.TransferOwnership(&_USDTERC20.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_USDTERC20 *USDTERC20Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDTERC20.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_USDTERC20 *USDTERC20Session) Unpause() (*types.Transaction, error) {
	return _USDTERC20.Contract.Unpause(&_USDTERC20.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_USDTERC20 *USDTERC20TransactorSession) Unpause() (*types.Transaction, error) {
	return _USDTERC20.Contract.Unpause(&_USDTERC20.TransactOpts)
}

// USDTERC20AddedBlackListIterator is returned from FilterAddedBlackList and is used to iterate over the raw logs and unpacked data for AddedBlackList events raised by the USDTERC20 contract.
type USDTERC20AddedBlackListIterator struct {
	Event *USDTERC20AddedBlackList // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *USDTERC20AddedBlackListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDTERC20AddedBlackList)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(USDTERC20AddedBlackList)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *USDTERC20AddedBlackListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDTERC20AddedBlackListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDTERC20AddedBlackList represents a AddedBlackList event raised by the USDTERC20 contract.
type USDTERC20AddedBlackList struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddedBlackList is a free log retrieval operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_USDTERC20 *USDTERC20Filterer) FilterAddedBlackList(opts *bind.FilterOpts) (*USDTERC20AddedBlackListIterator, error) {

	logs, sub, err := _USDTERC20.contract.FilterLogs(opts, "AddedBlackList")
	if err != nil {
		return nil, err
	}
	return &USDTERC20AddedBlackListIterator{contract: _USDTERC20.contract, event: "AddedBlackList", logs: logs, sub: sub}, nil
}

// WatchAddedBlackList is a free log subscription operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_USDTERC20 *USDTERC20Filterer) WatchAddedBlackList(opts *bind.WatchOpts, sink chan<- *USDTERC20AddedBlackList) (event.Subscription, error) {

	logs, sub, err := _USDTERC20.contract.WatchLogs(opts, "AddedBlackList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDTERC20AddedBlackList)
				if err := _USDTERC20.contract.UnpackLog(event, "AddedBlackList", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddedBlackList is a log parse operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_USDTERC20 *USDTERC20Filterer) ParseAddedBlackList(log types.Log) (*USDTERC20AddedBlackList, error) {
	event := new(USDTERC20AddedBlackList)
	if err := _USDTERC20.contract.UnpackLog(event, "AddedBlackList", log); err != nil {
		return nil, err
	}
	return event, nil
}

// USDTERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the USDTERC20 contract.
type USDTERC20ApprovalIterator struct {
	Event *USDTERC20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *USDTERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDTERC20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(USDTERC20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *USDTERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDTERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDTERC20Approval represents a Approval event raised by the USDTERC20 contract.
type USDTERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_USDTERC20 *USDTERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*USDTERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _USDTERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &USDTERC20ApprovalIterator{contract: _USDTERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_USDTERC20 *USDTERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *USDTERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _USDTERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDTERC20Approval)
				if err := _USDTERC20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_USDTERC20 *USDTERC20Filterer) ParseApproval(log types.Log) (*USDTERC20Approval, error) {
	event := new(USDTERC20Approval)
	if err := _USDTERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// USDTERC20DeprecateIterator is returned from FilterDeprecate and is used to iterate over the raw logs and unpacked data for Deprecate events raised by the USDTERC20 contract.
type USDTERC20DeprecateIterator struct {
	Event *USDTERC20Deprecate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *USDTERC20DeprecateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDTERC20Deprecate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(USDTERC20Deprecate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *USDTERC20DeprecateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDTERC20DeprecateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDTERC20Deprecate represents a Deprecate event raised by the USDTERC20 contract.
type USDTERC20Deprecate struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeprecate is a free log retrieval operation binding the contract event 0xcc358699805e9a8b7f77b522628c7cb9abd07d9efb86b6fb616af1609036a99e.
//
// Solidity: event Deprecate(address newAddress)
func (_USDTERC20 *USDTERC20Filterer) FilterDeprecate(opts *bind.FilterOpts) (*USDTERC20DeprecateIterator, error) {

	logs, sub, err := _USDTERC20.contract.FilterLogs(opts, "Deprecate")
	if err != nil {
		return nil, err
	}
	return &USDTERC20DeprecateIterator{contract: _USDTERC20.contract, event: "Deprecate", logs: logs, sub: sub}, nil
}

// WatchDeprecate is a free log subscription operation binding the contract event 0xcc358699805e9a8b7f77b522628c7cb9abd07d9efb86b6fb616af1609036a99e.
//
// Solidity: event Deprecate(address newAddress)
func (_USDTERC20 *USDTERC20Filterer) WatchDeprecate(opts *bind.WatchOpts, sink chan<- *USDTERC20Deprecate) (event.Subscription, error) {

	logs, sub, err := _USDTERC20.contract.WatchLogs(opts, "Deprecate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDTERC20Deprecate)
				if err := _USDTERC20.contract.UnpackLog(event, "Deprecate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeprecate is a log parse operation binding the contract event 0xcc358699805e9a8b7f77b522628c7cb9abd07d9efb86b6fb616af1609036a99e.
//
// Solidity: event Deprecate(address newAddress)
func (_USDTERC20 *USDTERC20Filterer) ParseDeprecate(log types.Log) (*USDTERC20Deprecate, error) {
	event := new(USDTERC20Deprecate)
	if err := _USDTERC20.contract.UnpackLog(event, "Deprecate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// USDTERC20DestroyedBlackFundsIterator is returned from FilterDestroyedBlackFunds and is used to iterate over the raw logs and unpacked data for DestroyedBlackFunds events raised by the USDTERC20 contract.
type USDTERC20DestroyedBlackFundsIterator struct {
	Event *USDTERC20DestroyedBlackFunds // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *USDTERC20DestroyedBlackFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDTERC20DestroyedBlackFunds)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(USDTERC20DestroyedBlackFunds)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *USDTERC20DestroyedBlackFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDTERC20DestroyedBlackFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDTERC20DestroyedBlackFunds represents a DestroyedBlackFunds event raised by the USDTERC20 contract.
type USDTERC20DestroyedBlackFunds struct {
	BlackListedUser common.Address
	Balance         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDestroyedBlackFunds is a free log retrieval operation binding the contract event 0x61e6e66b0d6339b2980aecc6ccc0039736791f0ccde9ed512e789a7fbdd698c6.
//
// Solidity: event DestroyedBlackFunds(address _blackListedUser, uint256 _balance)
func (_USDTERC20 *USDTERC20Filterer) FilterDestroyedBlackFunds(opts *bind.FilterOpts) (*USDTERC20DestroyedBlackFundsIterator, error) {

	logs, sub, err := _USDTERC20.contract.FilterLogs(opts, "DestroyedBlackFunds")
	if err != nil {
		return nil, err
	}
	return &USDTERC20DestroyedBlackFundsIterator{contract: _USDTERC20.contract, event: "DestroyedBlackFunds", logs: logs, sub: sub}, nil
}

// WatchDestroyedBlackFunds is a free log subscription operation binding the contract event 0x61e6e66b0d6339b2980aecc6ccc0039736791f0ccde9ed512e789a7fbdd698c6.
//
// Solidity: event DestroyedBlackFunds(address _blackListedUser, uint256 _balance)
func (_USDTERC20 *USDTERC20Filterer) WatchDestroyedBlackFunds(opts *bind.WatchOpts, sink chan<- *USDTERC20DestroyedBlackFunds) (event.Subscription, error) {

	logs, sub, err := _USDTERC20.contract.WatchLogs(opts, "DestroyedBlackFunds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDTERC20DestroyedBlackFunds)
				if err := _USDTERC20.contract.UnpackLog(event, "DestroyedBlackFunds", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDestroyedBlackFunds is a log parse operation binding the contract event 0x61e6e66b0d6339b2980aecc6ccc0039736791f0ccde9ed512e789a7fbdd698c6.
//
// Solidity: event DestroyedBlackFunds(address _blackListedUser, uint256 _balance)
func (_USDTERC20 *USDTERC20Filterer) ParseDestroyedBlackFunds(log types.Log) (*USDTERC20DestroyedBlackFunds, error) {
	event := new(USDTERC20DestroyedBlackFunds)
	if err := _USDTERC20.contract.UnpackLog(event, "DestroyedBlackFunds", log); err != nil {
		return nil, err
	}
	return event, nil
}

// USDTERC20IssueIterator is returned from FilterIssue and is used to iterate over the raw logs and unpacked data for Issue events raised by the USDTERC20 contract.
type USDTERC20IssueIterator struct {
	Event *USDTERC20Issue // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *USDTERC20IssueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDTERC20Issue)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(USDTERC20Issue)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *USDTERC20IssueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDTERC20IssueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDTERC20Issue represents a Issue event raised by the USDTERC20 contract.
type USDTERC20Issue struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIssue is a free log retrieval operation binding the contract event 0xcb8241adb0c3fdb35b70c24ce35c5eb0c17af7431c99f827d44a445ca624176a.
//
// Solidity: event Issue(uint256 amount)
func (_USDTERC20 *USDTERC20Filterer) FilterIssue(opts *bind.FilterOpts) (*USDTERC20IssueIterator, error) {

	logs, sub, err := _USDTERC20.contract.FilterLogs(opts, "Issue")
	if err != nil {
		return nil, err
	}
	return &USDTERC20IssueIterator{contract: _USDTERC20.contract, event: "Issue", logs: logs, sub: sub}, nil
}

// WatchIssue is a free log subscription operation binding the contract event 0xcb8241adb0c3fdb35b70c24ce35c5eb0c17af7431c99f827d44a445ca624176a.
//
// Solidity: event Issue(uint256 amount)
func (_USDTERC20 *USDTERC20Filterer) WatchIssue(opts *bind.WatchOpts, sink chan<- *USDTERC20Issue) (event.Subscription, error) {

	logs, sub, err := _USDTERC20.contract.WatchLogs(opts, "Issue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDTERC20Issue)
				if err := _USDTERC20.contract.UnpackLog(event, "Issue", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIssue is a log parse operation binding the contract event 0xcb8241adb0c3fdb35b70c24ce35c5eb0c17af7431c99f827d44a445ca624176a.
//
// Solidity: event Issue(uint256 amount)
func (_USDTERC20 *USDTERC20Filterer) ParseIssue(log types.Log) (*USDTERC20Issue, error) {
	event := new(USDTERC20Issue)
	if err := _USDTERC20.contract.UnpackLog(event, "Issue", log); err != nil {
		return nil, err
	}
	return event, nil
}

// USDTERC20ParamsIterator is returned from FilterParams and is used to iterate over the raw logs and unpacked data for Params events raised by the USDTERC20 contract.
type USDTERC20ParamsIterator struct {
	Event *USDTERC20Params // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *USDTERC20ParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDTERC20Params)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(USDTERC20Params)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *USDTERC20ParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDTERC20ParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDTERC20Params represents a Params event raised by the USDTERC20 contract.
type USDTERC20Params struct {
	FeeBasisPoints *big.Int
	MaxFee         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterParams is a free log retrieval operation binding the contract event 0xb044a1e409eac5c48e5af22d4af52670dd1a99059537a78b31b48c6500a6354e.
//
// Solidity: event Params(uint256 feeBasisPoints, uint256 maxFee)
func (_USDTERC20 *USDTERC20Filterer) FilterParams(opts *bind.FilterOpts) (*USDTERC20ParamsIterator, error) {

	logs, sub, err := _USDTERC20.contract.FilterLogs(opts, "Params")
	if err != nil {
		return nil, err
	}
	return &USDTERC20ParamsIterator{contract: _USDTERC20.contract, event: "Params", logs: logs, sub: sub}, nil
}

// WatchParams is a free log subscription operation binding the contract event 0xb044a1e409eac5c48e5af22d4af52670dd1a99059537a78b31b48c6500a6354e.
//
// Solidity: event Params(uint256 feeBasisPoints, uint256 maxFee)
func (_USDTERC20 *USDTERC20Filterer) WatchParams(opts *bind.WatchOpts, sink chan<- *USDTERC20Params) (event.Subscription, error) {

	logs, sub, err := _USDTERC20.contract.WatchLogs(opts, "Params")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDTERC20Params)
				if err := _USDTERC20.contract.UnpackLog(event, "Params", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseParams is a log parse operation binding the contract event 0xb044a1e409eac5c48e5af22d4af52670dd1a99059537a78b31b48c6500a6354e.
//
// Solidity: event Params(uint256 feeBasisPoints, uint256 maxFee)
func (_USDTERC20 *USDTERC20Filterer) ParseParams(log types.Log) (*USDTERC20Params, error) {
	event := new(USDTERC20Params)
	if err := _USDTERC20.contract.UnpackLog(event, "Params", log); err != nil {
		return nil, err
	}
	return event, nil
}

// USDTERC20PauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the USDTERC20 contract.
type USDTERC20PauseIterator struct {
	Event *USDTERC20Pause // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *USDTERC20PauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDTERC20Pause)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(USDTERC20Pause)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *USDTERC20PauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDTERC20PauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDTERC20Pause represents a Pause event raised by the USDTERC20 contract.
type USDTERC20Pause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_USDTERC20 *USDTERC20Filterer) FilterPause(opts *bind.FilterOpts) (*USDTERC20PauseIterator, error) {

	logs, sub, err := _USDTERC20.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &USDTERC20PauseIterator{contract: _USDTERC20.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_USDTERC20 *USDTERC20Filterer) WatchPause(opts *bind.WatchOpts, sink chan<- *USDTERC20Pause) (event.Subscription, error) {

	logs, sub, err := _USDTERC20.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDTERC20Pause)
				if err := _USDTERC20.contract.UnpackLog(event, "Pause", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePause is a log parse operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_USDTERC20 *USDTERC20Filterer) ParsePause(log types.Log) (*USDTERC20Pause, error) {
	event := new(USDTERC20Pause)
	if err := _USDTERC20.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	return event, nil
}

// USDTERC20RedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the USDTERC20 contract.
type USDTERC20RedeemIterator struct {
	Event *USDTERC20Redeem // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *USDTERC20RedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDTERC20Redeem)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(USDTERC20Redeem)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *USDTERC20RedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDTERC20RedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDTERC20Redeem represents a Redeem event raised by the USDTERC20 contract.
type USDTERC20Redeem struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0x702d5967f45f6513a38ffc42d6ba9bf230bd40e8f53b16363c7eb4fd2deb9a44.
//
// Solidity: event Redeem(uint256 amount)
func (_USDTERC20 *USDTERC20Filterer) FilterRedeem(opts *bind.FilterOpts) (*USDTERC20RedeemIterator, error) {

	logs, sub, err := _USDTERC20.contract.FilterLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return &USDTERC20RedeemIterator{contract: _USDTERC20.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0x702d5967f45f6513a38ffc42d6ba9bf230bd40e8f53b16363c7eb4fd2deb9a44.
//
// Solidity: event Redeem(uint256 amount)
func (_USDTERC20 *USDTERC20Filterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *USDTERC20Redeem) (event.Subscription, error) {

	logs, sub, err := _USDTERC20.contract.WatchLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDTERC20Redeem)
				if err := _USDTERC20.contract.UnpackLog(event, "Redeem", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRedeem is a log parse operation binding the contract event 0x702d5967f45f6513a38ffc42d6ba9bf230bd40e8f53b16363c7eb4fd2deb9a44.
//
// Solidity: event Redeem(uint256 amount)
func (_USDTERC20 *USDTERC20Filterer) ParseRedeem(log types.Log) (*USDTERC20Redeem, error) {
	event := new(USDTERC20Redeem)
	if err := _USDTERC20.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	return event, nil
}

// USDTERC20RemovedBlackListIterator is returned from FilterRemovedBlackList and is used to iterate over the raw logs and unpacked data for RemovedBlackList events raised by the USDTERC20 contract.
type USDTERC20RemovedBlackListIterator struct {
	Event *USDTERC20RemovedBlackList // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *USDTERC20RemovedBlackListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDTERC20RemovedBlackList)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(USDTERC20RemovedBlackList)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *USDTERC20RemovedBlackListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDTERC20RemovedBlackListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDTERC20RemovedBlackList represents a RemovedBlackList event raised by the USDTERC20 contract.
type USDTERC20RemovedBlackList struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRemovedBlackList is a free log retrieval operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_USDTERC20 *USDTERC20Filterer) FilterRemovedBlackList(opts *bind.FilterOpts) (*USDTERC20RemovedBlackListIterator, error) {

	logs, sub, err := _USDTERC20.contract.FilterLogs(opts, "RemovedBlackList")
	if err != nil {
		return nil, err
	}
	return &USDTERC20RemovedBlackListIterator{contract: _USDTERC20.contract, event: "RemovedBlackList", logs: logs, sub: sub}, nil
}

// WatchRemovedBlackList is a free log subscription operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_USDTERC20 *USDTERC20Filterer) WatchRemovedBlackList(opts *bind.WatchOpts, sink chan<- *USDTERC20RemovedBlackList) (event.Subscription, error) {

	logs, sub, err := _USDTERC20.contract.WatchLogs(opts, "RemovedBlackList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDTERC20RemovedBlackList)
				if err := _USDTERC20.contract.UnpackLog(event, "RemovedBlackList", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemovedBlackList is a log parse operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_USDTERC20 *USDTERC20Filterer) ParseRemovedBlackList(log types.Log) (*USDTERC20RemovedBlackList, error) {
	event := new(USDTERC20RemovedBlackList)
	if err := _USDTERC20.contract.UnpackLog(event, "RemovedBlackList", log); err != nil {
		return nil, err
	}
	return event, nil
}

// USDTERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the USDTERC20 contract.
type USDTERC20TransferIterator struct {
	Event *USDTERC20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *USDTERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDTERC20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(USDTERC20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *USDTERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDTERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDTERC20Transfer represents a Transfer event raised by the USDTERC20 contract.
type USDTERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_USDTERC20 *USDTERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*USDTERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDTERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &USDTERC20TransferIterator{contract: _USDTERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_USDTERC20 *USDTERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *USDTERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDTERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDTERC20Transfer)
				if err := _USDTERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_USDTERC20 *USDTERC20Filterer) ParseTransfer(log types.Log) (*USDTERC20Transfer, error) {
	event := new(USDTERC20Transfer)
	if err := _USDTERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// USDTERC20UnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the USDTERC20 contract.
type USDTERC20UnpauseIterator struct {
	Event *USDTERC20Unpause // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *USDTERC20UnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDTERC20Unpause)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(USDTERC20Unpause)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *USDTERC20UnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDTERC20UnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDTERC20Unpause represents a Unpause event raised by the USDTERC20 contract.
type USDTERC20Unpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_USDTERC20 *USDTERC20Filterer) FilterUnpause(opts *bind.FilterOpts) (*USDTERC20UnpauseIterator, error) {

	logs, sub, err := _USDTERC20.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &USDTERC20UnpauseIterator{contract: _USDTERC20.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_USDTERC20 *USDTERC20Filterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *USDTERC20Unpause) (event.Subscription, error) {

	logs, sub, err := _USDTERC20.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDTERC20Unpause)
				if err := _USDTERC20.contract.UnpackLog(event, "Unpause", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpause is a log parse operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_USDTERC20 *USDTERC20Filterer) ParseUnpause(log types.Log) (*USDTERC20Unpause, error) {
	event := new(USDTERC20Unpause)
	if err := _USDTERC20.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	return event, nil
}
